package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var NoSpecialCharsRule = noSpecialCharsRule{name: "NoSpecialCharsRule"}

var (
	forbiddenRanges = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 0x2600, Hi: 0x26FF, Stride: 1}, // Misc Symbols ☀☁☂
			{Lo: 0x2700, Hi: 0x27BF, Stride: 1}, // Dingbats ✂✈✉
		},
		R32: []unicode.Range32{
			{Lo: 0x1F300, Hi: 0x1F5FF, Stride: 1}, // Misc Symbols & Pictographs 🌍🔥
			{Lo: 0x1F600, Hi: 0x1F64F, Stride: 1}, // Emoticons 😀😂
			{Lo: 0x1F680, Hi: 0x1F6FF, Stride: 1}, // Transport & Map 🚀🚗
			{Lo: 0x1F900, Hi: 0x1F9FF, Stride: 1}, // Supplemental Symbols 🤖🧠
		},
	}

	forbiddenChars = map[rune]bool{
		'!': true,
		'?': true,
		';': true,
		':': true,
		',': true,
		'.': true,

		'(': true,
		')': true,
		'[': true,
		']': true,
		'{': true,
		'}': true,
		'<': true,
		'>': true,

		'+': true,
		'=': true,
		'*': true,
		'%': true,
		'^': true,
		'~': true,

		'"':  true,
		'\'': true,
		'`':  true,

		'|':  true,
		'\\': true,
		'/':  true,
		'@':  true,
		'#':  true,
		'$':  true,
		'&':  true,
	}
)

type noSpecialCharsRule struct {
	name string
}

func (r noSpecialCharsRule) CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr) {
	basicLit := msg.(*ast.BasicLit) // we already checked that it's a string literal in extractMsgArgExpr

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	msgStr := basicLit.Value[1 : len(basicLit.Value)-1]

	for _, ch := range msgStr {
		if forbiddenChars[ch] {
			pass.Reportf(msg.Pos(), "log message should not contain special characters")
			return
		}
		if unicode.Is(forbiddenRanges, ch) {
			pass.Reportf(msg.Pos(), "log message should not contain emoji")
			return
		}
	}
}

func (r noSpecialCharsRule) Name() string {
	return r.name
}
