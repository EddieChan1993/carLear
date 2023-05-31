package util

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
)

// DirFiles 扫描路径下所有文件
func DirFiles(path string) []string {
	res := make([]string, 0, 1000)
	//i := 100
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		//if i <= 0 {
		//	return err
		//}
		if info.IsDir() {
			return nil
		}
		res = append(res, path)
		//i--
		return nil
	})
	return res
}

func NewBufferFile(name string) (file *os.File, buf *bufio.Writer, err error) {
	fileHandle, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, nil, err
	}
	buf = bufio.NewWriter(fileHandle)
	return fileHandle, buf, nil
}

func IsExtraFile(filePath string) bool {
	// 文件不存在则返回error
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
