package commands

import (
	"bytes"

	"github.com/go-rillas/gorx/utilities"
)

// GetGorText returns a byte string for a go runner file from a valid Go source file byte string parameter
func GetGorText(inByteString []byte) []byte {
	if len(inByteString) == 0 {
		return []byte("")
	}

	var buf bytes.Buffer

	buf.Write([]byte("#!/usr/bin/env gor\n"))
	buf.Write(inByteString)
	return buf.Bytes()
}

// GetGorPath returns the prePath file path string with the existing file extension converted to .gor
func GetGorPath(prePath string) string {
	return utilities.SwapExtension(prePath, "gor")
}
