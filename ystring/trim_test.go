package ystring

import (
	"testing"
)

var (
	oneCharString    = "A"
	twoCharsString   = "KO"
	threeCharsString = "Luv"
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
		{"Empty string", args{emptyString, threeCharsString}, emptyString},
		{"Empty substring", args{threeCharsString, emptyString}, threeCharsString},
		{"Empty string and substring", args{emptyString, emptyString}, emptyString},
		{"Same string and substring", args{threeCharsString, threeCharsString}, emptyString},
		{"Equal length with diff content", args{"ABC", "ABD"}, "ABC"},
		{"Substring contains string", args{"ABC", "ABCDE"}, "ABC"},
		{"String starts with substring", args{"AppleComputer", "Apple"}, emptyString},
		{"String ends with substring", args{"AppleComputer", "Computer"}, "Apple"},
		{"String doesn't contain substring", args{"AppleComputer", "Banana"}, "AppleComputer"},
		{"String contains multiple substring", args{"Long, long ago, long ago, long ago, long ago.", "long ago"}, "Long, "},
		{"String and substring contains non-ASCII", args{"我真的非常非常感谢你🤙", "非常"}, "我真的"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAfterFirst(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimAfterFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimAfterFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimAfterFirst("Long, long ago, long ago, long ago, long ago.", "long ago")
	}
}
