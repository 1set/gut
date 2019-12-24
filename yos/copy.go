package yos

import (
	"errors"
	"io"
	"math/bits"
	"os"
	"path/filepath"

	"github.com/1set/gut/ystring"
)

var (
	ErrShortRead  = errors.New("short read")
	ErrEmptyPath  = errors.New("path is empty")
	ErrSameFile   = errors.New("files are identical")
	ErrNotRegular = errors.New("file is not regular")
)

// CopyFile copies a file to a target file or directory. Symbolic links are followed.
// If the target is an existing file, the target will be overwritten with the source file.
// If the target is an existing directory, the source file will be copied to the directory with the same file name.
// If the target doesn't exist but its parent directory does, the source file will be copied to the parent directory with the target name.
// ErrSameFile is returned if it detects an attempt to copy a file to itself.
func CopyFile(src, dest string) (err error) {
	if ystring.IsBlank(src) || ystring.IsBlank(dest) {
		err = ErrEmptyPath
		return
	}

	// clean up paths
	src, dest = filepath.Clean(src), filepath.Clean(dest)

	// check if source exists
	var srcInfo, destInfo os.FileInfo
	if srcInfo, err = os.Stat(src); err != nil {
		return
	}

	// check if destination exists
	if destInfo, err = os.Stat(dest); err != nil {
		// check existence of parent of the missing destination
		if os.IsNotExist(err) {
			_, err = os.Stat(filepath.Dir(dest))
		}
	} else {
		// append file name of source to path of the existing destination
		if destInfo.IsDir() {
			dest = JoinPath(dest, srcInfo.Name())
		}
	}

	if err != nil {
		return
	}

	return bufferCopyFile(src, dest, 256*1024)
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
		if os.SameFile(srcInfo, destInfo) {
			err = ErrSameFile
		} else if !destInfo.Mode().IsRegular() {
			err = ErrNotRegular
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
		if nr, err = srcFile.Read(buf); err != nil && err != io.EOF {
			return
		}
		if nr == 0 {
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
