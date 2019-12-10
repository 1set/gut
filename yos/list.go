package yos

import (
	"os"
	"path/filepath"
)

type FilePathInfo struct {
	Path string
	Info *os.FileInfo
}

func ListAll(root string) (items []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		items = append(items, &FilePathInfo{
			Path: path,
			Info: &info,
		})
		return nil
	})
	return
}

func ListFile(root string) (items []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// only save files
		if !info.IsDir() {
			items = append(items, &FilePathInfo{
				Path: path,
				Info: &info,
			})
		}
		return nil
	})
	return
}

func ListDir(root string) (items []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// only save dirs
		if info.IsDir() {
			items = append(items, &FilePathInfo{
				Path: path,
				Info: &info,
			})
		}
		return nil
	})
	return
}
