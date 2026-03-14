package rules

import (
	"go/ast"
	"go/token"
	"strings"
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
		fixed := string(basicLit.Value[0]) + strings.ToLower(string(firstRune)) + basicLit.Value[2:]
		pass.Report(analysis.Diagnostic{
			Pos:     msg.Pos(),
			End:     msg.End(),
			Message: "log message must start with lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "lowercase the first letter",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     msg.Pos(),
							End:     msg.End(),
							NewText: []byte(fixed),
						},
					},
				},
			},
		})
	}
}

func (r lowercaseLetterRule) Name() string {
	return r.name
}
