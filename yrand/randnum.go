package yrand

import (
	"crypto/rand"
	"math/big"
)

// Int64Range returns a random int64 number [min, max).
func Int64Range(min, max int64) (n int64, err error) {
	if min >= max {
		return 0, MinMaxRangeError
	}

	num := new(big.Int).SetUint64(uint64(max - min))
	result, err := rand.Int(rand.Reader, num)
	if err != nil {
		return 0, err
	}
	return result.Int64() + min, nil
}
