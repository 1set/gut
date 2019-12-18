package yos

import (
	"io"
	"math/bits"
	"os"
)

func CopyFile(src, dest string) (err error) {
	return bufferCopyFile(src, dest, 256*1024)
}

func bufferCopyFile(src, dest string, bufferSize int64) (err error) {
	var srcFile, destFile *os.File
	if srcFile, err = os.Open(src); err != nil {
		return
	}
	defer srcFile.Close()

	var srcInfo os.FileInfo
	if srcInfo, err = os.Stat(src); err != nil {
		return
	}

	fileSize := srcInfo.Size()
	if bufferSize > fileSize {
		bufferSize = 1 << bits.Len64(uint64(fileSize))
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
