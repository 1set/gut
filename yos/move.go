package yos

import (
	"os"
)

func MoveFile(src, dest string) (err error) {
	if src, dest, err = refineCopyPaths(src, dest, false); err == nil {
		err = moveFile(src, dest)
	}
	return
}

func moveFile(src, dest string) (err error) {
	// check if source exists and is a file
	var srcInfo os.FileInfo
	if srcInfo, err = os.Lstat(src); err == nil && !srcInfo.Mode().IsRegular() {
		err = ErrNotRegular
	}
	if err != nil {
		return
	}

	// attempts to move file by renaming links
	if err = os.Rename(src, dest); os.IsExist(err) {
		// remove destination if fails for its existence
		os.Remove(dest)
		err = os.Rename(src, dest)
	}

	if err == nil || os.IsExist(err) || os.IsNotExist(err) {
		// if rename succeeds, or got unexpected errors
		return
	}

	// move == copy to dest + remove src
	if err = bufferCopyFile(src, dest, defaultBufferSize); err == nil {
		err = os.Remove(src)
	}
	return
}
