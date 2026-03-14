package logchecker

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const (
	slogPkgPath = "log/slog"

	zapPkgPath    = "go.uber.org/zap"
	zapObjectName = "Logger"
)

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
		return pkg.Path() == slogPkgPath
	}

	objIdent := pass.TypesInfo.ObjectOf(sel.Sel)
	if objIdent == nil {
		return false
	}
	pkg := objIdent.Pkg()
	if pkg == nil {
		return false
	}
	return pkg.Path() == slogPkgPath
}

func isZapCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	typ := pass.TypesInfo.TypeOf(sel.X)
	if typ == nil {
		return false
	}

	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}

	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}

	pkg := named.Obj().Pkg()
	if pkg == nil {
		return false
	}

	return pkg.Path() == zapPkgPath &&
		named.Obj().Name() == zapObjectName
}
