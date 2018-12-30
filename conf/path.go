package conf

import (
	"path/filepath"
	"realgo/env"
)

func FileAbsPath(confRealPath string) string {
	return  filepath.Join(env.GetRootPath(), confRealPath)
}
