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

func checkLoggerSupported(pass *analysis.Pass, call *ast.CallExpr) bool {
	return isSlogCall(pass, call) || isZapCall(pass, call)
}

func isSlogCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	obj, ok := pass.TypesInfo.Selections[sel]
	if ok {
		pkg := obj.Obj().Pkg()
		if pkg == nil {
			return false
		}
		return pkg.Path() == "log/slog"
	}

	objIdent := pass.TypesInfo.ObjectOf(sel.Sel)
	if objIdent == nil {
		return false
	}
	pkg := objIdent.Pkg()
	if pkg == nil {
		return false
	}
	return pkg.Path() == "log/slog"
}

func isZapCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	panic("not implemented")
}
