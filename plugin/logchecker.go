package plugin

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis"

	"github.com/pingvan/logchecker/internal/config"
	"github.com/pingvan/logchecker/internal/logchecker"
)

// New creates a new instance of the logchecker analyzer for golangci-lint module plugin system.
// https://golangci-lint.run/plugins/module-plugins/
func New(conf any) ([]*analysis.Analyzer, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("logchecker: %w", err)
	}

	cfg, err := config.Load(dir)
	if err != nil {
		return nil, fmt.Errorf("logchecker: %w", err)
	}

	return []*analysis.Analyzer{logchecker.NewAnalyzerFromConfig(cfg)}, nil
}
