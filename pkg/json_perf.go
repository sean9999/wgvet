package pkg

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var JsonPerf = &analysis.Analyzer{
	Name: "stdlibjson",
	Doc:  "Checks for encoding/json imports.",
	Run:  run,
}

const shouldBe = "github.com/goccy/go-json"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {

		//	is this an import spec
		importSpec, ok := node.(*ast.ImportSpec)
		if !ok {
			return true
		}

		//	if so, what is its value?
		if importSpec.Path.Value == `"encoding/json"` {
			pass.Reportf(node.Pos(), "Import: %s should be %q", importSpec.Path.Value, shouldBe)
			return false
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
