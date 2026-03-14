package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/pingvan/logchecker/internal/logchecker"
	"github.com/pingvan/logchecker/internal/rules"
)

func TestLowercaseRule(t *testing.T) {
	analyzer := logchecker.NewAnalyzer(rules.LowercaseLetterRule)
	analysistest.Run(t, testDataDir(t), analyzer, "lowercase")
}
