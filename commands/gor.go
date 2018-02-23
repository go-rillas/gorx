package commands

import (
	"bytes"

	"github.com/go-rillas/gorx/utilities"
)

func GetGorText(inByteString []byte) []byte {
	if len(inByteString) == 0 {
		return []byte("")
	}

	var buf bytes.Buffer

	buf.Write([]byte("#!/usr/bin/env gor\n"))
	buf.Write(inByteString)
	return buf.Bytes()
}

func GetGorPath(prePath string) string {
	return utilities.SwapExtension(prePath, "gor")
}
