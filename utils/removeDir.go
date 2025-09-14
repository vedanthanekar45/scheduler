package utils

import (
	"path/filepath"
)

func RemoveDirectory(path string) string {
	scheduleName := filepath.Base(path)
	scheduleName = scheduleName[:len(scheduleName)-len(filepath.Ext(scheduleName))]
	return scheduleName
}
