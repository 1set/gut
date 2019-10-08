package yrand_test

import (
	"fmt"

	"github.com/an63/gut/yrand"
)

// This example generates a random 6-digit PIN.
func ExampleInt64Range() {
	min, max := int64(0), int64(1000000)
	num, err := yrand.Int64Range(min, max)
	if err != nil {
		fmt.Println(err)
		return
	}
	pin := fmt.Sprintf("%06d", num)

	fmt.Println(min <= num && num < max, len(pin))
	// Output: true 6
}

// This example generates a random string of 20 A-Z0-9 chars.
func ExampleStringBase36() {
	s, err := yrand.StringBase36(20)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(s))
	// Output: 20
}
