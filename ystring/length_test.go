package ystring

import (
	"testing"
)

func TestLength(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Empty string", emptyString, 0},
		{"String contains one tab", "\t", 1},
		{"String contains two hanjis", "沢愛", 2},
		{"String contains three whitespaces", "\t \n", 3},
		{"String contains four emojis", "👍😍🌍🔥", 4},
		{"String contains five combined emojis", "😍✍🏻🗣️🕳️🤏", 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.s); got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Length("\t  🏖️💢\t❎lone  \n ly devel\t🌍\t  💎🕳▶️🔛\t️🈹🕞  \n oper \v  \f  ~~~")
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Negative n", args{"abc", -2}, emptyString, true},
		{"Zero n", args{"abc", 0}, emptyString, false},
		{"One-char string with n=1", args{"A", 1}, "A", false},
		{"One-char string with n=2", args{"A", 2}, "A", false},
		{"One-char hanji string with n=1", args{"倉", 1}, "倉", false},
		{"One-char hanji string with n=2", args{"倉", 2}, "倉", false},
		{"String contains only ASCII n=0", args{"BenchmarkTruncate-16", 0}, emptyString, false},
		{"String contains only ASCII n=1", args{"BenchmarkTruncate-16", 1}, "B", false},
		{"String contains only ASCII n=9", args{"BenchmarkTruncate-16", 9}, "Benchmark", false},
		{"String contains only ASCII n=19", args{"BenchmarkTruncate-16", 19}, "BenchmarkTruncate-1", false},
		{"String contains only ASCII n=20", args{"BenchmarkTruncate-16", 20}, "BenchmarkTruncate-16", false},
		{"String contains only ASCII n=21", args{"BenchmarkTruncate-16", 21}, "BenchmarkTruncate-16", false},
		{"String contains only ASCII n=100", args{"BenchmarkTruncate-16", 100}, "BenchmarkTruncate-16", false},
		{"String contains only non-ASCII n=0", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 0}, emptyString, false},
		{"String contains only non-ASCII n=1", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 1}, "す", false},
		{"String contains only non-ASCII n=7", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 7}, "すきやばし次郎", false},
		{"String contains only non-ASCII n=15", args{"すきやばし次郎―生涯一鮨職人🍣", 15}, "すきやばし次郎―生涯一鮨職人🍣", false},
		{"String contains only non-ASCII n=16", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 16}, "すきやばし次郎―生涯一鮨職人🍣🍱", false},
		{"String contains only non-ASCII n=17", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 17}, "すきやばし次郎―生涯一鮨職人🍣🍱", false},
		{"String contains only non-ASCII n=100", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 100}, "すきやばし次郎―生涯一鮨職人🍣🍱", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantErr {
					t.Errorf("Truncate() panic = %v, wantErr %v", r, tt.wantErr)
				}
			}()

			if got := Truncate(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Truncate() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkTruncate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Truncate("\t  🏖️💢\t❎lone  \n ly devel\t🌍\t  💎🕳▶️🔛\t️🈹🕞  \n oper \v  \f  ~~~", 10)
	}
}
