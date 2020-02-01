package yos

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/1set/gut/ystring"
)

var (
	// ErrShortRead means a read accepted fewer bytes than expected.
	ErrShortRead = errors.New("short read")
	// ErrEmptyPath means the given path is empty or blank.
	ErrEmptyPath = errors.New("path is empty or blank")
	// ErrSameFile means the given two files are actually the same one.
	ErrSameFile = errors.New("files are identical")
	// ErrNotRegular means the file is not a regular file.
	ErrNotRegular = errors.New("file is not regular")
	// ErrIsDir indicates the given path is actually a directory
	ErrIsDir = errors.New("target is a directory")
	// ErrIsFile indicates the given path is actually a file
	ErrIsFile = errors.New("target is a file")
	// ErrIsNotSymlink indicates the given path is not a symbolic link
	ErrIsNotSymlink = errors.New("target is not a symbolic link")
)

var (
	errInvalidPath    = errors.New("invalid path")
	errSameFile       = errors.New("files are identical")
	errShortRead      = errors.New("short read")
	errNotDirectory   = errors.New("not a directory")
	errNotRegularFile = errors.New("not a regular file")
	errNotSymlink     = errors.New("not a symbolic link")
)

// operation names for the Op field of os.PathError.
var (
	opnCompare = "compare"
	opnCopy    = "copy"
	opnMove    = "move"
)

// opError returns error struct with given details.
func opError(op, path string, err error) *os.PathError {
	return &os.PathError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}

// refineOpPaths validates, cleans up and adjusts the source and destination paths for operations like copy or move.
func refineOpPaths(opName, srcRaw, destRaw string, followLink bool) (src, dest string, err error) {
	if ystring.IsBlank(srcRaw) {
		err = opError(opName, srcRaw, errInvalidPath)
		return
	}
	if ystring.IsBlank(destRaw) {
		err = opError(opName, destRaw, errInvalidPath)
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
		err = opError(opName, srcRaw, err)
		return
	}

	// check if destination exists
	if destInfo, err = statFunc(destRaw); err != nil {
		// check existence of parent of the missing destination
		if os.IsNotExist(err) {
			if _, err = os.Stat(filepath.Dir(destRaw)); err != nil {
				err = opError(opName, destRaw, err)
			}
		} else {
			err = opError(opName, destRaw, err)
		}
	} else {
		if os.SameFile(srcInfo, destInfo) {
			err = opError(opName, destRaw, errSameFile)
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

// refineComparePaths validates, cleans up for file comparison.
func refineComparePaths(pathRaw1, pathRaw2 string) (path1, path2 string, err error) {
	if ystring.IsBlank(pathRaw1) {
		err = opError(opnCompare, pathRaw1, errInvalidPath)
		return
	}
	if ystring.IsBlank(pathRaw2) {
		err = opError(opnCompare, pathRaw2, errInvalidPath)
		return
	}

	// clean up paths
	path1, path2 = filepath.Clean(pathRaw1), filepath.Clean(pathRaw2)
	return
}
