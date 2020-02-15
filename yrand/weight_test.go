package yrand

import (
	"fmt"
	"math"
	"testing"
)

func TestWeightedChoice(t *testing.T) {
	var (
		times = 300000
	)
	tests := []struct {
		name    string
		weights []float64
		wantErr bool
	}{
		{"nil", nil, true},
		{"empty weights", []float64{}, true},
		{"only zero weights", []float64{0, 0}, true},
		{"only non-positive weights", []float64{0, 0, -1}, true},
		{"single weight", []float64{1}, false},
		{"two diff weights", []float64{1, 3}, false},
		{"two equal weights", []float64{2, 2}, false},
		{"three diff weights", []float64{2, 4, 4}, false},
		{"contains one larger weight", []float64{1, 100, 1}, false},
		{"contains two larger weights", []float64{1, 100, 100}, false},
		{"contains non-positive weight", []float64{10, 0, 10}, false},
		{"contains non-positive weights", []float64{-1, 10, 0}, false},
		{"three increasing weights", []float64{1, 100, 1000}, false},
		{"four increasing weights", []float64{2.333, 4.666, 8.888, 10.101}, false},
		{"five increasing weights", []float64{1, 2, 3, 4, 5}, false},
		{"contains extremely larger weight", []float64{1e-6, 1e30, 1e-3, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, err := WeightedChoice(tt.weights)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeightedChoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !(0 <= gotIdx && gotIdx < len(tt.weights)) {
					t.Errorf("WeightedChoice() got invalid index = %v, want = [0, %v)", gotIdx, len(tt.weights))
					return
				}
				checkProbDist(t, "WeightedChoice", times, tt.weights, func() (idx int, err error) { return WeightedChoice(tt.weights) })
			}
		})
	}
}

func TestWeightedShuffle(t *testing.T) {
	var (
		times = 300000
	)
	tests := []struct {
		name    string
		weights []float64
		wantErr bool
	}{
		{"nil", nil, true},
		{"empty weights", []float64{}, true},
		{"only zero weights", []float64{0, 0}, true},
		{"only non-positive weights", []float64{0, 0, -1}, true},
		{"contains non-positive weights", []float64{-1, 10, 0}, true},
		{"contains extremely larger weight", []float64{1e-6, 1e30, 1e-3, 1}, true},
		{"single weight", []float64{1}, false},
		{"two diff weights", []float64{1, 3}, false},
		{"two equal weights", []float64{2, 2}, false},
		{"three diff weights", []float64{2, 4, 4}, false},
		{"contains one larger weight", []float64{1, 100, 1}, false},
		{"contains two larger weights", []float64{1, 100, 100}, false},
		{"three increasing weights", []float64{1, 100, 1000}, false},
		{"four increasing weights", []float64{2.333, 4.666, 8.888, 10.101}, false},
		{"five increasing weights", []float64{1, 2, 3, 4, 5}, false},
		{"six repeated weights", []float64{1, 2, 1, 2, 1, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WeightedShuffle(tt.weights, func(idx int) (err error) {
				return
			}); (err != nil) != tt.wantErr {
				t.Errorf("WeightedShuffle() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				checkProbDist(t, "WeightedShuffle", times, tt.weights, func() (idx int, err error) {
					savedIdx, cnt := 0, 0
					err = WeightedShuffle(tt.weights, func(idx int) error {
						if cnt == 0 {
							savedIdx = idx
						}
						if !(0 <= idx && idx < len(tt.weights)) {
							return fmt.Errorf("invalid index: %v", idx)
						}
						cnt++
						return nil
					})
					return savedIdx, err
				})
			}
		})
	}
}

func checkProbDist(t *testing.T, name string, times int, weights []float64, idxFunc func() (idx int, err error)) {
	var (
		tolerance        = 0.15
		minExpectedTimes = 10.0
	)

	result := make(map[int]int)
	for i := 1; i <= times; i++ {
		idx, err := idxFunc()
		if err != nil {
			t.Errorf("%s() got unexpected error = %v", name, err)
			return
		}
		if !(0 <= idx && idx < len(weights)) {
			t.Errorf("%s() got invalid index = %v, want = [0, %v)", name, idx, len(weights))
			return
		}

		if _, ok := result[idx]; ok {
			result[idx] += 1
		} else {
			result[idx] = 1
		}
	}

	weightSum := 0.0
	for _, w := range weights {
		if w > 0 {
			weightSum += w
		}
	}

	for i, w := range weights {
		if (w <= 0) || ((w / weightSum * float64(times)) < minExpectedTimes) {
			continue
		}
		expected := w / weightSum
		actual := float64(result[i]) / float64(times)
		diff := math.Abs(actual/expected - 1)
		if diff > tolerance {
			t.Errorf("%s() got unexpected result, weights: %v, index:[%d](%.2f), expected: %.3f, actual: %.3f, diff: %.2f%%, tole: %.2f%%",
				name, weights, i, w, expected, actual, diff*100, tolerance*100)
			return
		}
	}
}
