package logchecker

import (
	"go/ast"
)

func extractMsgArgExpr(call *ast.CallExpr) (ast.Expr, bool) {
	if len(call.Args) == 0 {
		return nil, false
	}

	return call.Args[0], true
}

func extractArgunets(call *ast.CallExpr) []ast.Expr {
	if len(call.Args) < 2 {
		return nil
	}

	return call.Args[1:]
}
