package conf

import (
	"path/filepath"
	"github.com/realgo/env"
)

func FileAbsPath(confRealPath string) string {
	return  filepath.Join(env.GetRootPath(), confRealPath)
}
