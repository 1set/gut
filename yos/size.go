package yos

import (
	"os"
	"path/filepath"
)

func FileSize(path string) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(path); err == nil {
		if isFileFi(&fi) {
			size = fi.Size()
		} else {
			err = opError(opnSize, path, errNotRegularFile)
		}
	}
	return
}

func SymlinkSize(path string) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = os.Lstat(path); err == nil {
		if isSymlinkFi(&fi) {
			size = fi.Size()
		} else {
			err = opError(opnSize, path, errNotSymlink)
		}
	}
	return
}

func DirSize(path string) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(path); err == nil {
		if isDirFi(&fi) {
			//size = fi.Size()

			err = filepath.Walk(path, func(pathIn string, info os.FileInfo, errIn error) (errOut error) {
				errOut = errIn
				if path == pathIn || errOut != nil {
					return
				}
				if isFileFi(&info) || isSymlinkFi(&info) {
					size += info.Size()
				}
				return
			})

			//listCondEntries(path, func(info os.FileInfo) (bool, error) {
			//	if isFileFi(&info) || isSymlinkFi(&info) {
			//		size += info.Size()
			//	}
			//	return true, nil
			//})
		} else {
			err = opError(opnSize, path, errNotDirectory)
		}
	}
	return
}
