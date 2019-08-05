package conf

import (
	"github.com/greatming/realgo/env"
	"path/filepath"
)

func FileAbsPath(confRealPath string) string {
	return filepath.Join(env.GetRootPath(), confRealPath)
}
