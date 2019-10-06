package yrand

import (
	"crypto/rand"
	"math/big"
)

// IntRange returns a random int64 number [min, max).
func IntRange(min, max int64) (n int64, err error) {
	if min >= max {
		return 0, MinMaxRangeError
	}
	result, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return 0, err
	}
	return result.Int64() + min, nil
}
