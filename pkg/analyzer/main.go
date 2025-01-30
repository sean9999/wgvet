package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "stdlibjson",
	Doc:  "Checks for encoding/json imports.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {

		funcDecl, ok := node.(*ast.ImportSpec)
		if !ok {
			return true
		}

		if funcDecl.Path.Value == `"encoding/json"` {
			pass.Reportf(node.Pos(), "Import: %s\n", funcDecl.Path.Value)
			return false
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
