package ystring

import (
	"strings"
)

func TrimAfterFirst(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	}

	if idx := strings.Index(s, substr); idx >= 0 {
		return s[0:idx]
	}
	return s
}
