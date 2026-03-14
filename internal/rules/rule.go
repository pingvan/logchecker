package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Name() string
	CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr)
}
