package yos

import (
	"os"
)

// Exist checks whether the given path exists.
//
// If the file is a symbolic link, it will attempt to follow the link and check if the source file exists.
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// NotExist checks whether the given path doesn't exist.
//
// If the file is a symbolic link, it will attempt to follow the link and check if the source file doesn't exist.
func NotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// ExistFile checks whether the specified path exists and is a file.
//
// If the path is a symbolic link, it will attempt to follow the link and check.
func ExistFile(path string) bool {
	return checkPathExist(path, os.Stat, isFileFi)
}

// ExistDir checks whether the specified path exists and is a directory.
//
// If the path is a symbolic link, it will attempt to follow the link and check.
func ExistDir(path string) bool {
	return checkPathExist(path, os.Stat, isDirFi)
}

// ExistSymlink checks whether the specified path exists and is a symbolic link.
//
// It only checks the path itself and makes no attempt to follow the link.
func ExistSymlink(path string) bool {
	return checkPathExist(path, os.Lstat, isSymlinkFi)
}

func checkPathExist(path string, stat funcStatFileInfo, check funcCheckFileInfo) bool {
	fi, err := stat(path)
	return err == nil && check(&fi)
}
