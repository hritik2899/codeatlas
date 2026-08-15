// Package parser defines the boundary between source-language parsers and the
// language-independent CodeAtlas graph model.
package parser

import (
	"context"
	"fmt"

	"github.com/hritik2899/codeatlas/internal/model"
)

// Parser converts source files into CodeAtlas nodes and edges.
//
// The interface intentionally knows nothing about Tree-sitter, Go, Java, or
// any other parsing implementation. This keeps the graph pipeline stable as
// language support grows.
type Parser interface {
	Parse(ctx context.Context, source []byte, path string) (Result, error)
}

// Result is the parser's language-independent output.
type Result struct {
	Nodes []model.Node
	Edges []model.Edge
}

// Registry selects a parser from a file extension.
type Registry struct {
	parsers map[string]Parser
}

func NewRegistry() *Registry {
	return &Registry{parsers: make(map[string]Parser)}
}

func (r *Registry) Register(extension string, p Parser) {
	r.parsers[extension] = p
}

func (r *Registry) ForExtension(extension string) (Parser, error) {
	p, ok := r.parsers[extension]
	if !ok {
		return nil, fmt.Errorf("no parser registered for extension %q", extension)
	}
	return p, nil
}
