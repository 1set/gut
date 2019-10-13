package yrand

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkShuffleNoop(b *testing.B) {
	count := 1000
	noopNoop := func(i, j int) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(count, noopNoop)
	}
}

func BenchmarkShuffleWithSlice(b *testing.B) {
	count := 1000
	num, _ := rangeInt(count)
	swapFunc := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(count, swapFunc)
	}
}

func TestShuffle(t *testing.T) {
	num, _ := rangeInt(0)
	swapFunc := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}

	tests := []struct {
		name     string
		n        int
		expected int
		wantErr  bool
	}{
		{"n = -7", -7, 0, true},
		{"n = 0", 0, 0, false},
		{"n = 1", 1, factorial(1), false},
		{"n = 2", 2, factorial(2), false},
		{"n = 3", 3, factorial(3), false},
		{"n = 4", 4, factorial(4), false},
		{"n = 5", 5, factorial(5), false},
		{"n = 6", 6, factorial(6), false},
		{"n = 7", 7, factorial(7), false},
		{"n = 8", 8, factorial(8), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			times := 1 + tt.expected*20
			counters := map[string]int{}
			for i := 0; i < times; i++ {
				num, _ = rangeInt(tt.n)
				err := Shuffle(tt.n, swapFunc)
				if (err != nil) != tt.wantErr {
					t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
				}
				if str := numSlice2String(num); len(str) > 0 {
					counters[str]++
				}
			}
			if actual := len(counters); actual != tt.expected {
				t.Errorf("Shuffle() order count: %v, expected: %v", actual, tt.expected)
			}
		})
	}
}

func rangeInt(max int) (sl []int, err error) {
	if max < 0 {
		return nil, errors.New("max should be non-negative")
	}

	sl = make([]int, 0, max)
	for i := 0; i < max; i++ {
		sl = append(sl, i)
	}
	return
}

func numSlice2String(num []int) (s string) {
	if len(num) == 0 {
		return
	}

	str := make([]string, 0, len(num))
	for _, v := range num {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, "|")
}

func factorial(n int) int {
	res := int(1)
	for i := int(2); i <= n; i++ {
		res *= i
	}
	return res
}
