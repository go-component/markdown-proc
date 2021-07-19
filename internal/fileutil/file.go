package fileutil

import (
	"os"
	"path/filepath"
	"regexp"
)

func MkdirAll(path string) error {
	err := os.MkdirAll(path, 0755)
	if !os.IsExist(err) {
		return err
	}
	return nil
}

func Ext(path string) string {

	reg := regexp.MustCompile("\\.\\w+")

	return reg.FindString(filepath.Ext(path))
}
