package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/pingvan/logchecker/internal/logchecker"
	"github.com/pingvan/logchecker/internal/rules"
)

func TestEnglishLanguageRule(t *testing.T) {
	analyzer := logchecker.NewAnalyzer(rules.EnglishLanguageRule)
	analysistest.Run(t, testDataDir(t), analyzer, "english")
}
