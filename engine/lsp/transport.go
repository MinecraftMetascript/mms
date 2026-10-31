// Package lsp is the MMS language server: JSON-RPC 2.0 over a transport,
// mapping LSP requests onto the ide package. The same server runs over
// stdio for desktop editors and over postMessage in the browser (WASM).
package lsp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

// Transport carries whole JSON-RPC messages.
type Transport interface {
	// Read returns the next message, or io.EOF when the client is gone.
	Read() ([]byte, error)
	Write(msg []byte) error
}

// Stream frames messages with Content-Length headers, as LSP does over
// stdio.
type Stream struct {
	r  *bufio.Reader
	w  io.Writer
	mu sync.Mutex
}

// NewStream returns a Content-Length transport over r and w.
func NewStream(r io.Reader, w io.Writer) *Stream {
	return &Stream{r: bufio.NewReader(r), w: w}
}

func (s *Stream) Read() ([]byte, error) {
	length := -1
	for {
		line, err := s.r.ReadString('\n')
		if err != nil {
			if err == io.EOF && line == "" {
				return nil, io.EOF
			}
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			if length < 0 {
				continue
			}
			break
		}
		name, value, ok := strings.Cut(line, ":")
		if ok && strings.EqualFold(strings.TrimSpace(name), "Content-Length") {
			n, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil || n < 0 {
				return nil, fmt.Errorf("lsp: bad Content-Length %q", value)
			}
			length = n
		}
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(s.r, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (s *Stream) Write(msg []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, err := fmt.Fprintf(s.w, "Content-Length: %d\r\n\r\n", len(msg)); err != nil {
		return err
	}
	_, err := s.w.Write(msg)
	return err
}

// Pipe is an in-memory transport for tests and the browser: the client
// side sends with Send and receives with Recv. Send never blocks, so it is
// safe to call from a JavaScript callback.
type Pipe struct {
	mu    sync.Mutex
	queue [][]byte
	ready chan struct{} // signalled when queue becomes non-empty
	out   chan []byte
	done  chan struct{}
	once  sync.Once
}

// NewPipe returns a connected in-memory transport.
func NewPipe() *Pipe {
	return &Pipe{ready: make(chan struct{}, 1), out: make(chan []byte, 64), done: make(chan struct{})}
}

func (p *Pipe) Read() ([]byte, error) {
	for {
		p.mu.Lock()
		if len(p.queue) > 0 {
			m := p.queue[0]
			p.queue = p.queue[1:]
			p.mu.Unlock()
			return m, nil
		}
		p.mu.Unlock()
		select {
		case <-p.ready:
		case <-p.done:
			return nil, io.EOF
		}
	}
}

func (p *Pipe) Write(msg []byte) error {
	select {
	case p.out <- msg:
		return nil
	case <-p.done:
		return io.ErrClosedPipe
	}
}

// Send delivers a message to the server.
func (p *Pipe) Send(msg []byte) {
	p.mu.Lock()
	p.queue = append(p.queue, msg)
	p.mu.Unlock()
	select {
	case p.ready <- struct{}{}:
	default:
	}
}

// Recv returns the server's next message.
func (p *Pipe) Recv() ([]byte, bool) {
	select {
	case m := <-p.out:
		return m, true
	case <-p.done:
		return nil, false
	}
}

// Close ends the connection.
func (p *Pipe) Close() { p.once.Do(func() { close(p.done) }) }
