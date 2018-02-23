package commands

import (
	"testing"
	"path/filepath"
)

const (
	gotext = `package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`

	gortext = `#!/usr/bin/env gor
package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`
)

func TestGetGorPath(t *testing.T) {
	tests := []struct{
		inPath string
		expectedOutPath string
	}{
		{"test.go", "test.gor"},
		{filepath.Join("onedeep", "test.go"), filepath.Join("onedeep", "test.gor")},
		{filepath.Join("twodeep", "onedeep", "test.go"), filepath.Join("twodeep", "onedeep", "test.gor")},
	}

	for _, testcase := range tests {
		testPath := GetGorPath(testcase.inPath)
		if testPath != testcase.expectedOutPath {
			t.Errorf("[FAIL] Returned path with %s inpath did not return expected outpath. Returned outpath: %s", testcase.inPath, testPath)
		}
	}
}

func TestGetGorText(t *testing.T) {
	testbytes := ([]byte)(gotext)
	gorbytes := GetGorText(testbytes)
	if string(gorbytes) != gortext {
		t.Errorf("[FAIL] Returned gor file text did not match the expected text:\n%s", string(gorbytes))
	}
}
