package treesitter

import (
	"strings"

	"github.com/hritik2899/codeatlas/internal/model"
)

// Declaration describes the small amount of information the AST walker needs
// to hand to the canonical model. The actual Tree-sitter node is deliberately
// not exposed beyond this package.
type Declaration struct {
	Name      string
	Kind      model.NodeKind
	StartLine int
	EndLine   int
}

// declarationsToNodes converts source-order declarations into stable graph
// nodes. This is the first step toward replacing parser-specific AST objects
// with CodeAtlas primitives.
func declarationsToNodes(declarations []Declaration, path, language string) []model.Node {
	nodes := make([]model.Node, 0, len(declarations))
	for _, declaration := range declarations {
		name := strings.TrimSpace(declaration.Name)
		if name == "" {
			continue
		}

		nodes = append(nodes, symbolNode(
			declaration.Kind,
			name,
			path,
			language,
			declaration.StartLine,
			declaration.EndLine,
		))
	}
	return nodes
}
