package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var LowercaseLetterRule = lowercaseLetterRule{name: "LowercaseLetterRule"}

type lowercaseLetterRule struct {
	name string
}

func (r lowercaseLetterRule) CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr) {
	basicLit := msg.(*ast.BasicLit) // we already checked that it's a string literal in extractMsgArgExpr

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	firstRune := rune(basicLit.Value[1])
	if !unicode.IsLetter(firstRune) {
		pass.Reportf(msg.Pos(), "log message should start with a letter")
		return
	}
	if !unicode.IsLower(firstRune) {
		pass.Reportf(msg.Pos(), "log message should start with an lowercase letter")
	}
}

func (r lowercaseLetterRule) Name() string {
	return r.name
}
