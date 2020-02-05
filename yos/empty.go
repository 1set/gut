package yos

import (
	"os"
	"path/filepath"
)

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

func IsDirEmpty(path string) (empty bool, err error) {
	var (
		fi  os.FileInfo
		raw = path
	)
	if path, fi, err = resolveDirInfo(path); err == nil {
		if isDirFi(&fi) {
			err = filepath.Walk(path, func(itemPath string, info os.FileInfo, errItem error) error {
				if path == itemPath || errItem != nil {
					return errItem
				}
				return errStepOutDir
			})

			if err == nil {
				empty = true
			} else if err == errStepOutDir {
				err = nil
			}
		} else {
			err = opError(opnEmpty, path, errNotDirectory)
		}
	} else {
		err = opError(opnEmpty, raw, err)
	}
	return
}
