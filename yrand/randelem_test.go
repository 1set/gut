package yrand

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Parallel()
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
		if testing.Short() && tt.n >= 5 {
			t.Skipf("skipping large case '%v' in short mode", tt.name)
		}
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

func TestChoiceInt(t *testing.T) {
	t.Parallel()
	int100, _ := rangeInt(100)
	int500000, _ := rangeInt(500000)
	tests := []struct {
		name          string
		list          []int
		checkContains bool
		wantN         int
		wantErr       bool
		wantContains  bool
	}{
		{"nil list", nil, false, 0, true, false},
		{"empty list", []int{}, false, 0, true, false},
		{"single item list", []int{100}, true, 100, false, true},
		{"list of 2 - a", []int{100, 200}, true, 100, false, true},
		{"list of 2 - b", []int{100, 200}, true, 100, false, true},
		{"list of 3 - same", []int{300, 300, 300}, true, 300, false, true},
		{"list of 5 - in", []int{100, 200, 300, 400, 500}, true, 300, false, true},
		{"list of 5 - not", []int{100, 200, 300, 400, 500}, true, 350, false, false},
		{"list of 100 - in", int100, true, 50, false, true},
		{"list of 100 - not", int100, true, 100, false, false},
		{"list of 500000 - in", int500000, true, 400000, false, true},
		{"list of 500000 - not", int500000, true, 600000, false, false},
	}
	for _, tt := range tests {
		if testing.Short() && len(tt.list) >= 500 {
			t.Skipf("skipping large case '%v' in short mode", tt.name)
		}
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := ChoiceInt(tt.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChoiceInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.checkContains {
				if contains := containInt(tt.list, tt.wantN); contains != tt.wantContains {
					t.Errorf("ChoiceInt() gotN = %v, list = %v, contains actual: %v, expected: %v", gotN, tt.list, contains, tt.wantContains)
				}
			} else if gotN != tt.wantN {
				t.Errorf("ChoiceInt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func BenchmarkChoiceInt(b *testing.B) {
	int1000, _ := rangeInt(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChoiceInt(int1000)
	}
}

func TestChoiceString(t *testing.T) {
	t.Parallel()
	str100, _ := rangeString(100)
	str500000, _ := rangeString(500000)
	tests := []struct {
		name          string
		list          []string
		checkContains bool
		wantS         string
		wantErr       bool
		wantContains  bool
	}{
		{"nil list", nil, false, "", true, false},
		{"empty list", []string{}, false, "", true, false},
		{"single item list", []string{"Good"}, true, "Good", false, true},
		{"list of 2 - a", []string{"Hello", "World"}, true, "Hello", false, true},
		{"list of 2 - b", []string{"Hello", "World"}, true, "World", false, true},
		{"list of 3 - same", []string{"Yes", "Yes", "Yes"}, true, "Yes", false, true},
		{"list of 100 - in", str100, true, "50", false, true},
		{"list of 100 - not", str100, true, "100", false, false},
		{"list of 500000 - in", str500000, true, "400000", false, true},
		{"list of 500000 - not", str500000, true, "600000", false, false},
	}
	for _, tt := range tests {
		if testing.Short() && len(tt.list) >= 500 {
			t.Skipf("skipping large case '%v' in short mode", tt.name)
		}
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := ChoiceString(tt.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChoiceString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.checkContains {
				if contains := containString(tt.list, tt.wantS); contains != tt.wantContains {
					t.Errorf("ChoiceString() gotS = %v, list = %v, contains actual: %v, expected: %v", gotS, tt.list, contains, tt.wantContains)
				}
			} else if gotS != tt.wantS {
				t.Errorf("ChoiceString() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func BenchmarkChoiceString(b *testing.B) {
	str1000, _ := rangeString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChoiceString(str1000)
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

func rangeString(max int) (sl []string, err error) {
	if max < 0 {
		return nil, errors.New("max should be non-negative")
	}

	sl = make([]string, 0, max)
	for i := 0; i < max; i++ {
		sl = append(sl, strconv.Itoa(i))
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
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func containInt(list []int, n int) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}

func containString(list []string, s string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
