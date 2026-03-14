package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var EnglishLanguageRule = englishLanguageRule{name: "EnglishLanguageRule"}

var allowedRanges = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 'A', Hi: 'Z', Stride: 1},
		{Lo: 'a', Hi: 'z', Stride: 1},
		{Lo: '0', Hi: '9', Stride: 1},
	},
}

type englishLanguageRule struct {
	name string
}

func (r englishLanguageRule) CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr) {
	basicLit := msg.(*ast.BasicLit) // we already checked that it's a string literal in extractMsgArgExpr

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	msgStr := basicLit.Value[1 : len(basicLit.Value)-1]
	for _, r := range msgStr {
		if unicode.In(r, allowedRanges) && !unicode.Is(unicode.Space, r) {
			pass.Reportf(msg.Pos(), "log message should contain only English letters, digits and spaces")
			return
		}
	}
}

func (r englishLanguageRule) Name() string {
	return r.name
}
