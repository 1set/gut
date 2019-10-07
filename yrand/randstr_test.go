package yrand

import (
	"fmt"
	"testing"
)

func TestStringBase62(t *testing.T) {
	s, e := StringBase62(10)
	fmt.Println(s, e)
	return

	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := StringBase62(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringBase62() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("StringBase62() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}