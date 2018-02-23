package utilities

import (
	"path/filepath"
	"strings"
)

func SwapExtension(inPath string, replaceExtension string) string {
	dirPath := filepath.Dir(inPath)
	basePath := filepath.Base(inPath)
	baseList := strings.Split(basePath, ".")
	baseName := baseList[0]
	finalBaseName := baseName + "." + replaceExtension

	return filepath.Join(dirPath, finalBaseName)
}
