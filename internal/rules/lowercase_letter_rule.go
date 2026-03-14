package rules

import (
	"go/ast"
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// LowercaseLetterRule checks that log messages start with a lowercase letter.
var LowercaseLetterRule = lowercaseLetterRule{name: "LowercaseLetterRule"}

type lowercaseLetterRule struct {
	name string
}

func (r lowercaseLetterRule) CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr) {
	basicLit, ok := msg.(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return
	}

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	firstRune := rune(basicLit.Value[1])
	if !unicode.IsLetter(firstRune) {
		pass.Reportf(msg.Pos(), "log message should start with a letter")
		return
	}
	if !unicode.IsLower(firstRune) {
		pass.Reportf(msg.Pos(), "log message must start with lowercase")
	}
}

func (r lowercaseLetterRule) Name() string {
	return r.name
}
