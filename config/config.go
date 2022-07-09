package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsPathOfExecutable() string {
	// 実行ファイルの絶対パスを取得
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 取得したパスのディレクトリを取得
	executableDir := filepath.Dir(executablePath)

	return executableDir
}
