package ystring

import (
	"strings"
)

// TrimAfterFirst returns s without the part after the first instance of substr.
// If substr is empty or not present in s, s is returned unchanged.
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

// TrimAfterLast returns s without the part after the last instance of substr.
// If substr is empty or not present in s, s is returned unchanged.
func TrimAfterLast(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	}

	if idx := strings.LastIndex(s, substr); idx >= 0 {
		return s[0:idx]
	}
	return s
}
