package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // to be able to run this like "go run main.go -- input.go"
			continue
		}

		node, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}

		fmt.Println("Imports:")
		for _, i := range node.Imports {
			fmt.Println(i.Path.Value)
		}

		//ast.Walk(&v, node)
	}
}

type visitor struct {
	fset *token.FileSet
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {

	// fmt.Println("Imports:")
	// for _, i := range node.Imports {
	// 	fmt.Println(i.Path.Value)
	// }

	funcDecl, ok := node.(*ast.FuncDecl)
	if !ok {
		return v
	}

	params := funcDecl.Type.Params.List
	if len(params) != 2 { // [0] must be format (string), [1] must be args (...interface{})
		return v
	}

	firstParamType, ok := params[0].Type.(*ast.Ident)
	if !ok { // first param type isn't identificator so it can't be of type "string"
		return v
	}

	if firstParamType.Name != "string" { // first param (format) type is not string
		return v
	}

	secondParamType, ok := params[1].Type.(*ast.Ellipsis)
	if !ok { // args are not ellipsis (...args)
		return v
	}

	elementType, ok := secondParamType.Elt.(*ast.InterfaceType)
	if !ok { // args are not of interface type, but we need interface{}
		return v
	}

	if elementType.Methods != nil && len(elementType.Methods.List) != 0 {
		return v // has >= 1 method in interface, but we need an empty interface "interface{}"
	}

	if strings.HasSuffix(funcDecl.Name.Name, "f") {
		return v
	}

	fmt.Printf("%s: printf-like formatting function '%s' should be named '%sf'\n",
		v.fset.Position(node.Pos()), funcDecl.Name.Name, funcDecl.Name.Name)
	return v
}

// type v struct {
// 	info *types.Info
// }

// func (v v) viz(node ast.Node) (w ast.Visitor) {
// 	switch node := node.(type) {
// 	case *ast.CallExpr:
// 		// Get some kind of *ast.Ident for the CallExpr that represents the
// 		// package. Then we can look it up in v.info. Where exactly it sits in
// 		// the ast depends on the form of the function call.

// 		switch node := node.Fun.(type) {
// 		case *ast.SelectorExpr: // foo.ReadFile
// 			pkgID := node.X.(*ast.Ident)
// 			fmt.Println(v.info.Uses[pkgID].(*types.PkgName).Imported().Path())

// 		case *ast.Ident: // ReadFile
// 			pkgID := node
// 			fmt.Println(v.info.Uses[pkgID].Pkg().Path())
// 		}
// 	}

// 	return v
// }
