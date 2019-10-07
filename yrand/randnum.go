// Package yrand is yet another wrapper of cryptographically secure random number generator.
package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	errMinMaxRange = errors.New("Min should be less than max.")
)

// Int64Range returns a random int64 number in [min, max).
func Int64Range(min, max int64) (n int64, err error) {
	n = 0
	if min >= max {
		err = errMinMaxRange
		return
	}

	randMax := new(big.Int).SetUint64(uint64(max - min))
	randNum, err := rand.Int(rand.Reader, randMax)
	if err == nil {
		n = randNum.Int64() + min
	}
	return
}

// Float64 returns a random float64 number in [0.0, 1.0).
func Float64() (n float64, err error) {
	n = 0
	const max = 1 << 53
	randNum, err := rand.Int(rand.Reader, big.NewInt(max))
	if err == nil {
		n = float64(randNum.Int64()) / max
	}
	return
}
