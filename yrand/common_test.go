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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := iterateRandomNumbers(tt.args.count, tt.args.max, tt.args.callback); (err != nil) != tt.wantErr {
				t.Errorf("iterateRandomNumbers() error = %v, wantErr %v", err, tt.wantErr)
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
