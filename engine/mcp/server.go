// Package mcp is the MMS Model Context Protocol server: it gives an AI
// agent that writes MMS the engine's checks, compiled output, test
// interpreter, reference and vanilla data as tools, and the docs as
// resources. `mms mcp` runs it over stdin and stdout.
//
// The protocol is JSON-RPC 2.0, one message per line. Only what the
// tools and resources need is here: initialize, ping, tools/list,
// tools/call, resources/list and resources/read.
package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"slices"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// protocolVersions are the protocol revisions the server speaks, newest
// first. None of what it uses differs between them.
var protocolVersions = []string{"2025-06-18", "2025-03-26", "2024-11-05"}

// Server answers one client.
type Server struct {
	// Version is reported to the client, as in `mms version`.
	Version string
	// Log receives a line per failed message; nil for none.
	Log io.Writer

	pack  *spec.Pack
	tools []*tool
	ref   lazyReference
}

// NewServer returns a server that checks against pack.
func NewServer(pack *spec.Pack) *Server {
	s := &Server{pack: pack, Version: "dev"}
	s.tools = s.toolList()
	return s
}

type request struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      *json.RawMessage `json:"id,omitempty"`
	Method  string           `json:"method"`
	Params  json.RawMessage  `json:"params,omitempty"`
}

type response struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      *json.RawMessage `json:"id"`
	Result  any              `json:"result,omitempty"`
	Error   *rpcError        `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	codeParseError     = -32700
	codeInvalidRequest = -32600
	codeMethodNotFound = -32601
	codeInvalidParams  = -32602
	codeNotFound       = -32002 // a resource that does not exist
)

// Run reads requests from r and writes answers to w until r ends.
func (s *Server) Run(ctx context.Context, r io.Reader, w io.Writer) error {
	in := bufio.NewReader(r)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	for {
		line, err := in.ReadBytes('\n')
		if len(line) > 0 {
			if resp := s.handle(ctx, line); resp != nil {
				if werr := enc.Encode(resp); werr != nil {
					return werr
				}
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// handle answers one message, or returns nil for a notification and for
// a blank line.
func (s *Server) handle(ctx context.Context, line []byte) *response {
	if len(bytesTrim(line)) == 0 {
		return nil
	}
	var req request
	if err := json.Unmarshal(line, &req); err != nil {
		return &response{JSONRPC: "2.0", Error: &rpcError{codeParseError, "not a JSON-RPC message: " + err.Error()}}
	}
	if req.Method == "" {
		// A response to a request of ours; the server sends none.
		return nil
	}
	result, rerr := s.dispatch(ctx, req.Method, req.Params)
	if req.ID == nil {
		return nil
	}
	if rerr != nil {
		if s.Log != nil {
			fmt.Fprintf(s.Log, "mms mcp: %s: %s\n", req.Method, rerr.Message)
		}
		return &response{JSONRPC: "2.0", ID: req.ID, Error: rerr}
	}
	return &response{JSONRPC: "2.0", ID: req.ID, Result: result}
}

func bytesTrim(b []byte) []byte {
	for len(b) > 0 && (b[len(b)-1] == '\n' || b[len(b)-1] == '\r' || b[len(b)-1] == ' ' || b[len(b)-1] == '\t') {
		b = b[:len(b)-1]
	}
	return b
}

func (s *Server) dispatch(ctx context.Context, method string, params json.RawMessage) (any, *rpcError) {
	switch method {
	case "initialize":
		var p struct {
			ProtocolVersion string `json:"protocolVersion"`
		}
		_ = json.Unmarshal(params, &p)
		version := protocolVersions[0]
		if slices.Contains(protocolVersions, p.ProtocolVersion) {
			version = p.ProtocolVersion
		}
		return map[string]any{
			"protocolVersion": version,
			"capabilities":    map[string]any{"tools": map[string]any{}, "resources": map[string]any{}},
			"serverInfo":      map[string]any{"name": "mms", "title": "Minecraft Metascript", "version": s.Version},
			"instructions":    instructions,
		}, nil
	case "ping":
		return map[string]any{}, nil
	case "tools/list":
		list := make([]map[string]any, 0, len(s.tools))
		for _, t := range s.tools {
			list = append(list, t.describe())
		}
		return map[string]any{"tools": list}, nil
	case "tools/call":
		var p struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, &rpcError{codeInvalidParams, err.Error()}
		}
		for _, t := range s.tools {
			if t.name == p.Name {
				return s.call(ctx, t, p.Arguments), nil
			}
		}
		return nil, &rpcError{codeInvalidParams, "unknown tool " + p.Name}
	case "resources/list":
		return map[string]any{"resources": resourceList()}, nil
	case "resources/templates/list":
		return map[string]any{"resourceTemplates": []any{}}, nil
	case "resources/read":
		var p struct {
			URI string `json:"uri"`
		}
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, &rpcError{codeInvalidParams, err.Error()}
		}
		text, ok := readResource(p.URI)
		if !ok {
			return nil, &rpcError{codeNotFound, "no resource " + p.URI}
		}
		return map[string]any{"contents": []map[string]any{{"uri": p.URI, "mimeType": resourceMIME(p.URI), "text": text}}}, nil
	}
	return nil, &rpcError{codeMethodNotFound, "unknown method " + method}
}

// call runs a tool. A tool's own failure is a result with isError set,
// which the model reads and can act on; it is not a protocol error.
func (s *Server) call(ctx context.Context, t *tool, args json.RawMessage) map[string]any {
	if len(args) == 0 || string(args) == "null" {
		args = json.RawMessage("{}")
	}
	text, err := func() (text string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("mms failed while running %s (please report this): %v", t.name, r)
			}
		}()
		return t.run(ctx, args)
	}()
	if err != nil {
		return map[string]any{"content": []map[string]any{{"type": "text", "text": err.Error()}}, "isError": true}
	}
	return map[string]any{"content": []map[string]any{{"type": "text", "text": clip(text)}}}
}

// maxResult bounds a tool's answer, which lands in the model's context.
const maxResult = 60_000

func clip(text string) string {
	if len(text) <= maxResult {
		return text
	}
	cut := maxResult
	for cut > 0 && text[cut]&0xC0 == 0x80 {
		cut--
	}
	return text[:cut] + fmt.Sprintf("\n… (%d more bytes left out; ask for less at a time)\n", len(text)-cut)
}

const instructions = `MMS (Minecraft Metascript) is a language that compiles to Minecraft Java Edition data packs and resource packs. Declarations such as LootTable, Recipe and Function become the JSON and .mcfunction files the game reads.

MMS is not in your training data, so read before you write: call docs with no arguments for the list of pages, and read guides/resources and guides/functions before writing a first file. Use lookup for the exact name and parameters of a kind, helper, Template or lint rule, and vanilla for the game's ids, instead of guessing them.

After every edit, call check and fix what it reports. Use explain to see the files a declaration compiles to, and test to run the project's Test declarations in the built-in interpreter. If you know the vanilla JSON for something, convert_json turns it into MMS.

Tools that read a project take either dir, a project folder on disk (the one with mms.json), or source, the text of one .mms file checked on its own.`
