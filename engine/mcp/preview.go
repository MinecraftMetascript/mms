package mcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// The preview tools (biomes, density, jigsaw) run the editors' own
// preview code, which is TypeScript: the worldgen sampling is Deepslate's,
// and the jigsaw simulator lives in packages/preview. Rather than keep a
// second copy in Go, the engine embeds that code as one JavaScript file
// (built by packages/preview/scripts/build-agent.mjs) and runs it with a
// JavaScript runtime on the machine. A pure-Go interpreter was tried and
// was about 150 times slower than Node, too slow to sample a map.
//
// The engine sends one request as JSON on the runtime's stdin: the tool,
// its arguments, the project's data pack files, its structure files'
// jigsaw blocks, and the vanilla worldgen files previews fall back on.
// It reads back {text} or {error}.

//go:embed preview/agent.mjs
var agentJS []byte

// previewRequest is what the JavaScript side reads (agent/common.ts).
type previewRequest struct {
	Tool    string           `json:"tool"`
	Args    json.RawMessage  `json:"args"`
	Files   []map[string]any `json:"files"`
	Assets  []map[string]any `json:"assets,omitempty"`
	Vanilla spec.Vanilla     `json:"vanilla,omitempty"`
}

type previewResponse struct {
	Text  string `json:"text"`
	Error string `json:"error"`
}

// vanillaWorldgen are the vanilla folders the preview worker reads.
var vanillaWorldgen = []string{"worldgen/noise", "worldgen/density_function", "worldgen/noise_settings", "worldgen/material_rule", "worldgen/material_condition", "tags/worldgen/biome"}

// errNoRuntime explains how to get the preview tools working.
var errNoRuntime = errors.New("the preview tools run the editors' preview code, which needs a JavaScript runtime: install Node.js 20 or later (or Bun or Deno) and make sure `node` is on PATH, or set MMS_JS to the runtime's path")

// jsRuntime finds the command that runs a JavaScript file: MMS_JS, then
// node, bun and deno on PATH.
func jsRuntime() (string, []string, error) {
	if p := os.Getenv("MMS_JS"); p != "" {
		return p, argsFor(p), nil
	}
	for _, name := range []string{"node", "bun", "deno"} {
		if p, err := exec.LookPath(name); err == nil {
			return p, argsFor(p), nil
		}
	}
	return "", nil, errNoRuntime
}

// PreviewRuntime says what runs the preview tools (biomes, density,
// jigsaw), for `mms mcp -check`: the runtime's path, or why there is none.
func PreviewRuntime() (string, error) {
	bin, _, err := jsRuntime()
	return bin, err
}

// argsFor returns the arguments a runtime needs before the script's path.
func argsFor(path string) []string {
	base := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".exe")
	if base == "deno" {
		return []string{"run", "--quiet", "--allow-read", "--allow-env"}
	}
	return nil
}

var (
	agentOnce sync.Once
	agentPath string
	agentErr  error
)

// agentFile writes the embedded script to the user's cache folder once,
// named by its hash, so every version of mms keeps its own copy.
func agentFile() (string, error) {
	agentOnce.Do(func() {
		dir, err := os.UserCacheDir()
		if err != nil {
			dir = os.TempDir()
		}
		sum := sha256.Sum256(agentJS)
		dir = filepath.Join(dir, "mms")
		agentPath = filepath.Join(dir, "preview-"+hex.EncodeToString(sum[:8])+".mjs")
		if b, err := os.ReadFile(agentPath); err == nil && bytes.Equal(b, agentJS) {
			return
		}
		if agentErr = os.MkdirAll(dir, 0o755); agentErr != nil {
			return
		}
		tmp := fmt.Sprintf("%s.%d.tmp", agentPath, os.Getpid())
		if agentErr = os.WriteFile(tmp, agentJS, 0o644); agentErr != nil {
			return
		}
		agentErr = os.Rename(tmp, agentPath)
	})
	return agentPath, agentErr
}

// preview runs one preview tool on a project's output.
func (s *Server) preview(ctx context.Context, tool string, args json.RawMessage, worldgen bool) (string, error) {
	var t target
	// The JavaScript side checks the tool's own arguments; here only
	// where the project is matters.
	var all map[string]json.RawMessage
	if err := json.Unmarshal(args, &all); err != nil {
		return "", fmt.Errorf("bad arguments: %w", err)
	}
	if v, ok := all["dir"]; ok {
		if err := json.Unmarshal(v, &t.Dir); err != nil {
			return "", fmt.Errorf("bad arguments: dir: %w", err)
		}
		delete(all, "dir")
	}
	if v, ok := all["source"]; ok {
		if err := json.Unmarshal(v, &t.Source); err != nil {
			return "", fmt.Errorf("bad arguments: source: %w", err)
		}
		delete(all, "source")
	}
	rest, _ := json.Marshal(all)

	p, r, err := s.load(ctx, t)
	if err != nil {
		return "", err
	}
	req := previewRequest{Tool: tool, Args: rest, Files: []map[string]any{}}
	for _, f := range r.Emit() {
		if f.Pack == sema.ResourcePack {
			continue
		}
		if f.Asset != "" {
			if info, ok := project.StructureAsset(p.Files, f.Path, f.Asset); ok {
				info["path"], info["src"] = f.Path, f.Asset
				req.Assets = append(req.Assets, info)
			}
			continue
		}
		req.Files = append(req.Files, map[string]any{"path": f.Path, "text": string(f.Data)})
	}
	if worldgen {
		if v := s.pack.Vanilla(); v != nil {
			req.Vanilla = spec.Vanilla{}
			for _, folder := range vanillaWorldgen {
				req.Vanilla[folder] = v[folder]
			}
		}
	}

	res, err := runAgent(ctx, req)
	if err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", errors.New(res.Error)
	}
	if r.HasErrors() {
		return "The project has errors, so its output may be missing parts; run check.\n\n" + res.Text, nil
	}
	return res.Text, nil
}

// runAgent runs the embedded script on one request.
func runAgent(ctx context.Context, req previewRequest) (previewResponse, error) {
	bin, pre, err := jsRuntime()
	if err != nil {
		return previewResponse{}, err
	}
	script, err := agentFile()
	if err != nil {
		return previewResponse{}, fmt.Errorf("writing the preview script: %w", err)
	}
	in, err := json.Marshal(req)
	if err != nil {
		return previewResponse{}, err
	}
	cmd := exec.CommandContext(ctx, bin, append(pre, script)...)
	cmd.Stdin = bytes.NewReader(in)
	var out, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &stderr
	cmd.Env = append(os.Environ(), "NODE_NO_WARNINGS=1")
	runErr := cmd.Run()
	var res previewResponse
	if err := json.Unmarshal(out.Bytes(), &res); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if runErr != nil && msg == "" {
			msg = runErr.Error()
		}
		if msg == "" {
			msg = "no answer"
		}
		return previewResponse{}, fmt.Errorf("the preview script failed under %s: %s", bin, lastLines(msg, 8))
	}
	return res, nil
}

// lastLines keeps the end of a long error, where the cause usually is.
func lastLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}
