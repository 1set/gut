package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

var alphabetDigit = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	errStringLimit = errors.New("limit of string should be positive")
)

func StringBase62(limit int) (s string, err error) {
	if limit <= 0 {
		err = errStringLimit
		return
	}

	base := uint64(len(alphabetDigit))
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
			sb.WriteByte(alphabetDigit[rm])
		}
	}

	s = sb.String()
	if len(s) > limit {
		s = s[:limit]
	}

	return
}
