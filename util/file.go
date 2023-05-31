package util

import (
	"io/fs"
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
