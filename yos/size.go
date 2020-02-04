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
	if fi, err = os.Lstat(path); err == nil {
		// resolve given symbolic link to real path
		if isSymlinkFi(&fi) {
			rawPath := path
			if path, err = filepath.EvalSymlinks(path); err == nil {
				fi, err = os.Lstat(path)
			} else {
				err = opError(opnSize, rawPath, err)
			}
			if err != nil {
				return
			}
		}

		if isDirFi(&fi) {
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
		} else {
			err = opError(opnSize, path, errNotDirectory)
		}
	}
	return
}
