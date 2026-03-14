package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Rule defines the interface for a log message checking rule.
type Rule interface {
	Name() string
	CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr)
}
