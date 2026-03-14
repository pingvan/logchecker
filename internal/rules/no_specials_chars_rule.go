package rules

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// NoSpecialCharsRule checks that log messages do not contain special characters or emojis.
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
	basicLit, ok := msg.(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return
	}

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	msgStr := basicLit.Value[1 : len(basicLit.Value)-1]

	if !containsSpecialChars(msgStr) {
		return
	}

	cleaned := removeSpecialChars(msgStr)
	quote := string(basicLit.Value[0])
	fixed := quote + cleaned + quote

	pass.Report(analysis.Diagnostic{
		Pos:     msg.Pos(),
		End:     msg.End(),
		Message: "log message should not contain special characters",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "remove special characters",
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

func containsSpecialChars(s string) bool {
	for _, ch := range s {
		if forbiddenChars[ch] || unicode.Is(forbiddenRanges, ch) {
			return true
		}
	}
	return false
}

func removeSpecialChars(s string) string {
	var b strings.Builder
	for _, ch := range s {
		if !forbiddenChars[ch] && !unicode.Is(forbiddenRanges, ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func (r noSpecialCharsRule) Name() string {
	return r.name
}
