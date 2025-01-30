package pkg

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// AppErrs complains if you use "errors" from stdlib, rather than "app_errors" from go-common.
var AppErrs = &analysis.Analyzer{
	Name: "app_errors",
	Doc:  "Checks for usage of errors from stdlib.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		inspect := func(node ast.Node) bool {

			//	is this an import spec
			importSpec, ok := node.(*ast.ImportSpec)
			if !ok {
				return true
			}

			//	if so, what is its value?
			if importSpec.Path.Value == `"errors"` {
				pass.Reportf(node.Pos(), "Import: %s should be %q", importSpec.Path.Value, "app_errors")
				return false
			}
			return true
		}

		for _, f := range pass.Files {
			ast.Inspect(f, inspect)
		}
		return nil, nil
	},
}
