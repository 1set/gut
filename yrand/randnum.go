// Package yrand is yet another wrapper package of cryptographically secure random number generator.
package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	minMaxRangeError = errors.New("Min should be less than max.")
)

// Int64Range returns a random int64 number [min, max).
func Int64Range(min, max int64) (n int64, err error) {
	n = 0
	if min >= max {
		err = minMaxRangeError
		return
	}

	randMax := new(big.Int).SetUint64(uint64(max - min))
	randNum, err := rand.Int(rand.Reader, randMax)
	if err == nil {
		n = randNum.Int64() + min
	}
	return
}
