package yrand

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	errShuffleNegative = errors.New("n should be non-negative")
)

// Shuffle randomizes the order of elements. n is the number of elements. swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int)) (err error) {
	if n < 0 {
		return errShuffleNegative
	} else if n <= 1 {
		return
	}

	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	for i := uint64(n - 1); i > 0; {
		if _, err = rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > i && i > 0; i-- {
			max := i + 1
			j := int(num % max)
			num /= max
			swap(int(i), j)
		}
	}

	return
}