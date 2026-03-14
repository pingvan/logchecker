package rules

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var NoSensitiveDataRule = noSensitiveDataRule{
	name:     "NoSensitiveDataRule",
	patterns: defaultSensitivePatterns,
}

func NewNoSensitiveDataRule(extraPatterns []string) Rule {
	patterns := make([]string, 0, len(defaultSensitivePatterns)+len(extraPatterns))
	patterns = append(patterns, defaultSensitivePatterns...)
	patterns = append(patterns, extraPatterns...)
	return &noSensitiveDataRule{
		name:     "NoSensitiveDataRule",
		patterns: patterns,
	}
}

var defaultSensitivePatterns = []string{
	"password",
	"passwd",
	"secret",
	"token",
	"api_key",
	"apikey",
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
	if pattern, found := r.containsSensitiveExpr(msg); found {
		pass.Reportf(msg.Pos(), "log message contains potentially sensitive data (pattern: %q)", pattern)
		return
	}

	for _, arg := range args {
		keyExpr, pattern, found := r.findSensitiveFieldKey(arg)
		if found {
			pass.Reportf(keyExpr.Pos(), "log field key contains potentially sensitive data (pattern: %q)", pattern)
			return
		}
	}
}

func (r noSensitiveDataRule) containsSensitiveExpr(expr ast.Expr) (string, bool) {
	switch expr := expr.(type) {
	case *ast.BasicLit:
		if expr.Kind != token.STRING || len(expr.Value) < 2 {
			return "", false
		}

		msgStr := strings.ToLower(expr.Value[1 : len(expr.Value)-1])
		return r.containsSensitive(msgStr)
	case *ast.BinaryExpr:
		if expr.Op != token.ADD {
			return "", false
		}
		if pattern, found := r.containsSensitiveExpr(expr.X); found {
			return pattern, true
		}
		return r.containsSensitiveExpr(expr.Y)
	case *ast.ParenExpr:
		return r.containsSensitiveExpr(expr.X)
	default:
		return "", false
	}
}

func (r noSensitiveDataRule) findSensitiveFieldKey(expr ast.Expr) (ast.Expr, string, bool) {
	call, ok := expr.(*ast.CallExpr)
	if !ok || len(call.Args) == 0 {
		return nil, "", false
	}

	pattern, found := r.containsSensitiveExpr(call.Args[0])
	if !found {
		return nil, "", false
	}

	return call.Args[0], pattern, true
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
