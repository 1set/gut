package ystring

import (
	"strings"
)

var (
	emptyString = ""
)

func TrimAfterFirst(s, substr string) string {
	ls, lsub := len(s), len(substr)
	switch {
	case ls == 0 || lsub == 0:
		return s
	case ls == lsub && s == substr:
		return emptyString
	case ls == lsub && s != substr:
		return s
	case ls < lsub:
		return s
	}

	idx := strings.Index(s, substr)
	switch {
	case idx == 0:
		return emptyString
	case idx > 0:
		return s[0:idx]
	}
	return s
}
