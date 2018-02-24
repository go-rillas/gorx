package commands

import (
	"bytes"

	"github.com/go-rillas/gorx/utilities"
)

// GetGoText returns a byte string for a Go source file byte string from a go runner file byte string parameter
func GetGoText(inByteString []byte) []byte {
	if len(inByteString) == 0 {
		return []byte("")
	}

	var outByteString []byte
	if bytes.HasPrefix(inByteString, []byte("#!")) {
		preFileList := bytes.SplitN(inByteString, []byte("\n"), 2)
		outByteString = preFileList[1]
	} else {
		outByteString = inByteString
	}

	return outByteString
}

// GetGoPath returns the prePath file path string with the existing file extension converted to .go
func GetGoPath(prePath string) string {
	return utilities.SwapExtension(prePath, "go")
}
