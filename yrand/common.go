package yrand

import (
	"crypto/rand"
	"math"
	"math/big"
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
