package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsPathOfExecutable() string {
	relativePath := filepath.Dir(os.Args[0])
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return absPath
}
