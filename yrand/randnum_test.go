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
		{"same min and max", args{int64(10), int64(10)}, args{int64(0), int64(1)}, true},
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
