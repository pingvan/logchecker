package rules_test

import (
	"path/filepath"
	"runtime"
	"testing"
)

func testDataDir(t *testing.T) string {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve test file path")
	}

	return filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
}
