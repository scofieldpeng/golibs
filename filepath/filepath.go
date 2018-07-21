package filepath

import (
	"path/filepath"
	"os"
)

// 获取程序运行时的当前目录
func GetRunDir() (path string, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return dir, nil
}
