package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/pingvan/logchecker/internal/logchecker"
	"github.com/pingvan/logchecker/internal/rules"
)

func TestNoSpecialCharsRule(t *testing.T) {
	analyzer := logchecker.NewAnalyzer(rules.NoSpecialCharsRule)
	analysistest.Run(t, testDataDir(t), analyzer, "special_chars")
}
