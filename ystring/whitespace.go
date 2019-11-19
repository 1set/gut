package ystring

import (
	"strings"
	"unicode"
)

// IsEmpty checks if the string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsBlank checks if the string contains only whitespaces.
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// Shrink returns a string that replaces consecutive whitespace characters in s with the sep string.
func Shrink(s, sep string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))

	wFlag := false
	for _, c := range strings.TrimSpace(s) {
		if unicode.IsSpace(c) {
			wFlag = true
		} else {
			if wFlag {
				sb.WriteString(sep)
				wFlag = false
			}
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
