package yos

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var (
	// compareFileModeMask is a mask for file mode bits to compare in SameDirEntries.
	compareFileModeMask = os.ModeDir | os.ModeSymlink
	// fileCompareChunkSize represents the buffer size for readers of SameFileContent.
	fileCompareChunkSize = 64 * 1024
)

// SameSymlinkContent checks if the two symbolic links have the same destination.
func SameSymlinkContent(path1, path2 string) (same bool, err error) {
	if path1, path2, err = refineComparePaths(path1, path2, nil, nil); err != nil {
		return
	}

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
	switch path1, path2, err = refineComparePaths(path1, path2, isFileFi, errNotRegularFile); {
	case err == errSameFile:
		same, err = true, nil
		return
	case err == errDiffFileSize:
		same, err = false, nil
		return
	case err != nil:
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

	buf1, buf2 := make([]byte, fileCompareChunkSize), make([]byte, fileCompareChunkSize)
	for {
		nr1, err1 := file1.Read(buf1)
		nr2, err2 := file2.Read(buf2)

		if err1 == io.EOF && err2 == io.EOF {
			if nr1 == 0 && nr2 == 0 {
				same = true
				break
			} else if nr1 > 0 {
				err = opError(opnCompare, path1, io.ErrUnexpectedEOF)
			} else if nr2 > 0 {
				err = opError(opnCompare, path2, io.ErrUnexpectedEOF)
			}
		} else if err1 != nil {
			err = opError(opnCompare, path1, err1)
		} else if err2 != nil {
			err = opError(opnCompare, path2, err2)
		} else if nr1 < nr2 {
			err = opError(opnCompare, path1, errShortRead)
		} else if nr2 < nr1 {
			err = opError(opnCompare, path2, errShortRead)
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

// SameDirEntries checks if the two directories have the same entries. Symbolic links will be not be followed, and only compares the contents.
func SameDirEntries(path1, path2 string) (same bool, err error) {
	switch path1, path2, err = refineComparePaths(path1, path2, isDirFi, errNotDirectory); {
	case err == errSameFile:
		same, err = true, nil
		return
	case err != nil:
		return
	}

	var items1, items2 []*FilePathInfo
	if items1, err = ListAll(path1); err != nil {
		return
	}
	if items2, err = ListAll(path2); err != nil {
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
		if same = entryMode1&compareFileModeMask == entryMode2&compareFileModeMask; !same {
			break
		}

		switch entryMode1 & os.ModeType {
		case os.ModeSymlink:
			if same, err = SameSymlinkContent(entry1.Path, entry2.Path); err != nil || !same {
				break IterateItems
			}
		case os.ModeDir:
		case 0:
			if same, err = SameFileContent(entry1.Path, entry2.Path); err != nil || !same {
				break IterateItems
			}
		}
	}

	return
}
