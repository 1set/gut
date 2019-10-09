package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
	"unicode/utf8"
)

var (
	errStringAlphabet = errors.New("length of alphabet should be greater than one")
	errStringLength   = errors.New("length of string should be positive")
	errIterateMax     = errors.New("max value should be greater than one")
	errIterateCount   = errors.New("count should be positive")
	alphabetLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetBase36    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetBase62    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// String returns a random string of given length with given ASCII chars only.
func String(alphabet string, length int) (s string, err error) {
	if length <= 0 {
		err = errStringLength
		return
	}

	base := uint64(len(alphabet))
	if base <= 1 {
		err = errStringAlphabet
		return
	}

	sb := strings.Builder{}
	sb.Grow(length)
	writeBack := func(num uint64) {
		sb.WriteByte(alphabet[int(num)])
	}

	if err = iterateRandomNumbers(length, base, writeBack); err == nil {
		s = sb.String()
	}
	return
}

// Runes returns a random string of given length with given Unicode chars only.
func Runes(alphabet string, length int) (s string, err error) {
	if length <= 0 {
		err = errStringLength
		return
	}

	base := uint64(utf8.RuneCountInString(alphabet))
	if base <= 1 {
		err = errStringAlphabet
		return
	}

	sb := strings.Builder{}
	sb.Grow(length * 4)
	abRunes := []rune(alphabet)
	writeBack := func(num uint64) {
		sb.WriteRune(abRunes[int(num)])
	}

	if err = iterateRandomNumbers(length, base, writeBack); err == nil {
		s = sb.String()
	}
	return
}

func iterateRandomNumbers(count int, max uint64, callback func(num uint64)) (err error) {
	if count <= 0 {
		return errIterateCount
	}
	if max <= 1 {
		return errIterateMax
	}

	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	for left := count; left > 0; {
		if _, err = rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > 0 && left > 0; left-- {
			rm := num % max
			num /= max
			callback(rm)
		}
	}
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
