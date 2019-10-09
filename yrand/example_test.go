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
		fmt.Println("got error:", err)
		return
	}
	pin := fmt.Sprintf("%06d", num)

	fmt.Println(min <= num && num < max, len(pin))
	// Output: true 6
}

// This example simulates coin toss experiments
func ExampleFloat64() {
	head, tail := 0, 0
	count := 50000
	for i := 0; i < count; i++ {
		n, err := yrand.Float64()
		if err != nil {
			fmt.Println("got error:", err)
			return
		}

		if n < 0.5 {
			head++
		} else {
			tail++
		}
	}

	fmt.Printf("%.1f", float64(head)/float64(tail))
	// Output: 1.0
}

// This example generates a random string of 20 A-Z0-9 chars.
func ExampleStringBase36() {
	s, err := yrand.StringBase36(20)
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println(len(s))
	// Output: 20
}
