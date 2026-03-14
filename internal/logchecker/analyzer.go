package logchecker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/pingvan/logchecker/internal/rules"
)

const (
	name = "logchecker"
	doc  = "Checks slog and zap logging calls for correct usage."
)

var logcheckerAnalyzer = &analysis.Analyzer{
	Name: name,
	Doc:  doc,
	Run:  run,
	// using it because of standart practise: this analyzer once build AST which will be reused by other analyzer
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

// now using new is ambigious bute will need it later, when config using will be implimented
func NewAnalyzer() *analysis.Analyzer {
	return logcheckerAnalyzer
}

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if !checkLoggerSupported(pass, call) {
			return
		}

		for _, rule := range rules.AllRules {
			rule.CheckRule(pass, call)
		}
	})
	return nil, nil
}
