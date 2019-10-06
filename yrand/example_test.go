package yrand_test

import (
	"fmt"

	"github.com/an63/gut/yrand"
)

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
