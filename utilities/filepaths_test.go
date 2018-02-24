package utilities

import (
	"path/filepath"
	"testing"
)

func TestSwapExtensionToGo(t *testing.T) {
	tests := []struct {
		inPath          string
		expectedOutPath string
	}{
		{"test.gor", "test.go"},
		{filepath.Join("onedeep", "test.gor"), filepath.Join("onedeep", "test.go")},
		{filepath.Join("twodeep", "onedeep", "test.gor"), filepath.Join("twodeep", "onedeep", "test.go")},
	}

	for _, testcase := range tests {
		testPath := SwapExtension(testcase.inPath, "go")
		if testPath != testcase.expectedOutPath {
			t.Errorf("[FAIL] Returned path with %s inpath did not return expected outpath. Returned outpath: %s", testcase.inPath, testPath)
		}
	}
}

func TestSwapExtensionToGor(t *testing.T) {
	tests := []struct {
		inPath          string
		expectedOutPath string
	}{
		{"test.go", "test.gor"},
		{filepath.Join("onedeep", "test.go"), filepath.Join("onedeep", "test.gor")},
		{filepath.Join("twodeep", "onedeep", "test.go"), filepath.Join("twodeep", "onedeep", "test.gor")},
	}

	for _, testcase := range tests {
		testPath := SwapExtension(testcase.inPath, "gor")
		if testPath != testcase.expectedOutPath {
			t.Errorf("[FAIL] Returned path with %s in path did not return expected outpath. Returned outpath: %s", testcase.inPath, testPath)
		}
	}
}
