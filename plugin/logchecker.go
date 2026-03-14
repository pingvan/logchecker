package plugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/pingvan/logchecker/internal/logchecker"
)

// New creates a new instance of the logchecker analyzer for golangci-lint module plugin system.
// https://golangci-lint.run/plugins/module-plugins/
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{logchecker.NewAnalyzer()}, nil
}
