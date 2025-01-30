package pkg

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

/**
 *	Check for usage of "encoding/json" and suggest using "github.com/goccy/go-json"
 **/

// JsonPerf complains if you're not using [PreferredJson] as a drop-in replacement for [encoding/json]
var JsonPerf = &analysis.Analyzer{
	Name: "stdlibjson",
	Doc:  "Checks for encoding/json imports.",
	Run:  check_json,
}

// PreferredJson is the [encoding/json] drop-in replacement package you need to prefer
const PreferredJson = "github.com/goccy/go-json"

func check_json(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {

		//	is this an import spec
		importSpec, ok := node.(*ast.ImportSpec)
		if !ok {
			return true
		}

		//	if so, what is its value?
		if importSpec.Path.Value == `"encoding/json"` {
			pass.Reportf(node.Pos(), "Import: %s should be %q", importSpec.Path.Value, PreferredJson)
			return false
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
