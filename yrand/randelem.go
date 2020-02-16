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
func Shuffle(n int, swap ShuffleSwapFunc) (err error) {
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

// ChoiceInt returns a random element from the non-empty slice of int.
func ChoiceInt(list []int) (n int, err error) {
	var idx int
	if idx, err = randomIndex(len(list)); err == nil {
		n = list[idx]
	}
	return
}

// ChoiceString returns a random element from the non-empty slice of string.
func ChoiceString(list []string) (s string, err error) {
	var idx int
	if idx, err = randomIndex(len(list)); err == nil {
		s = list[idx]
	}
	return
}
