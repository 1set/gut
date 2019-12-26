package yos

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/bits"
	"os"
	"path/filepath"

	"github.com/1set/gut/ystring"
)

var (
	// ErrShortRead means a read accepted fewer bytes than expected.
	ErrShortRead = errors.New("short read")
	// ErrEmptyPath means the given path is empty or blank.
	ErrEmptyPath = errors.New("path is empty")
	// ErrSameFile means the given two files are actually the same one.
	ErrSameFile = errors.New("files are identical")
	// ErrNotRegular means the file is not a regular file.
	ErrNotRegular = errors.New("file is not regular")
)

const (
	defaultDirectoryFileMode = os.FileMode(0755)
	defaultBufferSize        = 256 * 1024
)

// CopyFile copies a file to a target file or directory. Symbolic links are followed.
// If the target is an existing file, the target will be overwritten with the source file.
// If the target is an existing directory, the source file will be copied to the directory with the same file name.
// If the target doesn't exist but its parent directory does, the source file will be copied to the parent directory with the target name.
// ErrSameFile is returned if it detects an attempt to copy a file to itself.
func CopyFile(src, dest string) (err error) {
	if src, dest, err = refinePaths(src, dest); err == nil {
		err = bufferCopyFile(src, dest, defaultBufferSize)
	}
	return
}

func bufferCopyFile(src, dest string, bufferSize int64) (err error) {
	var srcFile, destFile *os.File
	if srcFile, err = os.Open(src); err != nil {
		return
	}
	defer srcFile.Close()

	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = os.Stat(src); err == nil && !srcInfo.Mode().IsRegular() {
		err = ErrNotRegular
	}
	if err != nil {
		return
	}

	// check if source and destination files are identical
	if destInfo, err = os.Stat(dest); err == nil {
		if !destInfo.Mode().IsRegular() {
			err = ErrNotRegular
		} else if os.SameFile(srcInfo, destInfo) {
			err = ErrSameFile
		}
	} else if os.IsNotExist(err) {
		// it's okay if destination file doesn't exist
		err = nil
	}

	if err != nil {
		return
	}

	// use smaller buffer if source file is not big enough
	fileSize := srcInfo.Size()
	if bufferSize > fileSize {
		bufferSize = 1 << uint(bits.Len64(uint64(fileSize)))
	}

	if destFile, err = os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, srcInfo.Mode()); err != nil {
		return
	}
	defer func() {
		if fe := destFile.Close(); fe != nil {
			err = fe
		}
	}()

	var nr, nw int
	buf := make([]byte, bufferSize)
	for {
		if nr, err = srcFile.Read(buf); err != nil {
			if err == io.EOF && nr > 0 {
				err = io.ErrUnexpectedEOF
			}
			break
		} else if nr == 0 {
			break
		}

		if nw, err = destFile.Write(buf[:nr]); err != nil {
			break
		} else if nw != nr {
			err = io.ErrShortWrite
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	// err = destFile.Sync()
	return
}

func refinePaths(srcRaw, destRaw string) (src, dest string, err error) {
	if ystring.IsBlank(srcRaw) || ystring.IsBlank(destRaw) {
		err = ErrEmptyPath
		return
	}

	// clean up paths
	srcRaw, destRaw = filepath.Clean(srcRaw), filepath.Clean(destRaw)

	// check if source exists
	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = os.Stat(srcRaw); err != nil {
		return
	}

	// check if destination exists
	if destInfo, err = os.Stat(destRaw); err != nil {
		// check existence of parent of the missing destination
		if os.IsNotExist(err) {
			_, err = os.Stat(filepath.Dir(destRaw))
		}
	} else {
		// append file name of source to path of the existing destination
		if destInfo.IsDir() {
			destRaw = JoinPath(destRaw, srcInfo.Name())
		}
	}

	if err == nil {
		src, dest = srcRaw, destRaw
	}
	return
}

func copySymlink(src, dest string) (err error) {
	var link string
	if link, err = os.Readlink(src); err == nil {
		err = os.Symlink(link, dest)
	}
	return
}

func CopyDir(src, dest string) (err error) {
	/*
		if dest dir exist, append src name to dest dir, and copy.
		if dest dir not exist, check if the parent exists.
		what if cp -R src1 dest --- the parent dir is .
		stop if any error occurs
	*/

	if src, dest, err = refinePaths(src, dest); err == nil {
		err = copyDir(src, dest)
	}
	return
}

// copyDir copies all entries of source directory to destination directory recursively.
func copyDir(src, dest string) (err error) {
	var srcInfo, destInfo os.FileInfo

	// check if source exists and is a directory
	if srcInfo, err = os.Stat(src); err == nil && !srcInfo.IsDir() {
		err = fmt.Errorf("%v: source is not a directory", src)
	}
	if err != nil {
		return
	}

	// check if destination doesn't exist or is not a file or source itself
	if destInfo, err = os.Stat(dest); err == nil {
		if destInfo.Mode().IsRegular() {
			err = fmt.Errorf("%v: destination is a regular file", src)
		} else if os.SameFile(srcInfo, destInfo) {
			err = ErrSameFile
		}
	} else if os.IsNotExist(err) {
		err = nil
		if err = os.MkdirAll(dest, defaultDirectoryFileMode); err == nil {
			originMode := srcInfo.Mode()
			defer os.Chmod(dest, originMode)
		}
	}
	if err != nil {
		return
	}

	// loop through entries in source directory
	var entries []os.FileInfo
	if entries, err = ioutil.ReadDir(src); err != nil {
		return
	}

IterateEntry:
	for _, entry := range entries {
		srcPath, destPath := JoinPath(src, entry.Name()), JoinPath(dest, entry.Name())

		switch entry.Mode() & os.ModeType {
		case os.ModeDir:
			if err = copyDir(srcPath, destPath); err != nil {
				break IterateEntry
			}
		case os.ModeSymlink:
			if err = copySymlink(srcPath, destPath); err != nil {
				break IterateEntry
			}
		default:
			if err = bufferCopyFile(srcPath, destPath, defaultBufferSize); err != nil {
				break IterateEntry
			}
		}
	}

	return
}
