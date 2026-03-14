package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/pingvan/logchecker/internal/config"
	"github.com/pingvan/logchecker/internal/logchecker"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "logchecker: %v\n", err)
		os.Exit(1)
	}

	cfg, err := config.Load(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "logchecker: %v\n", err)
		os.Exit(1)
	}

	singlechecker.Main(logchecker.NewAnalyzerFromConfig(cfg))
}
