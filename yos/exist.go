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
func IsFileExist(path string) bool {
	return checkPathExist(path, os.Stat, isFileFi)
}

// IsDirExist checks if the specified path exists and is a directory.
// If the path is a symbolic link, it will attempt to follow the link and check.
func IsDirExist(path string) bool {
	return checkPathExist(path, os.Stat, isDirFi)
}

// IsSymlinkExist checks if the specified path exists and is a symbolic link.
// It only checks the path itself and makes no attempt to follow the link.
func IsSymlinkExist(path string) bool {
	return checkPathExist(path, os.Lstat, isSymlinkFi)
}

func checkPathExist(path string, stat funcStatFileInfo, check funcCheckFileInfo) bool {
	fi, err := stat(path)
	return err == nil && check(&fi)
}

// JoinPath joins any number of path elements into a single path, adding a separator if necessary.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
