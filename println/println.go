// Package println defines an Analyzer that reports println usages
package println

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "println",
	Doc:      "reports println usage",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ce, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		i, ok := ce.Fun.(*ast.Ident)
		if !ok {
			return
		}

		if i.String() != "println" {
			return
		}

		pass.Reportf(i.Pos(), "println usage found")
	})

	return nil, nil
}
