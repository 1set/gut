package yos

import (
	"os"
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
	return moveEntry(
		src, dest,
		isFileFi, errNotRegularFile,
		os.Remove,
		func(src, dest string) error { return bufferCopyFile(src, dest, defaultBufferSize) })
}

// MoveSymlink moves a symbolic link to a target file.
// It makes no attempt to read the referenced file.
func MoveSymlink(src, dest string) (err error) {
	return moveEntry(
		src, dest,
		isSymlinkFi, errNotSymlink,
		os.Remove,
		func(src, dest string) error { return copySymlink(src, dest) })
}

// MoveDir moves a directory to a target directory recursively. Symbolic links inside the directories will not be followed.
//
// If the target is an existing file, it will be replaced by the source directory under the same target name.
//
// If the target is an existing directory, the source directory will be moved to the directory with the same name.
//
// If the target doesn't exist but its parent directory does, the source directory will be moved to the parent directory with the target name.
//
// It stops and returns immediately if any error occurs. ErrSameFile is returned if it detects an attempt to move a file to itself.
func MoveDir(src, dest string) (err error) {
	return moveEntry(
		src, dest,
		isDirFi, errNotDirectory,
		os.RemoveAll,
		func(src, dest string) error { return copyDir(src, dest) })
}

// moveEntry moves source to target by renaming or copying.
func moveEntry(src, dest string, check funcCheckFileInfo, errMode error, remove funcRemoveEntry, copy funcCopyEntry) (err error) {
	// validate and refine paths
	if src, dest, err = refineOpPaths(opnMove, src, dest, false); err != nil {
		return
	}

	// check if source exists and its file mode
	var srcInfo os.FileInfo
	if srcInfo, err = os.Lstat(src); err == nil && !check(&srcInfo) {
		err = opError(opnMove, src, errMode)
	}
	if err != nil {
		return
	}

	// attempts to move file by renaming links
	if err = os.Rename(src, dest); os.IsExist(err) || isLinkErrorNotDirectory(err) {
		// remove destination if fails for its existence or not directory
		_ = remove(dest)
		err = os.Rename(src, dest)
	}

	// if rename succeeds, or got unexpected errors
	switch {
	case err == nil:
		return
	case os.IsNotExist(err):
		err = opError(opnMove, src, err)
		return
	case os.IsExist(err):
		err = opError(opnMove, dest, err)
		return
	}

	// cross device: move == remove dest + copy to dest + remove src
	if isLinkErrorCrossDevice(err) {
		// remove destination file, and ignore the non-existence error
		if err = remove(dest); err != nil && !os.IsNotExist(err) {
			err = opError(opnMove, dest, err)
			return
		}
		if err = copy(src, dest); err == nil {
			err = remove(src)
		}
	}
	return
}
