package yos

import (
	"os"
	"path/filepath"
	"strings"
)

type FilePathInfo struct {
	Path string
	Info os.FileInfo
}

// listCondEntries returns a list of conditional directory entries.
func listCondEntries(root string, cond func(os.FileInfo) (bool, error)) (entries []*FilePathInfo, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var ok bool
		if ok, err = cond(info); ok {
			entries = append(entries, &FilePathInfo{
				Path: path,
				Info: info,
			})
		}
		return err
	})
	return
}

// ListAll returns a list of all directory entries in the given directory in lexical order, including root.
// It searches recursively, but symbolic links will be not be followed.
func ListAll(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return true, nil })
}

// ListFile returns a list of file directory entries in the given directory in lexical order.
// It searches recursively, but symbolic links will be not be followed.
func ListFile(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return !info.IsDir(), nil })
}

// ListDir returns a list of nested directory entries in the given directory in lexical order, including root.
// It searches recursively, but symbolic links will be not be followed.
func ListDir(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return info.IsDir(), nil })
}

const (
	ListRecursive int = 1 << iota
	ListToLower
	ListIncludeFile
	ListIncludeDir
)

func ListMatch(root string, flag int, patterns ...string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		fileName := info.Name()
		if flag&ListToLower != 0 {
			fileName = strings.ToLower(fileName)
		}
		isDir := info.IsDir()
		if (isDir && (flag&ListIncludeDir != 0)) || (!isDir && (flag&ListIncludeFile != 0)) {
			for _, pattern := range patterns {
				ok, err = filepath.Match(pattern, fileName)
				if ok || err != nil {
					break
				}
			}
		}
		if err == nil && isDir && (flag&ListRecursive == 0) {
			err = filepath.SkipDir
		}
		return
	})
}

func ListMatchAll(root string, patterns ...string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		fileName := info.Name()
		for _, pattern := range patterns {
			ok, err = filepath.Match(pattern, fileName)
			if ok || err != nil {
				break
			}
		}
		return
	})
}

func ListMatchFile(root string, patterns ...string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		if info.IsDir() {
			// ignore directories
			return
		}
		fileName := info.Name()
		for _, pattern := range patterns {
			ok, err = filepath.Match(pattern, fileName)
			if ok || err != nil {
				break
			}
		}
		return
	})
}

func ListMatchDir(root string, patterns ...string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		if !info.IsDir() {
			// ignore files
			return
		}
		fileName := info.Name()
		for _, pattern := range patterns {
			ok, err = filepath.Match(pattern, fileName)
			if ok || err != nil {
				break
			}
		}
		return
	})
}
