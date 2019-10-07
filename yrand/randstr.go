package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

var (
	errStringAlphabet = errors.New("length of alphabet should be greater than one")
	errStringLimit = errors.New("limit of string should be positive")
	alphabetBase62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// String returns a random string of given length and composed of given chars only.
func String(alphabet string, limit int) (s string, err error) {
	base := uint64(len(alphabet))
	if base <= 1 {
		err = errStringAlphabet
		return
	}
	if limit <= 0 {
		err = errStringLimit
		return
	}

	numBig := new(big.Int)
	bytes := make([]byte, 8)
	sb := strings.Builder{}
	for sb.Len() < limit {
		_, err = rand.Read(bytes)
		if err != nil {
			return
		}

		numBig.SetBytes(bytes)
		num := numBig.Uint64()

		for num > 0 {
			rm := int(num % base)
			num = num / base
			sb.WriteByte(alphabet[rm])
		}
	}

	s = sb.String()
	if len(s) > limit {
		s = s[:limit]
	}
	return
}

// StringBase62 returns a random string of given length and composed of a-zA-Z0-9 chars only.
func StringBase62(limit int) (s string, err error) {
	return String(alphabetBase62, limit)
}
