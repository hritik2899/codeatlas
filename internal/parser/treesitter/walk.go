package treesitter

import (
	"context"
	"fmt"

	"github.com/hritik2899/codeatlas/internal/model"
	"github.com/hritik2899/codeatlas/internal/parser"
)

// WalkResult contains the declarations discovered while walking a parsed file.
// Keeping this result separate makes the AST walker independently testable.
type WalkResult struct {
	Declarations []Declaration
}

// Walk performs the parser-stage traversal contract. The concrete Tree-sitter
// node visitor will populate declarations as language grammars are enabled.
func Walk(ctx context.Context, source []byte) (WalkResult, error) {
	if err := ctx.Err(); err != nil {
		return WalkResult{}, err
	}
	if len(source) == 0 {
		return WalkResult{}, fmt.Errorf("cannot walk empty source")
	}

	return WalkResult{Declarations: nil}, nil
}

// resultFromWalk combines the file node with declarations discovered by the
// walker. This is the first explicit connection between traversal and the
// canonical parser result.
func resultFromWalk(path, language string, walkResult WalkResult) parser.Result {
	nodes := []model.Node{{
		ID:       "file:" + path,
		Kind:     model.File,
		Name:     path,
		Path:     path,
		Language: language,
	}}
	nodes = append(nodes, declarationsToNodes(walkResult.Declarations, path, language)...)

	return parser.Result{Nodes: nodes}
}
