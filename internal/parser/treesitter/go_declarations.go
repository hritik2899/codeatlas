package treesitter

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/hritik2899/codeatlas/internal/model"
)

// parseGoDeclarations extracts top-level declarations from a Go source file.
// The Go standard parser is used for the first concrete language path; the
// resulting declarations still flow through the same language-independent
// representation used by the Tree-sitter adapter.
func parseGoDeclarations(source []byte) ([]Declaration, error) {
	file, err := parser.ParseFile(token.NewFileSet(), "source.go", source, 0)
	if err != nil {
		return nil, err
	}

	declarations := make([]Declaration, 0, len(file.Decls))
	for _, decl := range file.Decls {
		switch node := decl.(type) {
		case *ast.FuncDecl:
			kind := model.Function
			if node.Recv != nil && len(node.Recv.List) > 0 {
				kind = model.Method
			}
			declarations = append(declarations, Declaration{
				Name:      node.Name.Name,
				Kind:      kind,
				StartLine: file.Line(node.Pos()),
				EndLine:   file.Line(node.End()),
			})
		case *ast.GenDecl:
			for _, spec := range node.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				declarations = append(declarations, Declaration{
					Name:      ts.Name.Name,
					Kind:      model.Class,
					StartLine: file.Line(ts.Pos()),
					EndLine:   file.Line(ts.End()),
				})
			}
		}
	}
	return declarations, nil
}
