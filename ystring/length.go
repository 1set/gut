package ystring

import (
	"unicode/utf8"
)

// Length returns the number of runes in a given string.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Truncate returns first n runes of s.
func Truncate(s string, n int) string {
	switch {
	case n < 0:
		panic("ystring: negative Truncate length n")
	case n == 0:
		return s[0:0]
	case n >= len(s):
		return s
	}

	cnt := 0
	for idx := range s {
		if n <= cnt {
			return s[0:idx]
		}
		cnt++
	}
	return s
}
