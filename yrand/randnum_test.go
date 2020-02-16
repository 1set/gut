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
		{"choose from 0/1", args{int64(0), int64(2)}, args{int64(0), int64(2)}, false},
		{"between [0, 100)", args{int64(0), int64(100)}, args{int64(0), int64(100)}, false},
		{"between [0, 10000)", args{int64(0), int64(10000)}, args{int64(0), int64(10000)}, false},
		{"between [0, 100000000)", args{int64(0), int64(100000000)}, args{int64(0), int64(100000000)}, false},
		{"between [0, 2147483647)", args{int64(0), int64(2147483647)}, args{int64(0), int64(2147483647)}, false},
		{"between [0, 2147483648)", args{int64(0), int64(2147483648)}, args{int64(0), int64(2147483648)}, false},
		{"between [0, 2147483649)", args{int64(0), int64(2147483649)}, args{int64(0), int64(2147483649)}, false},
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
		Int64Range(int64(0), int64(2147483649))
	}
}

func TestInt32Range(t *testing.T) {
	type args struct {
		min int32
		max int32
	}
	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{int32(20), int32(10)}, args{int32(0), int32(1)}, true},
		{"same min and max", args{int32(10), int32(10)}, args{int32(0), int32(1)}, true},
		{"always same number", args{int32(1000), int32(1001)}, args{int32(1000), int32(1001)}, false},
		{"choose from 0/1", args{int32(0), int32(2)}, args{int32(0), int32(2)}, false},
		{"between [0, 100)", args{int32(0), int32(100)}, args{int32(0), int32(100)}, false},
		{"between [0, 10000)", args{int32(0), int32(10000)}, args{int32(0), int32(10000)}, false},
		{"between [0, 100000000)", args{int32(0), int32(100000000)}, args{int32(0), int32(100000000)}, false},
		{"between [0, 2147483647)", args{int32(0), int32(2147483647)}, args{int32(0), int32(2147483647)}, false},
		{"between [-100, 0)", args{int32(-100), int32(0)}, args{int32(-100), int32(0)}, false},
		{"between [-2147483647, 0)", args{int32(-2147483647), int32(0)}, args{int32(-2147483647), int32(0)}, false},
		{"between [-100, 100)", args{int32(-100), int32(100)}, args{int32(-100), int32(100)}, false},
		{"between [-2147483648, 2147483647)", args{int32(-2147483648), int32(2147483647)}, args{int32(-2147483648), int32(2147483647)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := Int32Range(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int32Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("Int32Range() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkInt32Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32Range(int32(0), int32(1073741825))
	}
}

func TestIntRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{20, 10}, args{0, 1}, true},
		{"same min and max", args{10, 10}, args{0, 1}, true},
		{"always same number", args{1000, 1001}, args{1000, 1001}, false},
		{"choose from 0/1", args{0, 2}, args{0, 2}, false},
		{"between [0, 100)", args{0, 100}, args{0, 100}, false},
		{"between [0, 10000)", args{0, 10000}, args{0, 10000}, false},
		{"between [0, 100000000)", args{0, 100000000}, args{0, 100000000}, false},
		{"between [0, 2147483647)", args{0, 2147483647}, args{0, 2147483647}, false},
		{"between [0, 2147483648)", args{0, 2147483648}, args{0, 2147483648}, false},
		{"between [0, 2147483649)", args{0, 2147483649}, args{0, 2147483649}, false},
		{"between [0, 9223372036854775807)", args{0, 9223372036854775807}, args{0, 9223372036854775807}, false},
		{"between [-100, 0)", args{-100, 0}, args{-100, 0}, false},
		{"between [-9223372036854775807, 0)", args{-9223372036854775807, 0}, args{-9223372036854775807, 0}, false},
		{"between [-100, 100)", args{-100, 100}, args{-100, 100}, false},
		{"between [-2147483648, 2147483648)", args{-2147483648, 2147483648}, args{-2147483648, 2147483648}, false},
		{"between [-9223372036854775808, 9223372036854775807)", args{-9223372036854775808, 9223372036854775807}, args{-9223372036854775808, 9223372036854775807}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := IntRange(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("IntRange() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntRange(0, 2147483649)
	}
}

func TestUint64Range(t *testing.T) {
	type args struct {
		min uint64
		max uint64
	}
	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{uint64(20), uint64(10)}, args{uint64(0), uint64(1)}, true},
		{"same min and max", args{uint64(10), uint64(10)}, args{uint64(0), uint64(1)}, true},
		{"always same number", args{uint64(1000), uint64(1001)}, args{uint64(1000), uint64(1001)}, false},
		{"choose from 0/1", args{uint64(0), uint64(2)}, args{uint64(0), uint64(2)}, false},
		{"between [0, 100)", args{uint64(0), uint64(100)}, args{uint64(0), uint64(100)}, false},
		{"between [0, 10000)", args{uint64(0), uint64(10000)}, args{uint64(0), uint64(10000)}, false},
		{"between [0, 100000000)", args{uint64(0), uint64(100000000)}, args{uint64(0), uint64(100000000)}, false},
		{"between [0, 2147483647)", args{uint64(0), uint64(2147483647)}, args{uint64(0), uint64(2147483647)}, false},
		{"between [0, 2147483648)", args{uint64(0), uint64(2147483648)}, args{uint64(0), uint64(2147483648)}, false},
		{"between [0, 2147483649)", args{uint64(0), uint64(2147483649)}, args{uint64(0), uint64(2147483649)}, false},
		{"between [0, 9223372036854775807)", args{uint64(0), uint64(9223372036854775807)}, args{uint64(0), uint64(9223372036854775807)}, false},
		{"between [0, 18446744073709551615)", args{uint64(0), uint64(18446744073709551615)}, args{uint64(0), uint64(18446744073709551615)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := Uint64Range(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("Uint64Range() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkUint64Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64Range(uint64(2), uint64(2147483659))
	}
}

func TestUint32Range(t *testing.T) {
	type args struct {
		min uint32
		max uint32
	}
	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{uint32(20), uint32(10)}, args{uint32(0), uint32(1)}, true},
		{"same min and max", args{uint32(10), uint32(10)}, args{uint32(0), uint32(1)}, true},
		{"always same number", args{uint32(1000), uint32(1001)}, args{uint32(1000), uint32(1001)}, false},
		{"choose from 0/1", args{uint32(0), uint32(2)}, args{uint32(0), uint32(2)}, false},
		{"between [0, 100)", args{uint32(0), uint32(100)}, args{uint32(0), uint32(100)}, false},
		{"between [0, 10000)", args{uint32(0), uint32(10000)}, args{uint32(0), uint32(10000)}, false},
		{"between [0, 100000000)", args{uint32(0), uint32(100000000)}, args{uint32(0), uint32(100000000)}, false},
		{"between [0, 2147483647)", args{uint32(0), uint32(2147483647)}, args{uint32(0), uint32(2147483647)}, false},
		{"between [0, 2147483648)", args{uint32(0), uint32(2147483648)}, args{uint32(0), uint32(2147483648)}, false},
		{"between [0, 2147483649)", args{uint32(0), uint32(2147483649)}, args{uint32(0), uint32(2147483649)}, false},
		{"between [0, 4294967295)", args{uint32(0), uint32(4294967295)}, args{uint32(0), uint32(4294967295)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := Uint32Range(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("Uint32Range() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkUint32Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32Range(uint32(2), uint32(2147483659))
	}
}

func TestUintRange(t *testing.T) {
	type args struct {
		min uint
		max uint
	}
	tests := []struct {
		name       string
		args       args
		wantNRange args
		wantErr    bool
	}{
		{"invalid min and max", args{uint(20), uint(10)}, args{uint(0), uint(1)}, true},
		{"same min and max", args{uint(10), uint(10)}, args{uint(0), uint(1)}, true},
		{"always same number", args{uint(1000), uint(1001)}, args{uint(1000), uint(1001)}, false},
		{"choose from 0/1", args{uint(0), uint(2)}, args{uint(0), uint(2)}, false},
		{"between [0, 100)", args{uint(0), uint(100)}, args{uint(0), uint(100)}, false},
		{"between [0, 10000)", args{uint(0), uint(10000)}, args{uint(0), uint(10000)}, false},
		{"between [0, 100000000)", args{uint(0), uint(100000000)}, args{uint(0), uint(100000000)}, false},
		{"between [0, 2147483647)", args{uint(0), uint(2147483647)}, args{uint(0), uint(2147483647)}, false},
		{"between [0, 2147483648)", args{uint(0), uint(2147483648)}, args{uint(0), uint(2147483648)}, false},
		{"between [0, 2147483649)", args{uint(0), uint(2147483649)}, args{uint(0), uint(2147483649)}, false},
		{"between [0, 9223372036854775807)", args{uint(0), uint(9223372036854775807)}, args{uint(0), uint(9223372036854775807)}, false},
		{"between [0, 18446744073709551615)", args{uint(0), uint(18446744073709551615)}, args{uint(0), uint(18446744073709551615)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := UintRange(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("UintRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(tt.wantNRange.min <= gotN && gotN < tt.wantNRange.max) {
				t.Errorf("UintRange() gotN = %v, want %v", gotN, tt.wantNRange)
			}
		})
	}
}

func BenchmarkUintRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UintRange(uint(2), uint(2147483659))
	}
}

func TestFloat64(t *testing.T) {
	count := 100000
	total := 0.0
	for i := 0; i < count; i++ {
		gotN, err := Float64()
		total += gotN
		if err != nil {
			t.Errorf("Float64() got error = %v", err)
			return
		}
		if !(0 <= gotN && gotN < 1) {
			t.Errorf("Float64() got N = %v", gotN)
			return
		}
	}

	avg := total / float64(count)
	if !(isFloatEqual(avg, 0.5, 0.01)) {
		t.Errorf("Float64() got unexpected average = %v", avg)
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

func TestFloat32(t *testing.T) {
	count := 100000
	total := float32(0.0)
	for i := 0; i < count; i++ {
		gotN, err := Float32()
		total += gotN
		if err != nil {
			t.Errorf("Float32() got error = %v", err)
			return
		}
		if !(0 <= gotN && gotN < 1) {
			t.Errorf("Float32() got N = %v", gotN)
			return
		}
	}

	avg := total / float32(count)
	if !(isFloatEqual(float64(avg), 0.5, 0.01)) {
		t.Errorf("Float32() got unexpected average = %v", avg)
	}
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32()
	}
}
