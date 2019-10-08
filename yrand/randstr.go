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
	alphabetLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetBase36    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetBase62    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// String returns a random string of given length with given chars only.
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

	for left := length; left > 0; {
		_, err = rand.Read(bytes)
		if err != nil {
			return
		}

		numBig.SetBytes(bytes)
		for num := numBig.Uint64(); num > 0 && left > 0; left-- {
			rm := int(num % base)
			num /= base
			sb.WriteByte(alphabet[rm])
		}
	}

	s = sb.String()
	return
}

// StringLetters returns a random string of given length with A-Z chars only.
func StringLetters(length int) (s string, err error) {
	return String(alphabetLetters, length)
}

// StringBase36 returns a random string of given length with A-Z0-9 chars only.
func StringBase36(length int) (s string, err error) {
	return String(alphabetBase36, length)
}

// StringBase62 returns a random string of given length with a-zA-Z0-9 chars only.
func StringBase62(length int) (s string, err error) {
	return String(alphabetBase62, length)
}
