package yos

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"

	"github.com/1set/gut/ystring"
)

var (
	errInvalidPath    = errors.New("invalid path")
	errSameFile       = errors.New("files are identical")
	errShortRead      = errors.New("short read")
	errIsDirectory    = errors.New("is a directory")
	errNotDirectory   = errors.New("not a directory")
	errNotRegularFile = errors.New("not a regular file")
	errNotSymlink     = errors.New("not a symbolic link")
)

// operation names for the Op field of os.PathError.
var (
	opnCompare = "compare"
	opnCopy    = "copy"
	opnMove    = "move"
	opnList    = "list"
	opnSize    = "size"
)

// internal use
var (
	emptyStr = ""
)

// underlyingError returns the underlying error for known os error types. forked from: os/error.go
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.LinkError:
		return err.Err
	case *os.PathError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	}
	return err
}

// opError returns error struct with given details.
func opError(op, path string, err error) *os.PathError {
	return &os.PathError{
		Op:   op,
		Path: path,
		Err:  underlyingError(err),
	}
}

type (
	funcStatFileInfo  func(name string) (os.FileInfo, error)
	funcCheckFileInfo func(fi *os.FileInfo) bool
	funcRemoveEntry   func(path string) error
	funcCopyEntry     func(src, dest string) error
)

// isFileFi indicates whether the FileInfo is for a regular file.
func isFileFi(fi *os.FileInfo) bool {
	return fi != nil && (*fi).Mode().IsRegular()
}

// isDirFi indicates whether the FileInfo is for a directory.
func isDirFi(fi *os.FileInfo) bool {
	return fi != nil && (*fi).Mode().IsDir()
}

// isSymlinkFi indicates whether the FileInfo is for a symbolic link.
func isSymlinkFi(fi *os.FileInfo) bool {
	return fi != nil && ((*fi).Mode()&os.ModeType == os.ModeSymlink)
}

func isLinkErrorCrossDevice(err error) bool {
	lerr, ok := err.(*os.LinkError)
	return ok && lerr.Err == syscall.EXDEV
}

func isLinkErrorNotDirectory(err error) bool {
	lerr, ok := err.(*os.LinkError)
	return ok && lerr.Err == syscall.ENOTDIR
}

// refineOpPaths validates, cleans up and adjusts the source and destination paths for operations like copy or move.
func refineOpPaths(opName, srcRaw, destRaw string, followLink bool) (src, dest string, err error) {
	// validate paths, and quit if got error
	if ystring.IsBlank(srcRaw) {
		err = opError(opName, srcRaw, errInvalidPath)
	} else if ystring.IsBlank(destRaw) {
		err = opError(opName, destRaw, errInvalidPath)
	}
	if err != nil {
		return
	}

	// clean up paths
	src, dest = filepath.Clean(srcRaw), filepath.Clean(destRaw)

	// use os.Stat to follow symbolic links
	statFunc := os.Lstat
	if followLink {
		statFunc = os.Stat
	}

	// check if source exists
	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = statFunc(src); err != nil {
		return
	}

	// check if destination exists
	if destInfo, err = statFunc(dest); err != nil {
		// check existence of parent of the missing destination
		if os.IsNotExist(err) {
			_, err = os.Stat(filepath.Dir(dest))
		}
	} else {
		if os.SameFile(srcInfo, destInfo) {
			err = opError(opName, dest, errSameFile)
		} else if destInfo.IsDir() {
			// append file name of source to path of the existing destination
			dest = JoinPath(dest, srcInfo.Name())
		}
	}
	return
}

// refineComparePaths validates, cleans up path for file comparison.
func refineComparePaths(pathRaw1, pathRaw2 string) (path1, path2 string, err error) {
	// validate paths
	if ystring.IsBlank(pathRaw1) {
		err = opError(opnCompare, pathRaw1, errInvalidPath)
	} else if ystring.IsBlank(pathRaw2) {
		err = opError(opnCompare, pathRaw2, errInvalidPath)
	}

	// clean up paths
	if err == nil {
		path1, path2 = filepath.Clean(pathRaw1), filepath.Clean(pathRaw2)
	}
	return
}

// resolveDirInfo returns file info of a path if it's a directory or a symbolic link to a directory, otherwise returns an error.
func resolveDirInfo(pathRaw string) (path string, fi os.FileInfo, err error) {
	if fi, err = os.Lstat(pathRaw); err == nil {
		// resolve to real path if the given path is a symbolic link
		if isSymlinkFi(&fi) {
			if path, err = filepath.EvalSymlinks(pathRaw); err == nil {
				// update file info for the real path
				fi, err = os.Lstat(path)
			}
			if err != nil {
				path = emptyStr
				return
			}
		} else {
			// simply clean the path if the raw path isn't a symbolic link to resolve
			path = filepath.Clean(pathRaw)
		}

		// check if the final path is a directory
		if !isDirFi(&fi) {
			err, path = errNotDirectory, emptyStr
		}
	}
	return
}

// openFileInfo returns file descriptor and info of a path if it's a regular file, otherwise returns an error.
func openFileInfo(path string) (file *os.File, fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); err == nil {
		if isFileFi(&fi) {
			if file, err = os.Open(path); err == nil {
				return
			}
		} else {
			err = errNotRegularFile
		}
	}
	return
}
