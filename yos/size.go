package yos

import (
	"os"
	"path/filepath"
)

// GetFileSize returns the size in bytes for a regular file.
//
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
//
// If the given path is a symbolic link, it will be followed, but symbolic links inside the directory won't.
func GetDirSize(path string) (size int64, err error) {
	var (
		rootFi os.FileInfo
		root   string
	)
	if root, rootFi, err = resolveDirInfo(path); err == nil {
		err = filepath.Walk(root, func(itemPath string, itemFi os.FileInfo, errIn error) (errOut error) {
			errOut = errIn
			if os.SameFile(rootFi, itemFi) || errOut != nil {
				return
			}
			if isFileFi(&itemFi) || isSymlinkFi(&itemFi) {
				size += itemFi.Size()
			}
			return
		})
	} else {
		err = opError(opnSize, path, err)
	}
	return
}
