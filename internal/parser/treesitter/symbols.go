package treesitter

import (
	"strings"

	"github.com/hritik2899/codeatlas/internal/model"
)

// symbolNode turns a parser-discovered declaration into the canonical graph
// representation. Keeping this conversion here prevents language-specific AST
// nodes from leaking into the rest of CodeAtlas.
func symbolNode(kind model.NodeKind, name, path, language string, start, end int) model.Node {
	id := strings.Join([]string{string(kind), path, name}, ":")
	return model.Node{
		ID:       id,
		Kind:     kind,
		Name:     name,
		Path:     path,
		Language: language,
		Position: model.Position{StartLine: start, EndLine: end},
	}
}

// symbolsFromNames is intentionally parser-agnostic. The concrete AST walker
// will supply declarations in source order; this helper owns their conversion
// into stable CodeAtlas nodes.
func symbolsFromNames(names []string, path, language string) []model.Node {
	nodes := make([]model.Node, 0, len(names))
	for _, name := range names {
		if name = strings.TrimSpace(name); name == "" {
			continue
		}
		nodes = append(nodes, symbolNode(model.Function, name, path, language, 0, 0))
	}
	return nodes
}
