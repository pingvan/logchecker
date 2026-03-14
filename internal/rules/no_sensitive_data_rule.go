package rules

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var NoSensitiveDataRule = noSensitiveDataRule{
	name:     "NoSensitiveDataRule",
	patterns: defaultSensitivePatterns,
}

var defaultSensitivePatterns = []string{
	"password",
	"passwd",
	"secret",
	"token",
	"api_key",
	"apikey",
	"auth",
	"credential",
	"private_key",
	"access_key",
	"jwt",
}

type noSensitiveDataRule struct {
	name     string
	patterns []string
}

func (r noSensitiveDataRule) CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr) {
	basicLit := msg.(*ast.BasicLit) // we already checked that it's a string literal in extractMsgArgExpr

	if len(basicLit.Value) < 2 {
		return // empty string or just quotes
	}

	msgStr := strings.ToLower(basicLit.Value[1 : len(basicLit.Value)-1])
	if pattern, found := r.containsSensitive(msgStr); found {
		pass.Reportf(msg.Pos(), "log message contains potentially sensitive data (pattern: %q)", pattern)
		return
	}
}

func (r noSensitiveDataRule) containsSensitive(s string) (string, bool) {
	for _, pattern := range r.patterns {
		if strings.Contains(s, pattern) {
			return pattern, true
		}
	}
	return "", false
}

func (r noSensitiveDataRule) Name() string {
	return r.name
}
