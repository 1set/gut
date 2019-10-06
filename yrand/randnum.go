package yrand

import (
	"crypto/rand"
	"math/big"
)

// Int64Range returns a random int64 number [min, max).
func Int64Range(min, max int64) (n int64, err error) {
	n = 0
	if min >= max {
		err = MinMaxRangeError
		return
	}

	randMax := new(big.Int).SetUint64(uint64(max - min))
	randNum, err := rand.Int(rand.Reader, randMax)
	if err == nil {
		n = randNum.Int64() + min
	}
	return
}
