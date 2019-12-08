package yos

import (
	"os"
	"path/filepath"
)

// IsExist checks if the file, directory or symbolic link exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsExist checks if the file, directory or symbolic link doesn't exist.
func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func IsFileExist(path string) (exist bool, err error) {
	panic("Todo")
}

func IsDirExist(path string) (exist bool, err error) {
	panic("Todo")
}

func IsSymlinkExist(path string) (exist bool, err error) {
	panic("Todo if symblink is not handled well by IsFileExist / IsDirExist")
}

func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
