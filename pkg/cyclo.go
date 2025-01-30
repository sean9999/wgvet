package pkg

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

/**
 *	check for cyclomatic complexity
 *	prior art: https://github.com/bkielbasa/cyclop/blob/v1.2.3/pkg/analyzer/analyzer.go
 **/

const skipTests = true

// CycloMaxComplexity is the maximum cyclomatic complexity allowed before we start complaining
const CycloMaxComplexity = 25
const packageAverage = CycloMaxComplexity / 2

// CyclopsAnalzyer analyzes cyclomatic complexity and complains if it finds anything above [CycloMaxComplexity]
func CyclopsAnalzyer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "cyclop",
		Doc:  "checks function and package cyclomatic complexity",
		Run:  cyclo,
	}
}

func cyclo(pass *analysis.Pass) (interface{}, error) {
	var sum, count float64
	var pkgName string
	var pkgPos token.Pos

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if !ok {
				if node == nil {
					return true
				}
				if file, ok := node.(*ast.File); ok {
					pkgName = file.Name.Name
					pkgPos = node.Pos()
				}
				// we check function by function
				return true
			}

			if skipTests && testFunc(funcDecl) {
				return true
			}

			count++
			comp := complexity(funcDecl)
			sum += float64(comp)
			if comp > CycloMaxComplexity {
				pass.Reportf(node.Pos(), "calculated cyclomatic complexity for function %s is %d, max is %d", funcDecl.Name.Name, comp, CycloMaxComplexity)
			}

			return true
		})
	}

	if packageAverage > 0 {
		avg := sum / count
		if avg > packageAverage {
			pass.Reportf(pkgPos, "the average complexity for the package %s is %f, max is %f", pkgName, avg, packageAverage)
		}
	}

	return nil, nil
}

func testFunc(f *ast.FuncDecl) bool {
	return strings.HasPrefix(f.Name.Name, "Test")
}

func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{}
	ast.Walk(&v, fn)
	return v.Complexity
}

type complexityVisitor struct {
	Complexity int
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.Complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.Complexity++
		}
	}
	return v
}
