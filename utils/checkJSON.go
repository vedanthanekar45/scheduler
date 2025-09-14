package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckJSON(path string) {
	if strings.ToLower(filepath.Ext(path)) != ".json" {
		fmt.Printf("Error: Invalid file type. Only .json files are accepted.\n")
		fmt.Printf("Provided file: %s\n", path)
		os.Exit(1)
	}
}
