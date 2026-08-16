// Package treesitter contains the concrete source-parser integration used by
// CodeAtlas. The adapter is intentionally thin: language-specific AST details
// stay here and the rest of the ingestion pipeline consumes parser.Result.
package treesitter

import (
	"context"
	"fmt"

	"github.com/hritik2899/codeatlas/internal/model"
	"github.com/hritik2899/codeatlas/internal/parser"
)

// Adapter is the first concrete Tree-sitter integration point.
type Adapter struct {
	language string
}

func New(language string) *Adapter {
	return &Adapter{language: language}
}

func (a *Adapter) Parse(ctx context.Context, source []byte, path string) (parser.Result, error) {
	if err := ctx.Err(); err != nil {
		return parser.Result{}, err
	}
	if len(source) == 0 {
		return parser.Result{}, fmt.Errorf("cannot parse empty source file %q", path)
	}
	if a.language == "" {
		return parser.Result{}, fmt.Errorf("tree-sitter language is required")
	}

	// First concrete extraction primitive: an accepted source file is represented
	// as a CodeAtlas file node. Symbol extraction will build on this result.
	return parser.Result{
		Nodes: []model.Node{{
			ID:       "file:" + path,
			Kind:     model.File,
			Name:     path,
			Path:     path,
			Language: a.language,
		}},
	}, nil
}
