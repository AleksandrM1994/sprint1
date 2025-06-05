package analyzers

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// OsExitAnalyzer - экземпляр объекта analysis.Analyzer, описывающий функцию поиска os.Exit в файлах main.go
var OsExitAnalyzer = &analysis.Analyzer{
	Name: "osExitAnalyzer",
	Doc:  "check for start os.Exit in package main",
	Run:  runOsExitAnalyzer,
}

// runOsExitAnalyzer - анализатор для поиска os.Exit в файлах main.go
func runOsExitAnalyzer(pass *analysis.Pass) (interface{}, error) {
	fset := pass.Fset
	for _, file := range pass.Files {
		// функцией ast.Inspect проходим по всем узлам AST
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.CallExpr:
				expr, okExpr := x.Fun.(*ast.SelectorExpr)
				if okExpr {
					if ident, okIdent := expr.X.(*ast.Ident); okIdent && ident.Name == "os" && expr.Sel.Name == "Exit" {
						pos := fset.Position(ident.Pos())
						if strings.Contains(pos.Filename, "main.go") {
							pass.Reportf(ident.NamePos, "dont use os.Exit in main file")
						}
					}
				}

			}
			return true
		})
	}
	return nil, nil
}
