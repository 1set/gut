package yrand

import (
	"testing"
)

func TestInt64Range(t *testing.T) {
	type args struct {
		min int64
		max int64
	}

	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{int64(20), int64(10)}, args{int64(0), int64(1)}, true},
		{"same min and max", args{int64(10), int64(10)}, args{int64(0), int64(1)}, true},
		{"always same number", args{int64(1000), int64(1001)}, args{int64(1000), int64(1001)}, false},
		{"between [0, 100)", args{int64(0), int64(100)}, args{int64(0), int64(100)}, false},
		{"between [0, 10000)", args{int64(0), int64(10000)}, args{int64(0), int64(10000)}, false},
		{"between [0, 100000000)", args{int64(0), int64(100000000)}, args{int64(0), int64(100000000)}, false},
		{"between [0, 2147483647)", args{int64(0), int64(2147483647)}, args{int64(0), int64(2147483647)}, false},
		{"between [0, 2147483648)", args{int64(0), int64(2147483648)}, args{int64(0), int64(2147483648)}, false},
		{"between [0, 9223372036854775807)", args{int64(0), int64(9223372036854775807)}, args{int64(0), int64(9223372036854775807)}, false},
		{"between [-100, 0)", args{int64(-100), int64(0)}, args{int64(-100), int64(0)}, false},
		{"between [-9223372036854775807, 0)", args{int64(-9223372036854775807), int64(0)}, args{int64(-9223372036854775807), int64(0)}, false},
		{"between [-100, 100)", args{int64(-100), int64(100)}, args{int64(-100), int64(100)}, false},
		{"between [-2147483648, 2147483648)", args{int64(-2147483648), int64(2147483648)}, args{int64(-2147483648), int64(2147483648)}, false},
		{"between [-9223372036854775808, 9223372036854775807)", args{int64(-9223372036854775808), int64(9223372036854775807)}, args{int64(-9223372036854775808), int64(9223372036854775807)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := Int64Range(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("Int64Range() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkInt64Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64Range(int64(-100), int64(1000))
	}
}
