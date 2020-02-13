package yos

import (
	"os"
	"path/filepath"
)

// JoinPath joins any number of path elements into a single path, adding a separator if necessary.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// ChangeExeDir changes the current working directory to the directory of the executable that started the current process.
//
// If a symbolic link is used to start the process, it will be changed to the directory of the executable that the link pointed to.
func ChangeExeDir() (err error) {
	var (
		ap string
		fi os.FileInfo
	)
	// get the path for the executable that started the current process
	if ap, err = os.Executable(); err != nil {
		err = opError(opnChange, ap, err)
		return
	}

	// get the file info of the executable and resolve the path if it's a symbolic link
	if fi, err = os.Lstat(ap); err == nil && isSymlinkFi(&fi) {
		ap, err = filepath.EvalSymlinks(ap)
	}
	if err != nil {
		err = opError(opnChange, ap, err)
		return
	}

	// get the executable directory and changes the current working directory to it
	if err = os.Chdir(filepath.Dir(ap)); err != nil {
		err = opError(opnChange, ap, err)
	}
	return
}

// MakeDir creates a directory named path with 0755 permission bits, along with any necessary parents.
//
// 0755 permission bits indicates that the owner can read, write and execute, whereas everyone else can read and execute but not modify.
//
// If the path is already a directory, MakeDir does nothing and returns nil.
func MakeDir(path string) (err error) {
	if err = os.MkdirAll(path, defaultDirectoryPermMode); err != nil {
		err = opError(opnMake, path, err)
	}
	return
}
