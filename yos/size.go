package yos

import (
	"os"
	"path/filepath"
)

// GetFileSize returns the size in bytes for a regular file.
// If the given path is a symbolic link, it will be followed.
func GetFileSize(path string) (size int64, err error) {
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

// GetSymlinkSize returns the size in bytes for a symbolic link.
func GetSymlinkSize(path string) (size int64, err error) {
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

// GetDirSize returns total size in bytes for all regular files and symbolic links in a directory.
// If the given path is a symbolic link, it will be followed, but symbolic links inside the directory won't.
func GetDirSize(path string) (size int64, err error) {
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
