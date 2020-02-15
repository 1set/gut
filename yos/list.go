package yos

import (
	"os"
	"path/filepath"
	"regexp"
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
	// ListToLower indicates ListMatch to convert file name to lower case before the pattern matching.
	ListToLower
	// ListUseRegExp indicates ListMatch to use regular expression for the pattern matching.
	ListUseRegExp
	// ListIncludeDir indicates ListMatch to include matched directories in the returned list.
	ListIncludeDir
	// ListIncludeFile indicates ListMatch to include matched files in the returned list.
	ListIncludeFile
	// ListIncludeSymlink indicates ListMatch to include matched symbolic link in the returned list.
	ListIncludeSymlink
)

const (
	// ListIncludeAll indicates ListMatch to include all the matched in the returned list.
	ListIncludeAll = ListIncludeDir | ListIncludeFile | ListIncludeSymlink
)

// ListMatch returns a list of directory entries that matches any given pattern in the directory in lexical order.
//
// Symbolic links other than the given path will be not be followed. The given directory is not included in the list.
//
// ListMatch requires the pattern to match the full file name, not just a substring. Errors are returned if any pattern is malformed.
func ListMatch(root string, flag int, patterns ...string) (entries []*FilePathInfo, err error) {
	var (
		useRegExp  = flag&ListUseRegExp != 0
		rePatterns []*regexp.Regexp
		patRe      *regexp.Regexp
	)
	if useRegExp {
		rePatterns = make([]*regexp.Regexp, 0, len(patterns))
		for _, pat := range patterns {
			if patRe, err = regexp.Compile(pat); err == nil {
				rePatterns = append(rePatterns, patRe)
			} else {
				err = opError(opnList, pat, err)
				return
			}
		}
	}

	return listCondEntries(root, func(info os.FileInfo) (ok bool, err error) {
		fileName := info.Name()
		if flag&ListToLower != 0 {
			fileName = strings.ToLower(fileName)
		}

		if tf := flag & ListIncludeAll; tf != 0 {
			if (tf == ListIncludeAll) ||
				(tf&ListIncludeDir != 0 && isDirFi(&info)) ||
				(tf&ListIncludeFile != 0 && isFileFi(&info)) ||
				(tf&ListIncludeSymlink != 0 && isSymlinkFi(&info)) {
				if useRegExp {
					for _, pat := range rePatterns {
						if ok = pat.MatchString(fileName); ok {
							break
						}
					}
				} else {
					for _, pat := range patterns {
						if ok, err = filepath.Match(pat, fileName); ok || err != nil {
							break
						}
					}
				}
			}
		}

		if err == nil && (flag&ListRecursive == 0) && isDirFi(&info) {
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
