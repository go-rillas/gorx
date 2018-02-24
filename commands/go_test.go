package commands

import (
	"path/filepath"
	"testing"
)

const (
	gotext_go = `package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`

	gortext_go = `#!/usr/bin/env gor
package example

import "fmt"

func AnExample() {
	fmt.Println("Hello world!")
}
`
)

func TestGetGoPath(t *testing.T) {
	tests := []struct {
		inPath          string
		expectedOutPath string
	}{
		{"test.gor", "test.go"},
		{filepath.Join("onedeep", "test.gor"), filepath.Join("onedeep", "test.go")},
		{filepath.Join("twodeep", "onedeep", "test.gor"), filepath.Join("twodeep", "onedeep", "test.go")},
	}

	for _, testcase := range tests {
		testPath := GetGoPath(testcase.inPath)
		if testPath != testcase.expectedOutPath {
			t.Errorf("[FAIL] Returned path with %s inpath did not return expected outpath. Returned outpath: %s", testcase.inPath, testPath)
		}
	}
}

func TestGetGoText(t *testing.T) {
	testbytes := ([]byte)(gortext_go)
	gobytes := GetGoText(testbytes)
	if string(gobytes) != gotext_go {
		t.Errorf("[FAIL] Returned gor file text did not match the expected text:\n%s", string(gobytes))
	}

	// test that valid Go source returns the same valid Go source
	testbytesGoText := ([]byte)(gotext_go)
	gobytesFromGo := GetGoText(testbytesGoText)
	if string(gobytesFromGo) != gotext_go {
		t.Errorf("[FAIL] Expected return of the same Go source when Go source is passed to function.  Did not receive expected text.")
	}

	emptygobytes := GetGoText([]byte(""))
	if len(emptygobytes) != 0 {
		t.Errorf("[FAIL] Empty byte string returned text and should not have returned text: %v", emptygobytes)
	}
}
