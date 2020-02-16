package yrand_test

import (
	"fmt"

	"github.com/1set/gut/yrand"
)

// This example generates a random 6-digit PIN.
func ExampleInt64Range() {
	min, max := int64(0), int64(1000000)
	pin, err := yrand.Int64Range(min, max)
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Printf("PIN: %06d\n", pin)
}

// This example chooses a random fish from a given list.
func ExampleIntRange() {
	fishes := []string{
		"Ahi",
		"Basa",
		"Kajiki",
		"Mahi Mahi",
		"Monchong",
		"Salmon",
		"Tilapia",
		"Tuna",
	}
	idx, err := yrand.IntRange(0, len(fishes))
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fish := fishes[idx]
	fmt.Println("I 🥰", fish, "🐠")
}

// This example simulates coin toss experiments
func ExampleFloat64() {
	head, tail := 0, 0
	count := 100000
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

	fmt.Printf("Head: %d\nTail: %d\nRate: %.3f", head, tail, float64(head)/float64(tail))
}

// This example generates a random string of 20 A-Z0-9 chars.
func ExampleStringBase36() {
	key, err := yrand.StringBase36(20)
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("Key:", key)
}

// This example randomizes the order of a list of numbers.
func ExampleShuffle() {
	num := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	swapFunc := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}
	fmt.Println("before:", num)

	count := len(num)
	if err := yrand.Shuffle(count, swapFunc); err != nil {
		fmt.Println("got error:", err)
		return
	}
	fmt.Println("after: ", num)
}

// This example chooses a random food from a given list.
func ExampleChoiceString() {
	foods := []string{
		"Ahi Poke",
		"Bouillabaisse",
		"Hukilau Chowder",
		"Kalua Pork",
		"Lau lau",
		"Lobster Casarecce",
		"Loco Moco",
		"Manapua",
	}
	food, err := yrand.ChoiceString(foods)
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("I 🍴", food)
}

// This example chooses a random person according to associated weights.
func ExampleWeightedChoice() {
	var (
		candidates = []string{"Knuth", "Morris", "Pratt"}
		weights    = []float64{0.8, 0.1, 0.1}
	)
	idx, err := yrand.WeightedChoice(weights)
	if err != nil {
		fmt.Println("got error:", err)
		return
	}
	fmt.Println("Luckiest Guy:", candidates[idx])
}

// This example chooses three random emojis according to associated weights.
func ExampleWeightedShuffle() {
	var (
		emojis  = []string{"🍔", "🥪", "🌮", "🌯", "🥞", "🍪", "🥮", "🥠"}
		weights = []float64{8, 7, 6, 5, 2, 2, 1, 1}
		count   = 0
	)
	err := yrand.WeightedShuffle(weights, func(idx int) (err error) {
		fmt.Print(emojis[idx], " ")
		if count++; count == 3 {
			fmt.Println()
			err = yrand.QuitShuffle
		}
		return
	})
	if err != nil {
		fmt.Println("got error:", err)
		return
	}
}
