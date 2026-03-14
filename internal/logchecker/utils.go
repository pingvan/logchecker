package logchecker

import (
	"go/ast"
	"go/token"
)

func extractMsgArgExpr(call *ast.CallExpr) (ast.Expr, bool) {
	if len(call.Args) == 0 {
		return nil, false
	}

	msgArg := call.Args[0]

	basicLit, ok := msgArg.(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return nil, false
	}

	return msgArg, true
}

func extractArgunets(call *ast.CallExpr) ([]ast.Expr, bool) {
	if len(call.Args) < 2 {
		return nil, false
	}

	return call.Args[1:], true
}
