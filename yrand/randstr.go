package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

var (
	errStringAlphabet = errors.New("length of alphabet should be greater than one")
	errStringLength   = errors.New("length of string should be positive")
	alphabetBase36    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetBase62    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// String returns a random string of given length and composed of given chars only.
func String(alphabet string, length int) (s string, err error) {
	base := uint64(len(alphabet))
	if base <= 1 {
		err = errStringAlphabet
		return
	}
	if length <= 0 {
		err = errStringLength
		return
	}

	numBig := new(big.Int)
	bytes := make([]byte, 8)

	sb := strings.Builder{}
	sb.Grow(length)
	left := length
	for left > 0 {
		_, err = rand.Read(bytes)
		if err != nil {
			return
		}

		numBig.SetBytes(bytes)
		num := numBig.Uint64()
		for num > 0 && left > 0 {
			rm := int(num % base)
			num = num / base
			sb.WriteByte(alphabet[rm])
			left--
		}
	}

	s = sb.String()
	if len(s) > length {
		s = s[:length]
	}
	return
}

// StringBase36 returns a random string of given length and composed of A-Z0-9 chars only.
func StringBase36(length int) (s string, err error) {
	return String(alphabetBase36, length)
}

// StringBase62 returns a random string of given length and composed of a-zA-Z0-9 chars only.
func StringBase62(length int) (s string, err error) {
	return String(alphabetBase62, length)
}
