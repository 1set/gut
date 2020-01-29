package yos

import (
	"os"
	"path/filepath"

	"github.com/1set/gut/ystring"
)

// refineOpPaths validates, cleans up and adjusts the source and destination paths for operations like copy or move.
func refineOpPaths(srcRaw, destRaw string, followLink bool) (src, dest string, err error) {
	if ystring.IsBlank(srcRaw) || ystring.IsBlank(destRaw) {
		err = ErrEmptyPath
		return
	}

	// clean up paths
	srcRaw, destRaw = filepath.Clean(srcRaw), filepath.Clean(destRaw)

	// use os.Lstat instead if not following symbolic links
	statFunc := os.Stat
	if !followLink {
		statFunc = os.Lstat
	}

	// check if source exists
	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = statFunc(srcRaw); err != nil {
		return
	}

	// check if destination exists
	if destInfo, err = statFunc(destRaw); err != nil {
		// check existence of parent of the missing destination
		if os.IsNotExist(err) {
			_, err = os.Stat(filepath.Dir(destRaw))
		}
	} else {
		if os.SameFile(srcInfo, destInfo) {
			err = ErrSameFile
		} else if destInfo.IsDir() {
			// append file name of source to path of the existing destination
			destRaw = JoinPath(destRaw, srcInfo.Name())
		}
	}

	if err == nil {
		src, dest = srcRaw, destRaw
	}
	return
}
