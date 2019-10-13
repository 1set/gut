// Package yrand is yet another wrapper of cryptographically secure random number generator.
package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	errMinMaxRange = errors.New("min should be less than max")
)

// Int64Range returns a random int64 number in [min, max).
func Int64Range(min, max int64) (n int64, err error) {
	n = 0
	if min >= max {
		err = errMinMaxRange
		return
	}

	randMax := new(big.Int).SetUint64(uint64(max - min))
	var randNum *big.Int
	if randNum, err = rand.Int(rand.Reader, randMax); err == nil {
		n = randNum.Int64() + min
	}
	return
}

// Int32Range returns a random int32 number in [min, max).
func Int32Range(min, max int32) (n int32, err error) {
	num, err := Int64Range(int64(min), int64(max))
	return int32(num), err
}

// IntRange returns a random int number in [min, max).
func IntRange(min, max int) (n int, err error) {
	num, err := Int64Range(int64(min), int64(max))
	return int(num), err
}

// Uint64Range returns a random uint64 number in [min, max).
func Uint64Range(min, max uint64) (n uint64, err error) {
	n = 0
	if min >= max {
		err = errMinMaxRange
		return
	}

	randMax := new(big.Int).SetUint64(max - min)
	var randNum *big.Int
	if randNum, err = rand.Int(rand.Reader, randMax); err == nil {
		n = randNum.Uint64() + min
	}
	return
}

// Float64 returns a random float64 number in [0.0, 1.0).
func Float64() (n float64, err error) {
	return getRandomFloat(1 << 53)
}

// Float32 returns a random float32 number in [0.0, 1.0).
func Float32() (n float32, err error) {
	num, err := getRandomFloat(1 << 24)
	return float32(num), err
}

func getRandomFloat(prec int64) (n float64, err error) {
	n = 0
	var randNum *big.Int
	if randNum, err = rand.Int(rand.Reader, big.NewInt(prec)); err == nil {
		n = float64(randNum.Int64()) / float64(prec)
	}
	return
}
