package yos

import (
	"os"
	"path/filepath"
)

type FilePathInfo struct {
	Path string
	Info *os.FileInfo
}

// listCondEntries returns a list of conditional directory entries.
func listCondEntries(root string, cond func(os.FileInfo) bool) (entries []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if cond(info) {
			entries = append(entries, &FilePathInfo{
				Path: path,
				Info: &info,
			})
		}
		return nil
	})
	return
}

func ListAll(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) bool { return true })
}

func ListFile(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) bool { return !info.IsDir() })
}

func ListDir(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) bool { return info.IsDir() })
}
