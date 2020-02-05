package yos

import (
	"os"
	"path/filepath"
)

func IsSymlinkBroken(path string) (broken bool, err error) {
	var fi os.FileInfo
	if fi, err = os.Lstat(path); err == nil {
		if isSymlinkFi(&fi) {
			_, err = filepath.EvalSymlinks(path)
			broken = err != nil
		} else {
			err = opError(opnEmpty, path, errNotSymlink)
		}
	}
	return
}

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
			err = filepath.Walk(path, func(pathIn string, info os.FileInfo, errIn error) (errOut error) {
				errOut = errIn
				if path == pathIn || errOut != nil {
					return
				}
				errOut = errSameFile
				return
			})

			if err == nil {
				empty = true
			} else if err == errSameFile {
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
