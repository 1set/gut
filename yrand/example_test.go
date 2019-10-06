package yrand_test

import (
	"fmt"

	"github.com/an63/gut/yrand"
)

// This example tries to get a random integer between 100 (inclusive) and 200 (exclusive)
// and verify if the number is in the range.
func ExampleInt64Range() {
	min, max := int64(100), int64(200)
	num, err := yrand.Int64Range(min, max)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(min <= num && num < max)

	// Output: true
}
