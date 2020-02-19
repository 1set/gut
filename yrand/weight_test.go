package yrand

import (
	"errors"
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
		{"contains extremely larger weight", []float64{1e-6, 1e30, 1e-3}, false},
		{"contains extremely larger weights", []float64{1e30, 1e-6, 1e30}, false},
		{"three increasing weights", []float64{1, 100, 1000}, false},
		{"four increasing weights", []float64{2.333, 4.666, 8.888, 10.101}, false},
		{"five increasing weights", []float64{1, 2, 3, 4, 5}, false},
		{"thirty-two large number weights", getLargeWeights(32, 100000), false},
		{"1k large number weights", getLargeWeights(1024, 1000000), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() && len(tt.weights) >= 4 {
				t.Skipf("skipping large case '%v' in short mode", tt.name)
			}

			gotIdx, err := WeightedChoice(tt.weights)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeightedChoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(tt.weights) <= 32 {
				if !(0 <= gotIdx && gotIdx < len(tt.weights)) {
					t.Errorf("WeightedChoice() got invalid index = %v, want = [0, %v)", gotIdx, len(tt.weights))
					return
				}
				checkProbDist(t, "WeightedChoice", times, tt.weights, func() (idx int, err error) { return WeightedChoice(tt.weights) })
			}
		})
	}
}

func BenchmarkWeightedChoiceInvalid(b *testing.B) {
	weights := []float64{0, -10, 0, 0, -1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = WeightedChoice(weights)
	}
}

func BenchmarkWeightedChoiceValid(b *testing.B) {
	weights := []float64{2.333, 4.666, 8.888, 10.101}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = WeightedChoice(weights)
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
		{"contains extremely larger weight 1", []float64{1e-6, 1e-1, 2e-2, 1e-3, 1e30}, false},
		{"contains extremely larger weight 2", []float64{1e30, 1e-6}, false},
		{"contains extremely larger weight 3", []float64{1e30, 1e30, 1e-6}, false},
		{"contains extremely larger weight 4", []float64{1e30, 1e-6, 1e30}, false},
		{"contains extremely larger weight 5", []float64{1e-6, 1e30, 1e30}, false},
		{"contains extremely larger weight 6", []float64{1e-6, 1e30, 1e-3, 1}, false},
		{"single weight", []float64{1}, false},
		{"two diff weights", []float64{1, 3}, false},
		{"two equal weights", []float64{2, 2}, false},
		{"three diff weights", []float64{2, 4, 4}, false},
		{"contains one larger weight", []float64{1, 100, 1}, false},
		{"contains two larger weights", []float64{1, 100, 100}, false},
		{"only extremely small weights", []float64{1e-30, 2e-30, 3e-30}, false},
		{"only extremely large weights", []float64{1e30, 2e30, 3e30}, false},
		{"three increasing weights", []float64{1, 100, 1000}, false},
		{"four increasing weights", []float64{2.333, 4.666, 8.888, 10.101}, false},
		{"five increasing weights", []float64{1, 2, 3, 4, 5}, false},
		{"eight repeated weights", []float64{1, 2, 1, 2, 1, 2, 1, 2}, false},
		{"thirty-two large number weights", getLargeWeights(32, 100000), false},
		{"1k large number weights", getLargeWeights(1024, 1000000), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() && len(tt.weights) >= 4 {
				t.Skipf("skipping large case '%v' in short mode", tt.name)
			}

			cnt, maxCnt := 0, len(tt.weights)
			switch err := WeightedShuffle(tt.weights, func(idx int) (err error) {
				cnt++
				if (idx < 0) || (idx >= maxCnt) {
					t.Errorf("WeightedShuffle() got invalid index = %v, want = [0, %v)", idx, maxCnt)
				}
				return
			}); {
			case (err != nil) != tt.wantErr:
				t.Errorf("WeightedShuffle() got error = %v, wantErr = %v", err, tt.wantErr)
				return
			case (err != nil) && (err != errInvalidWeights):
				t.Errorf("WeightedShuffle() got diff error = %v, want = %v, weights = %v", err, errInvalidWeights, tt.weights)
				return
			case (err == nil) && (cnt != maxCnt):
				t.Errorf("WeightedShuffle() got not enough indexes = %v, want = %v", cnt, maxCnt)
				return
			}

			if !tt.wantErr && len(tt.weights) <= 32 {
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

	errSample := errors.New("this is sample error")
	tests2 := []struct {
		name      string
		weights   []float64
		errReturn error
		expectCnt int
		expectErr error
	}{
		{"got yield func error", []float64{1, 2, 3, 4, 5, 6, 7, 8}, errSample, 3, errSample},
		{"quit the shuffle", []float64{1, 2, 3, 4, 5, 6}, QuitShuffle, 2, nil},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			cnt := 0
			if err := WeightedShuffle(tt.weights, func(idx int) (err error) {
				cnt++
				if cnt >= tt.expectCnt {
					return tt.errReturn
				}
				return nil
			}); err != tt.expectErr {
				t.Errorf("WeightedShuffle() got error = %v, want = %v", err, tt.expectErr)
				return
			} else if cnt != tt.expectCnt {
				t.Errorf("WeightedShuffle() quit at count = %v, want = %v", cnt, tt.expectCnt)
				return
			}
		})
	}
}

func BenchmarkWeightedShuffleInvalid(b *testing.B) {
	weights := []float64{1e-6, 1e-1, 2e-2, 1e-3, 1e30}
	noop := func(idx int) (err error) { return }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeightedShuffle(weights, noop)
	}
}

func BenchmarkWeightedShuffleValid(b *testing.B) {
	weights := []float64{2.333, 4.666, 8.888, 10.101, 12.3333}
	noop := func(idx int) (err error) { return }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeightedShuffle(weights, noop)
	}
}

func getLargeWeights(count, scale int) (weights []float64) {
	for i := 1; i <= count; i++ {
		num := 99 + math.Pi*float64(scale)*math.Log2(float64(i+1))
		if i%4 == 0 {
			num = math.Log10(num)
		}
		weights = append(weights, num)
	}
	return
}

func checkProbDist(t *testing.T, name string, times int, weights []float64, idxFunc func() (idx int, err error)) {
	var (
		tolerance        = 0.2
		minExpectedTimes = 20.0
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
			result[idx]++
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
		expected := w / weightSum
		if (w <= 0) || ((expected * float64(times)) < minExpectedTimes) {
			continue
		}
		actual := float64(result[i]) / float64(times)
		if !isFloatEqual(actual, expected, tolerance) {
			t.Errorf("%s() got unexpected result, weights: %v, index:[%d](%.2f), expected: %.3f, actual: %.3f, tole: %.2f%%",
				name, weights, i, w, expected, actual, tolerance*100)
			return
		}
	}
}
