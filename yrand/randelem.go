package yrand

import (
	"crypto/rand"
	"math/big"
)

func Shuffle(n int, swap func(i, j int)) (err error) {
	if n < 1 {
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
