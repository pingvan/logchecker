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

type logCheckerAnalyzer struct {
	rules []rules.Rule
}

func NewAnalyzer(customRules ...rules.Rule) *analysis.Analyzer {
	rulesList := customRules
	if len(customRules) == 0 {
		rulesList = rules.AllRules
	}
	l := newLogCheckerAnalyzer(rulesList)
	a := &analysis.Analyzer{
		Name: name,
		Doc:  doc,
		Run:  l.run,
		// using it because of standart practise: this analyzer once build AST which will be reused by other analyzer
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
	return a
}

func newLogCheckerAnalyzer(ruleList []rules.Rule) *logCheckerAnalyzer {
	return &logCheckerAnalyzer{
		rules: ruleList,
	}
}

func (l *logCheckerAnalyzer) run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if !checkLoggerSupported(pass, call) {
			return
		}

		msgExpr, ok := extractMsgArgExpr(call)
		if !ok {
			return
		}

		args := extractArgunets(call)

		for _, rule := range l.rules {
			rule.CheckRule(pass, call, msgExpr, args)
		}
	})
	return nil, nil
}
