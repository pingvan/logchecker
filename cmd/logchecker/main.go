package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/pingvan/logchecker/internal/logchecker"
)

func main() {
	singlechecker.Main(logchecker.NewAnalyzer())
}
