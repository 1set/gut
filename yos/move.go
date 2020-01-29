package yos

import (
	"os"
	"syscall"
)

// MoveFile moves a file to a target file or directory. Symbolic links will be not be followed.
//
// If the target is an existing file, the target will be overwritten with the source file.
//
// If the target is an existing directory, the source file will be moved to the directory with the same file name.
//
// If the target doesn't exist but its parent directory does, the source file will be moved to the parent directory with the target name.
//
// ErrSameFile is returned if it detects an attempt to copy a file to itself.
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

	// cross device: move == remove dest + copy to dest + remove src
	if lerr, ok := err.(*os.LinkError); ok && lerr.Err == syscall.EXDEV {
		// remove destination file, and ignore the non-existence error
		if err = os.Remove(dest); err != nil && !os.IsNotExist(err) {
			return
		}
		if err = bufferCopyFile(src, dest, defaultBufferSize); err == nil {
			err = os.Remove(src)
		}
	}

	return
}
