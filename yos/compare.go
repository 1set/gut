package yos

import (
	"bytes"
	"errors"
	"io"
	"os"
)

var (
	ErrShortRead     = errors.New("short read")
	ErrPathDirectory = errors.New("path is directory")
)

func SameContent(path1, path2 string) (same bool, err error) {
	var fi1, fi2 os.FileInfo
	if fi1, err = os.Stat(path1); err != nil {
		return
	}
	if fi2, err = os.Stat(path2); err != nil {
		return
	}

	if fi1.IsDir() || fi2.IsDir() {
		err = ErrPathDirectory
		return
	}

	if os.SameFile(fi1, fi2) {
		same = true
		return
	}

	if fi1.Size() != fi2.Size() {
		return
	}

	var file1, file2 *os.File
	if file1, err = os.Open(path1); err != nil {
		return
	}
	defer file1.Close()

	if file2, err = os.Open(path2); err != nil {
		return
	}
	defer file2.Close()

	const chunkSize = 64 * 1024
	buf1, buf2 := make([]byte, chunkSize), make([]byte, chunkSize)
	for {
		nr1, err1 := file1.Read(buf1)
		nr2, err2 := file2.Read(buf2)

		if err1 == io.EOF && err2 == io.EOF {
			if nr1 > 0 || nr2 > 0 {
				err = io.ErrUnexpectedEOF
			} else {
				same = true
			}
			break
		}

		if err1 != nil {
			err = err1
			break
		}

		if err2 != nil {
			err = err2
			break
		}

		if nr1 != nr2 {
			err = ErrShortRead
			break
		}

		if same = bytes.Equal(buf1[:nr1], buf2[:nr2]); !same {
			break
		}
	}

	return
}

/*

Path1 is a symlink to a directory
Path1 is a symlink to a file and path2 is the file
Path1 is a symlink to a file and path2 is a file with same content
Path1 is a symlink to a symlink and path2 is the symlink to a file
Path1 is a symlink to a symlink and path2 is the symlink to a directory
Path1 is a symlink to a symlink and path2 is the symlink to path1
Path1 is a symlink to a symlink and path2 is the symlink to itself
Path1 is a symlink to a symlink and path2 is the symlink which is broken
Path1 is a symlink to a symlink and path2 is the symlink to another symlink which is broken
Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a file
Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a directory
Path1 is a symlink to a symlink and path2 is the symlink to another symlink to path1
Path1 and path2 are symlinks to the same file
Path1 and path2 are symlinks to files with same content
*/
