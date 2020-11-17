package ystring

import (
	"strings"
	"unicode"
)

// IsEmpty checks if the string is empty.
func IsEmpty(s string) bool {
	return s == ""
}

// IsNotEmpty checks if the string is not empty.
func IsNotEmpty(s string) bool {
	return s != ""
}

// NotEmptyOrDefault returns the string if it is not empty, or returns the fallback.
func NotEmptyOrDefault(s, fallback string) string {
	if s != "" {
		return s
	}
	return fallback
}

// IsBlank checks if the string contains only whitespaces.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank checks if the string contains any non-whitespace characters.
func IsNotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

// NotBlankOrDefault returns the string if it is not blank, or returns the fallback.
func NotBlankOrDefault(s, fallback string) string {
	if strings.TrimSpace(s) != "" {
		return s
	}
	return fallback
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
