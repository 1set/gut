package yos

import (
	"path/filepath"
)

// JoinPath joins any number of path elements into a single path, adding a separator if necessary.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
