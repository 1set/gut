package yos

import (
	"os"
	"path/filepath"
)

// IsExist checks if the file, directory exists.
// If the file is a symbolic link, it will attempt to follow the link and check if the source file exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsNotExist checks if the file, directory doesn't exist.
// If the file is a symbolic link, it will attempt to follow the link and check if the source file doesn't exist.
func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// IsFileExist checks if the specified path exists and is a file.
// If the path is a symbolic link, it will attempt to follow the link and check.
func IsFileExist(path string) (exist bool, err error) {
	return checkPathExist(path, false, errNotRegularFile)
}

// IsDirExist checks if the specified path exists and is a directory.
// If the path is a symbolic link, it will attempt to follow the link and check.
func IsDirExist(path string) (exist bool, err error) {
	return checkPathExist(path, true, errNotDirectory)
}

// FIXME: check according to func var
func checkPathExist(path string, expectDir bool, fallbackErr error) (exist bool, err error) {
	exist, err = false, nil
	var info os.FileInfo
	if info, err = os.Stat(path); err == nil {
		if info.IsDir() == expectDir {
			exist = true
		} else {
			err = fallbackErr
		}
	}
	return
}

// IsSymlinkExist checks if the specified path exists and is a symbolic link.
// It only checks the path itself and makes no attempt to follow the link.
func IsSymlinkExist(path string) (exist bool, err error) {
	exist, err = false, nil
	var info os.FileInfo
	if info, err = os.Lstat(path); err == nil {
		if (info.Mode() & os.ModeSymlink) != 0 {
			exist = true
		} else {
			err = errNotSymlink
		}
	}
	return
}

// JoinPath joins any number of path elements into a single path, adding a separator if necessary.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
