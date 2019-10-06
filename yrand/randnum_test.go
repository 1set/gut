package yrand

import (
	"testing"
)

func TestIntRange(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"same min & max", args {int64(10), int64(10)}, int64(10), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := IntRange(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("IntRange() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
