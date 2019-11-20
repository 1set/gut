package ystring

import (
	"strings"
	"unicode"
	"unicode/utf8"
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

// Length returns the number of runes in a given string.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Truncate returns first n runes of s.
func Truncate(s string, n int) string {
	if n < 0 {
		panic("ystring: negative Truncate length n")
	} else if n == 0 {
		return s[0:0]
	}

	cnt := 0
	for idx := range s {
		if n <= cnt {
			return s[0:idx]
		}
		cnt += 1
	}
	return s
}
