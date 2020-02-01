package yos

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/1set/gut/ystring"
)

var (
	// CompareFileModeMask is a mask for file mode bits to compare in SameDirEntries.
	CompareFileModeMask = os.ModeDir | os.ModeSymlink
)

// SameFileContent checks if the two given files have the same content or are the same file. Symbolic links are followed.
// Errors are returned if any files doesn't exist or is broken.
func SameFileContent(path1, path2 string) (same bool, err error) {
	if path1, path2, err = refineComparePaths(path1, path2); err != nil {
		return
	}

	var fi1, fi2 os.FileInfo
	if fi1, err = os.Stat(path1); err != nil {
		err = &os.PathError{opnCompare, path1, err}
		return
	}
	if fi2, err = os.Stat(path2); err != nil {
		err = &os.PathError{opnCompare, path2, err}
		return
	}

	if !fi1.Mode().IsRegular() {
		err = &os.PathError{opnCompare, path1, errNotRegularFile}
		return
	}
	if !fi2.Mode().IsRegular() {
		err = &os.PathError{opnCompare, path2, errNotRegularFile}
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
		err = &os.PathError{opnCompare, path1, err}
		return
	}
	defer file1.Close()

	if file2, err = os.Open(path2); err != nil {
		err = &os.PathError{opnCompare, path2, err}
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
			if nr1 > 0 {
				err = &os.PathError{opnCompare, path1, io.ErrUnexpectedEOF}
			} else {
				err = &os.PathError{opnCompare, path2, io.ErrUnexpectedEOF}
			}
		} else if err1 != nil {
			err = &os.PathError{opnCompare, path1, err1}
		} else if err2 != nil {
			err = &os.PathError{opnCompare, path2, err2}
		} else if nr1 != nr2 {
			if nr1 < nr2 {
				err = &os.PathError{opnCompare, path1, errShortRead}
			} else {
				err = &os.PathError{opnCompare, path2, errShortRead}
			}
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

// SameSymlinkContent checks if the two symbolic links have the same destination.
func SameSymlinkContent(path1, path2 string) (same bool, err error) {
	if path1, path2, err = refineComparePaths(path1, path2); err != nil {
		return
	}

	var link1, link2 string
	if link1, err = os.Readlink(path1); err != nil {
		err = &os.PathError{opnCompare, path1, err}
		return
	}
	if link2, err = os.Readlink(path2); err != nil {
		err = &os.PathError{opnCompare, path2, err}
		return
	}

	same = link1 == link2
	return
}

// SameDirEntries checks if the two directories have the same entries. Symbolic links will be not be followed, and only compares the contents.
func SameDirEntries(path1, path2 string) (same bool, err error) {
	if path1, path2, err = refineComparePaths(path1, path2); err != nil {
		return
	}

	var fi1, fi2 os.FileInfo
	if fi1, err = os.Stat(path1); err != nil {
		err = &os.PathError{opnCompare, path1, err}
		return
	}
	if fi2, err = os.Stat(path2); err != nil {
		err = &os.PathError{opnCompare, path2, err}
		return
	}

	if !fi1.IsDir() {
		err = &os.PathError{opnCompare, path1, errNotDirectory}
		return
	}
	if !fi2.IsDir() {
		err = &os.PathError{opnCompare, path2, errNotDirectory}
		return
	}

	if os.SameFile(fi1, fi2) {
		same = true
		return
	}

	var items1, items2 []*FilePathInfo
	if items1, err = ListAll(path1); err != nil {
		err = &os.PathError{opnCompare, path1, err}
		return
	}
	if items2, err = ListAll(path2); err != nil {
		err = &os.PathError{opnCompare, path2, err}
		return
	}

	num1, num2 := len(items1), len(items2)
	if same = num1 == num2; !same {
		return
	}

IterateItems:
	for idx := 0; idx < num1; idx++ {
		entry1, entry2 := items1[idx], items2[idx]

		relativePath1, relativePath2 := strings.Replace(entry1.Path, path1, "", 1), strings.Replace(entry2.Path, path2, "", 1)
		if same = relativePath1 == relativePath2; !same {
			break
		}

		entryMode1, entryMode2 := entry1.Info.Mode(), entry2.Info.Mode()
		if same = entryMode1&CompareFileModeMask == entryMode2&CompareFileModeMask; !same {
			break
		}

		switch entryMode1 & os.ModeType {
		case os.ModeSymlink:
			if same, err = SameSymlinkContent(entry1.Path, entry2.Path); err != nil || !same {
				break IterateItems
			}
		case os.ModeDir:
		default:
			if same, err = SameFileContent(entry1.Path, entry2.Path); err != nil || !same {
				break IterateItems
			}
		}
	}

	return
}

// refineComparePaths validates, cleans up for file comparison.
func refineComparePaths(pathRaw1, pathRaw2 string) (path1, path2 string, err error) {
	if ystring.IsBlank(pathRaw1) {
		err = &os.PathError{opnCompare, pathRaw1, errInvalidPath}
		return
	}
	if ystring.IsBlank(pathRaw2) {
		err = &os.PathError{opnCompare, pathRaw2, errInvalidPath}
		return
	}

	// clean up paths
	path1, path2 = filepath.Clean(pathRaw1), filepath.Clean(pathRaw2)
	return
}
