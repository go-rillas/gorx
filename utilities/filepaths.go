package utilities

import (
	"path/filepath"
	"strings"
)

func SwapExtension(inPath string, replaceExtension string) string {
	absPath, _ := filepath.Abs(inPath)
	dirPath := filepath.Dir(absPath)
	basePath := filepath.Base(absPath)
	baseList := strings.Split(basePath, ".")
	baseName := baseList[0]
	finalBaseName := baseName + "." + replaceExtension

	return filepath.Join(dirPath, finalBaseName)
}
