package util

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	defaultSliceSize = 1024
)

// FileExists 检查文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetSortedFilenames 获取指定目录下，指定扩展名的所有文件
// 会返回 util.go 这样带basename和extname的文件名
func GetSortedFilenames(dir, ext string) ([]string, error) {
	filenames := make([]string, 0, defaultSliceSize)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ext) {
			filenames = append(filenames, info.Name())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(filenames)

	return filenames, nil
}
