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

// Adapter is the first concrete source parser. Language-specific extraction is
// delegated to small helpers while the returned shape remains parser.Result.
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

	walkResult, err := Walk(ctx, source)
	if err != nil {
		return parser.Result{}, err
	}

	if a.language == string(LanguageGo) {
		declarations, err := parseGoDeclarations(source)
		if err != nil {
			return parser.Result{}, fmt.Errorf("parse Go source %q: %w", path, err)
		}
		walkResult.Declarations = declarations
	}

	result := resultFromWalk(path, a.language, walkResult)
	if len(result.Nodes) == 0 {
		result.Nodes = []model.Node{{
			ID:       "file:" + path,
			Kind:     model.File,
			Name:     path,
			Path:     path,
			Language: a.language,
		}}
	}
	return result, nil
}
