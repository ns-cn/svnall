package action

import (
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(sourcePath, destPath string) (err error) {
	source, err := os.ReadFile(sourcePath)
	if err != nil {
		return
	}
	_ = os.Remove(destPath)
	// 创建文件夹
	destDir, _ := filepath.Split(destPath)
	_, err = os.Stat(destDir)
	if err != nil {
		err = os.MkdirAll(destDir, 0777)
		if err != nil {
			return
		}
	}
	// 写入文件
	err = os.WriteFile(destPath, source, 0777)
	return
}

func GetFilePath(parent, file string) (targetPath string) {
	targetPath = strings.ReplaceAll(parent+file, "\\\\", "\\")
	targetPath = strings.ReplaceAll(targetPath, "//", "/")
	return
}
