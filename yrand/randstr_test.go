package yrand

import (
	"testing"
)

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
