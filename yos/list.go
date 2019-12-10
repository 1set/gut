package yos

import (
	"os"
	"path/filepath"
)

type FilePathInfo struct {
	Path string
	Info *os.FileInfo
}

// listCondItems returns a list of conditional directory entries.
func listCondItems(root string, cond func(os.FileInfo) bool) (items []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if cond(info) {
			items = append(items, &FilePathInfo{
				Path: path,
				Info: &info,
			})
		}
		return nil
	})
	return
}

func ListAll(root string) (items []*FilePathInfo, err error) {
	return listCondItems(root, func(info os.FileInfo) bool { return true })
}

func ListFile(root string) (items []*FilePathInfo, err error) {
	return listCondItems(root, func(info os.FileInfo) bool { return !info.IsDir() })
}

func ListDir(root string) (items []*FilePathInfo, err error) {
	return listCondItems(root, func(info os.FileInfo) bool { return info.IsDir() })
}
