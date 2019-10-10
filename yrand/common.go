package yrand

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
)

var (
	errIterateMax      = errors.New("max value should be greater than one")
	errIterateCount    = errors.New("count should be positive")
	errIterateCallback = errors.New("callback should not be nil")
)

func isEqualFloat(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func iterateRandomNumbers(count int, max uint64, callback func(num uint64)) (err error) {
	if count <= 0 {
		return errIterateCount
	}
	if max <= 1 {
		return errIterateMax
	}
	if callback == nil {
		return errIterateCallback
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