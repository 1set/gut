package ystring

import (
	"strings"
)

// TrimAfterFirst returns s without the part after the first instance of substr and that instance itself.
// If substr is empty or not present in s, s is returned unchanged.
func TrimAfterFirst(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	default:
		if idx := strings.Index(s, substr); idx >= 0 {
			return s[0:idx]
		}
	}
	return s
}

// TrimAfterLast returns s without the part after the last instance of substr and that instance itself.
// If substr is empty or not present in s, s is returned unchanged.
func TrimAfterLast(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	default:
		if idx := strings.LastIndex(s, substr); idx >= 0 {
			return s[0:idx]
		}
	}
	return s
}

// TrimBeforeFirst returns s without the part before the first instance of substr and that instance itself.
// If substr is empty or not present in s, s is returned unchanged.
func TrimBeforeFirst(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	default:
		if idx := strings.Index(s, substr); idx >= 0 {
			return s[idx+lsub:]
		}
	}
	return s
}

// TrimBeforeLast returns s without the part before the last instance of substr and that instance itself.
// If substr is empty or not present in s, s is returned unchanged.
func TrimBeforeLast(s, substr string) string {
	switch ls, lsub := len(s), len(substr); {
	case ls == 0 || lsub == 0:
		return s
	case ls < lsub:
		return s
	default:
		if idx := strings.LastIndex(s, substr); idx >= 0 {
			return s[idx+lsub:]
		}
	}
	return s
}
