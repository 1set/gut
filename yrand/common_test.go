package yrand

import (
	"math"
	"testing"
)

func TestIsEqualFloat(t *testing.T) {
	type args struct {
		a         float64
		b         float64
		tolerance float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"compare 0 to 0 with tole=0", args{0, 0, 0}, true},
		{"compare 0 to 0 with tole=1e-6", args{0, 0, 1e-6}, true},
		{"compare 0 to 0.01 with tole=1e-3", args{0, 0.01, 1e-3}, false},
		{"compare 0 to 0.01 with tole=1e-1", args{0, 0.01, 1e-1}, true},
		{"compare 1 to 1 with tole=0", args{1, 1, 0}, true},
		{"compare 1 to 1 with tole=1e-6", args{1, 1, 1e-6}, true},
		{"compare 1 to 1.01 with tole=1e-3", args{1, 1.01, 1e-3}, false},
		{"compare 1 to 1.01 with tole=1e-1", args{1, 1.01, 1e-1}, true},
		{"compare 1 to 2 with tole=0", args{1, 2, 0}, false},
		{"compare 1 to 2 with tole=1e-3", args{1, 2, 1e-3}, false},
		{"compare 1 to 2 with tole=1e-6", args{1, 2, 1e-6}, false},
		{"compare NaN to NaN with tole=1e-6", args{math.NaN(), math.NaN(), 1e-6}, false},
		{"compare 0 to NaN with tole=1e-6", args{0, math.NaN(), 1e-6}, false},
		{"compare +Inf to +Inf with tole=1e-6", args{math.Inf(1), math.Inf(1), 1e-6}, false},
		{"compare +Inf to -Inf with tole=1e-6", args{math.Inf(1), math.Inf(-1), 1e-6}, false},
		{"compare +Inf to NaN with tole=1e-6", args{math.Inf(1), math.NaN(), 1e-6}, false},
		{"compare -Inf to NaN with tole=1e-6", args{math.Inf(-1), math.NaN(), 1e-6}, false},
		{"compare +Inf to 0 with tole=1e-6", args{math.Inf(1), 0, 1e-6}, false},
		{"compare -Inf to 0 with tole=1e-6", args{math.Inf(-1), 0, 1e-6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqualFloat(tt.args.a, tt.args.b, tt.args.tolerance); got != tt.want {
				t.Errorf("isEqualFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsEqualFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEqualFloat(1, 2, 0.01)
	}
}

func TestIterateRandomNumbers(t *testing.T) {
	noop := func(foo uint64) {
	}
	numbers := make([]uint64, 0)
	recordNumbers := func(num uint64) {
		numbers = append(numbers, num)
	}

	type args struct {
		count    int
		max      uint64
		callback func(num uint64)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid count=0", args{0, 2, noop}, true},
		{"invalid max=1", args{8, 1, noop}, true},
		{"nil callback", args{8, 2, nil}, true},
		{"count=8 and max=2", args{8, 2, recordNumbers}, false},
		{"count=8 and max=8", args{8, 8, recordNumbers}, false},
		{"count=100 and max=16", args{100, 16, recordNumbers}, false},
		{"count=10000 and max=32", args{10000, 32, recordNumbers}, false},
		{"count=1000000 and max=32", args{1000000, 32, recordNumbers}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers = make([]uint64, 0)
			if err := iterateRandomNumbers(tt.args.count, tt.args.max, tt.args.callback); (err != nil) != tt.wantErr {
				t.Errorf("iterateRandomNumbers() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if len(numbers) != tt.args.count {
					t.Errorf("iterateRandomNumbers() number count: %v, expect: %v", len(numbers), tt.args.count)
				}
				for _, v := range numbers {
					if !(v < tt.args.max) {
						t.Errorf("iterateRandomNumbers() number should be in [0, %v), got: %v", tt.args.max, v)
					}
				}
			}
		})
	}
}

func BenchmarkIterateRandomNumbers(b *testing.B) {
	noop := func(foo uint64) {
	}
	for i := 0; i < b.N; i++ {
		iterateRandomNumbers(16, 62, noop)
	}
}
