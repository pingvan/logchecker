// pkg/rules/sensitive_test.go
package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/pingvan/logchecker/internal/logchecker"
	"github.com/pingvan/logchecker/internal/rules"
)

func TestNoSensitiveDataRule(t *testing.T) {
	analyzer := logchecker.NewAnalyzer(rules.NoSensitiveDataRule)
	analysistest.Run(t, testDataDir(t), analyzer, "sensitive")
}
