package yos

import (
	"os"
	"path/filepath"
)

// IsFileEmpty checks whether the given file is empty.
func IsFileEmpty(path string) (empty bool, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(path); err == nil {
		if isFileFi(&fi) {
			empty = fi.Size() == 0
		} else {
			err = opError(opnEmpty, path, errNotRegularFile)
		}
	}
	return
}

// IsDirEmpty checks whether the given directory contains nothing.
func IsDirEmpty(path string) (empty bool, err error) {
	var (
		rootFi os.FileInfo
		root   string
	)
	if root, rootFi, err = resolveDirInfo(path); err == nil {
		err = filepath.Walk(root, func(itemPath string, itemFi os.FileInfo, errItem error) error {
			if os.SameFile(rootFi, itemFi) || errItem != nil {
				return errItem
			}
			// force exit for the first entry other than the root itself
			return errStepOutDir
		})

		if err == nil {
			empty = true
		} else if err == errStepOutDir {
			err = nil
		}
	} else {
		err = opError(opnEmpty, path, err)
	}
	return
}
