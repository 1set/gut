package yrand

import (
	"crypto/rand"
	"math/big"
)

// IntRange returns a random int64 number within a specified range from min (inclusive) to max (exclusive).
func IntRange(min, max int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return result.Int64() + min
}
