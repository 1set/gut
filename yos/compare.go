package yos

import (
	"bytes"
	"io"
	"os"
)

// SameSymlinkContent checks if the two symbolic links have the same destination.
func SameSymlinkContent(path1, path2 string) (same bool, err error) {
	var link1, link2 string
	if link1, err = os.Readlink(path1); err != nil {
		return
	}
	if link2, err = os.Readlink(path2); err != nil {
		return
	}
	same = link1 == link2
	return
}

// SameFileContent checks if the two given files have the same content or are the same file. Symbolic links are followed.
// Errors are returned if any files doesn't exist or is broken.
func SameFileContent(path1, path2 string) (same bool, err error) {
	var fi1, fi2 os.FileInfo
	if fi1, err = os.Stat(path1); err != nil {
		return
	}
	if fi2, err = os.Stat(path2); err != nil {
		return
	}

	if !(fi1.Mode().IsRegular() && fi2.Mode().IsRegular()) {
		err = ErrNotRegular
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
			if nr1 == 0 && nr2 == 0 {
				same = true
				break
			}
			err = io.ErrUnexpectedEOF
		} else if err1 != nil {
			err = err1
		} else if err2 != nil {
			err = err2
		} else if nr1 != nr2 {
			err = ErrShortRead
		}

		if err != nil {
			break
		}

		if same = bytes.Equal(buf1[:nr1], buf2[:nr2]); !same {
			break
		}
	}

	return
}
