package commands

import (
	"bytes"

	"github.com/go-rillas/gorx/utilities"
)

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

func GetGoPath(prePath string) string {
	return utilities.SwapExtension(prePath, "go")
}
