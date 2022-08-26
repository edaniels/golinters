// Package printf defines an Analyzer that reports various printf usages
package fatal

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "fatal",
	Doc:      "reports a usage of logger.Fatal",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
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

		if i.String() != "logger.Fatal" {
			return
		}
		pass.Reportf(i.Pos(), "logger.Fatal usage found")
		return
	})

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
