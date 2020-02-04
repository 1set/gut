package yos

import (
	"io"
	"io/ioutil"
	"math/bits"
	"os"
)

const (
	defaultDirectoryFileMode = os.FileMode(0755)
	defaultBufferSize        = 256 * 1024
)

// CopyFile copies a file to a target file or directory. Symbolic links are followed.
//
// If the target is an existing file, the target will be overwritten with the source file.
//
// If the target is an existing directory, the source file will be copied to the directory with the same file name.
//
// If the target doesn't exist but its parent directory does, the source file will be copied to the parent directory with the target name.
//
// If there is an error, it'll be of type *os.PathError.
func CopyFile(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, true); err == nil {
		err = bufferCopyFile(src, dest, defaultBufferSize)
	}
	return
}

// CopyDir copies a directory to a target directory recursively. Symbolic links inside the directories will be copied instead of being followed.
//
// If the target is an existing file, an error will be returned.
//
// If the target is an existing directory, the source directory will be copied to the directory with the same name.
//
// If the target doesn't exist but its parent directory does, the source directory will be copied to the parent directory with the target name.
//
// It stops and returns immediately if any error occurs, and the error will be of type *os.PathError.
func CopyDir(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, true); err == nil {
		err = copyDir(src, dest)
	}
	return
}

// CopySymlink copies a symbolic link to a target file.
// It only copies the contents and makes no attempt to read the referenced file.
// If there is an error, it'll be of type *os.PathError.
func CopySymlink(src, dest string) (err error) {
	if src, dest, err = refineOpPaths(opnCopy, src, dest, false); err == nil {
		err = copySymlink(src, dest)
	}
	return
}

// bufferCopyFile reads content from the source file and write to the destination file with a buffer.
func bufferCopyFile(src, dest string, bufferSize int64) (err error) {
	var (
		srcFile, destFile *os.File
		srcInfo, destInfo os.FileInfo
	)

	// check if source file exists and open for read
	if srcFile, srcInfo, err = openFileInfo(src); err == nil {
		defer srcFile.Close()
	} else {
		err = opError(opnCopy, src, err)
		return
	}

	// check if source and destination files are identical
	if destInfo, err = os.Stat(dest); err == nil {
		if !isFileFi(&destInfo) {
			err = opError(opnCopy, dest, errNotRegularFile)
		} else if os.SameFile(srcInfo, destInfo) {
			err = opError(opnCopy, dest, errSameFile)
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
		if nr, err = srcFile.Read(buf); err != nil || nr == 0 {
			if err == io.EOF && nr > 0 {
				err = opError(opnCopy, src, io.ErrUnexpectedEOF)
			}
			break
		}

		if nw, err = destFile.Write(buf[:nr]); err != nil {
			break
		} else if nw != nr {
			err = opError(opnCopy, dest, io.ErrShortWrite)
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	// err = destFile.Sync()
	return
}

// copySymlink reads content from the source symbolic link and write to the destination symbolic link.
func copySymlink(src, dest string) (err error) {
	var destInfo os.FileInfo
	if destInfo, err = os.Lstat(dest); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
	} else {
		if isDirFi(&destInfo) {
			// avoid overwriting directory
			err = opError(opnCopy, dest, errIsDirectory)
		} else {
			err = os.Remove(dest)
		}
	}
	if err != nil {
		return
	}

	var link string
	if link, err = os.Readlink(src); err != nil {
		err = opError(opnCopy, src, err)
	} else if err = os.Symlink(link, dest); err != nil {
		err = opError(opnCopy, dest, err)
	}
	return
}

// copyDir copies all entries of source directory to destination directory recursively.
func copyDir(src, dest string) (err error) {
	var srcInfo, destInfo os.FileInfo

	// check if source exists and is a directory
	if srcInfo, err = os.Stat(src); err == nil {
		if !isDirFi(&srcInfo) {
			err = opError(opnCopy, src, errNotDirectory)
		}
	}
	if err != nil {
		return
	}

	// check if destination doesn't exist or is not a file or source itself
	if destInfo, err = os.Stat(dest); err == nil {
		if !isDirFi(&destInfo) {
			err = opError(opnCopy, dest, errNotDirectory)
		} else if os.SameFile(srcInfo, destInfo) {
			err = opError(opnCopy, dest, errSameFile)
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
		case 0:
			if err = bufferCopyFile(srcPath, destPath, defaultBufferSize); err != nil {
				break IterateEntry
			}
		}
	}

	return
}
