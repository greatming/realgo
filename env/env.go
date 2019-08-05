package env

import (
	"os"
)

var rootPath string

func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	}
	var err error
	rootPath, err = os.Getwd()
	if err != nil {
		panic("can't init root path ")
	}
	return rootPath
}
