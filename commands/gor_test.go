package commands

import (
	"path/filepath"
	"testing"
)

const (
	gotext_gor = `package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`

	gortext_gor = `#!/usr/bin/env gor
package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`
)

func TestGetGorPath(t *testing.T) {
	tests := []struct {
		inPath          string
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
	testbytes := ([]byte)(gotext_gor)
	gorbytes := GetGorText(testbytes)
	if string(gorbytes) != gortext_gor {
		t.Errorf("[FAIL] Returned gor file text did not match the expected text:\n%s", string(gorbytes))
	}

	emptygorbytes := GetGorText([]byte(""))
	if len(emptygorbytes) != 0 {
		t.Errorf("[FAIL] Empty byte string returned text and should not have returned text: %v", emptygorbytes)
	}
}
