package yrand

import (
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		alphabet string
		limit    int
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantErr    bool
	}{
		{"negative limit value", args{"ABC", -2}, 0, true},
		{"zero limit value", args{"ABC", 0}, 0, true},
		{"empty alphabet", args{"", 5}, 0, true},
		{"alphabet of single char", args{"1", 5}, 0, true},
		{"alphabet of 3 and limit of 1", args{"ABC", 1}, 1, false},
		{"alphabet of 3 and limit of 2", args{"ABC", 2}, 2, false},
		{"alphabet of 5 and limit of 8", args{"ABCDE", 8}, 8, false},
		{"alphabet of 5 and limit of 40", args{"ABCDE", 40}, 40, false},
		{"alphabet of 5 and limit of 1000", args{"ABCDE", 1000}, 1000, false},
		{"alphabet of 8 and limit of 100000", args{"ABCDEFGH", 100000}, 100000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := String(tt.args.alphabet, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotS) != tt.wantLength {
				t.Errorf("StringBase62() gotS = %v, len = %v, want %v", gotS, len(gotS), tt.wantLength)
			}
		})
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String("abcABC123", 16)
	}
}

func TestStringBase36(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		wantLength int
		wantErr    bool
	}{
		{"negative limit value", -1, 0, true},
		{"zero limit value", 0, 0, true},
		{"set limit=1", 1, 1, false},
		{"set limit=2", 2, 2, false},
		{"set limit=8", 8, 8, false},
		{"set limit=10", 10, 10, false},
		{"set limit=16", 16, 16, false},
		{"set limit=20", 20, 20, false},
		{"set limit=40", 40, 40, false},
		{"set limit=1000", 1000, 1000, false},
		{"set limit=100000", 100000, 100000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := StringBase36(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringBase36() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotS) != tt.wantLength {
				t.Errorf("StringBase36() gotS = %v, len = %v, want %v", gotS, len(gotS), tt.wantLength)
			}
		})
	}
}

func BenchmarkStringBase36(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBase36(16)
	}
}

func TestStringBase62(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		wantLength int
		wantErr    bool
	}{
		{"negative limit value", -1, 0, true},
		{"zero limit value", 0, 0, true},
		{"set limit=1", 1, 1, false},
		{"set limit=2", 2, 2, false},
		{"set limit=8", 8, 8, false},
		{"set limit=10", 10, 10, false},
		{"set limit=16", 16, 16, false},
		{"set limit=20", 20, 20, false},
		{"set limit=40", 40, 40, false},
		{"set limit=1000", 1000, 1000, false},
		{"set limit=100000", 100000, 100000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := StringBase62(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringBase62() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotS) != tt.wantLength {
				t.Errorf("StringBase62() gotS = %v, len = %v, want %v", gotS, len(gotS), tt.wantLength)
			}
		})
	}
}

func BenchmarkStringBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBase62(16)
	}
}
