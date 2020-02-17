package filepath

import (
	"os"
	"path/filepath"
)

// 获取当前二进制文件运行时的当前目录
func GetRunDir() (dir string, err error) {
	filePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir, err = filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return "", err
	}

	return dir, nil
}

// 获取当前所在的目录，这个建议开发时使用，尤其是go run这种
func GetCurrentPwd() (dir string, err error){
	return os.Getwd()
}
