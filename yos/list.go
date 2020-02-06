package yos

import (
	"os"
	"path/filepath"
	"strings"
)

// A FilePathInfo describes path and stat of a file or directory.
type FilePathInfo struct {
	Path string
	Info os.FileInfo
}

// ListAll returns a list of all entries in the given directory in lexical order. The given directory is not included in the list.
//
// It searches recursively, but symbolic links other than the given path will be not be followed.
func ListAll(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return true, nil })
}

// ListFile returns a list of file entries in the given directory in lexical order. The given directory is not included in the list.
//
// It searches recursively, but symbolic links other than the given path will be not be followed.
func ListFile(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isFileFi(&info), nil })
}

// ListSymlink returns a list of symbolic link entries in the given directory in lexical order. The given directory is not included in the list.
//
// It searches recursively, but symbolic links other than the given path will be not be followed.
func ListSymlink(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isSymlinkFi(&info), nil })
}

// ListDir returns a list of nested directory entries in the given directory in lexical order. The given directory is not included in the list.
//
// It searches recursively, but symbolic links other than the given path will be not be followed.
func ListDir(root string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (bool, error) { return isDirFi(&info), nil })
}

// The flags are used by the ListMatch method.
const (
	// ListRecursive indicates ListMatch to recursively list directory entries encountered.
	ListRecursive int = 1 << iota
	// ListRecursive indicates ListMatch to convert file name to lower case before the pattern matching.
	ListToLower
	// ListRecursive indicates ListMatch to include matched files in the returned list.
	ListIncludeFile
	// ListRecursive indicates ListMatch to include matched directories in the returned list.
	ListIncludeDir
)

// ListMatch returns a list of directory entries that matches any given pattern in the directory in lexical order.
// ListMatch requires the pattern to match all of the filename, not just a substring.
// Symbolic links other than the given path will be not be followed. The given directory is not included in the list.
// filepath.ErrBadPattern is returned if any pattern is malformed.
func ListMatch(root string, flag int, patterns ...string) (entries []*FilePathInfo, err error) {
	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		fileName := info.Name()
		if flag&ListToLower != 0 {
			fileName = strings.ToLower(fileName)
		}
		isDir := info.IsDir()
		if (isDir && (flag&ListIncludeDir != 0)) || (!isDir && (flag&ListIncludeFile != 0)) {
			for _, pattern := range patterns {
				ok, err = filepath.Match(pattern, fileName)
				if ok || err != nil {
					break
				}
			}
		}
		if err == nil && isDir && (flag&ListRecursive == 0) {
			err = filepath.SkipDir
		}
		return
	})
}

// listCondEntries returns a list of conditional directory entries.
func listCondEntries(root string, cond func(os.FileInfo) (bool, error)) (entries []*FilePathInfo, err error) {
	var (
		rootFi   os.FileInfo
		rootPath string
	)
	if rootPath, rootFi, err = resolveDirInfo(root); err != nil {
		err = opError(opnList, root, err)
		return
	}

	err = filepath.Walk(rootPath, func(itemPath string, itemFi os.FileInfo, errIn error) (errOut error) {
		errOut = errIn
		if os.SameFile(rootFi, itemFi) || errOut != nil {
			return
		}
		var ok bool
		if ok, errOut = cond(itemFi); ok {
			entries = append(entries, &FilePathInfo{
				Path: itemPath,
				Info: itemFi,
			})
		}
		return
	})
	return
}
