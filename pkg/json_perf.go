package pkg

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

/**
 *	Check for usage of "encoding/json" and suggest using "github.com/goccy/go-json"
 **/

var JsonPerf = &analysis.Analyzer{
	Name: "stdlibjson",
	Doc:  "Checks for encoding/json imports.",
	Run:  check_json,
}

const jsonShouldBe = "github.com/goccy/go-json"

func check_json(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {

		//	is this an import spec
		importSpec, ok := node.(*ast.ImportSpec)
		if !ok {
			return true
		}

		//	if so, what is its value?
		if importSpec.Path.Value == `"encoding/json"` {
			pass.Reportf(node.Pos(), "Import: %s should be %q", importSpec.Path.Value, jsonShouldBe)
			return false
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
