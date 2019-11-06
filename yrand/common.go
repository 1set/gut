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
	errChoiceEmpty     = errors.New("slice should not be empty")
)

// Check if two float numbers are equal within a given tolerance.
func isEqualFloat(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

// Iterate over an newly generated list of `count` random uint64 numbers in [0, `max`).
func iterateRandomNumbers(count int, max uint64, callback func(uint64) error) (err error) {
	if count <= 0 {
		err = errIterateCount
	} else if max <= 1 {
		err = errIterateMax
	} else if callback == nil {
		err = errIterateCallback
	}
	if err != nil {
		return
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
			if err = callback(rm); err != nil {
				return
			}
		}
	}
	return
}

// Returns a random int number in [0, `max`) for index purpose.
func randomIndex(max int) (idx int, err error) {
	if max <= 0 {
		err = errChoiceEmpty
	} else {
		idx, err = IntRange(0, max)
	}
	return
}
