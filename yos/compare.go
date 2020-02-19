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
	if path1, path2, err = refineComparePaths(path1, path2); err != nil {
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
	if path1, path2, err = refineComparePaths(path1, path2); err != nil {
		return
	}

	var (
		fi1, fi2     os.FileInfo
		file1, file2 *os.File
	)

	// check file mode of path1, and then path2
	if file1, fi1, err = openFileInfo(path1); err == nil {
		defer file1.Close()
	} else {
		err = opError(opnCompare, path1, err)
		return
	}
	if file2, fi2, err = openFileInfo(path2); err == nil {
		defer file2.Close()
	} else {
		err = opError(opnCompare, path2, err)
		return
	}

	// quick check if it's the identical file and file size
	if same = os.SameFile(fi1, fi2); same {
		return
	} else if fi1.Size() != fi2.Size() {
		return
	}

	same, err = compareReaderContent(file1, file2, path1, path2)
	return
}

// SameDirEntries checks if the two directories have the same entries. Symbolic links other than the given paths will be not be followed, and only compares content of links.
func SameDirEntries(path1, path2 string) (same bool, err error) {
	var (
		fi1, fi2       os.FileInfo
		raw1, raw2     = path1, path2
		items1, items2 []*FilePathInfo
	)
	// resolve paths if they're symbolic links
	if path1, fi1, err = resolveDirInfo(path1); err != nil {
		err = opError(opnCompare, raw1, err)
		return
	}
	if path2, fi2, err = resolveDirInfo(path2); err != nil {
		err = opError(opnCompare, raw2, err)
		return
	}

	// quick check if it's the identical directory
	if same = os.SameFile(fi1, fi2); same {
		return
	}

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

CompareEntries:
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
				break CompareEntries
			}
		case os.ModeDir:
			// ignore the directory structure here, since it's already compared by the relative path logic before
		case 0:
			if same, err = SameFileContent(entry1.Path, entry2.Path); err != nil || !same {
				break CompareEntries
			}
		}
	}
	return
}

func compareReaderContent(rd1, rd2 io.Reader, path1, path2 string) (same bool, err error) {
	buf1, buf2 := make([]byte, fileCompareChunkSize), make([]byte, fileCompareChunkSize)
CompareContent:
	for {
		nr1, err1 := rd1.Read(buf1)
		nr2, err2 := rd2.Read(buf2)

		switch {
		case err1 == io.EOF && err2 == io.EOF:
			switch {
			case nr1 == 0 && nr2 == 0:
				same = true
				break CompareContent
			case nr1 > 0:
				err = opError(opnCompare, path1, io.ErrUnexpectedEOF)
			case nr2 > 0:
				err = opError(opnCompare, path2, io.ErrUnexpectedEOF)
			}
		case err1 != nil:
			err = opError(opnCompare, path1, err1)
		case err2 != nil:
			err = opError(opnCompare, path2, err2)
		case nr1 < nr2:
			err = opError(opnCompare, path1, errShortRead)
		case nr2 < nr1:
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
