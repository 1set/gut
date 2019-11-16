package ystring

import (
	"testing"
)

func TestTrimAfterFirst(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAfterFirst(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimAfterFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
