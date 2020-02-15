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

// Checks if two floats are equal within a given tolerance.
func isFloatEqual(a, b, tolerance float64) (equal bool) {
	const pos, neg = 1, -1
	switch {
	case math.IsNaN(a), math.IsInf(a, pos), math.IsInf(a, neg), math.IsNaN(b), math.IsInf(b, pos), math.IsInf(b, neg):
		equal = false
	case a != 0 && b != 0:
		equal = math.Abs((a-b)/a) <= tolerance
	default:
		equal = math.Abs(a-b) <= tolerance
	}
	return
}

// Iterates over an newly generated list of `count` random uint64 numbers in [0, `max`).
func iterateRandomNumbers(count int, max uint64, callback func(uint64) error) (err error) {
	switch {
	case count <= 0:
		err = errIterateCount
	case max <= 1:
		err = errIterateMax
	case callback == nil:
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
