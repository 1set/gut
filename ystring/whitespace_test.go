package ystring

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"Empty string", emptyString, true},
		{"String contains one whitespace", oneWhitespaceString, false},
		{"String contains one tab", "\t", false},
		{"String contains whitespaces", " \t\n \t   ", false},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.s); got != tt.want {
				t.Errorf("IsEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsEmpty("lonely developer")
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"Empty string", emptyString, true},
		{"String contains one tab", "\t", true},
		{"String contains one whitespace", oneWhitespaceString, true},
		{"String contains whitespaces", " \t\n \t \f \n\v ", true},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.s); got != tt.want {
				t.Errorf("IsBlank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsBlank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsBlank("lonely developer")
	}
}

func TestShrink(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty string", args{emptyString, "."}, emptyString},
		{"Empty substring", args{"a b c", emptyString}, "abc"},
		{"String contains one tab", args{"\t", "."}, emptyString},
		{"String contains one whitespace", args{oneWhitespaceString, "."}, emptyString},
		{"String contains only whitespaces", args{" \t\n \t \f \n\v ", "."}, emptyString},
		{"String contains whitespaces with only one char", args{" \t\n \t S\f \n\v ", "."}, "S"},
		{"String contains whitespaces with two chars", args{" \t\n O\f\t\t \nK\v ", "."}, "O.K"},
		{"String contains letters and whitespaces", args{"a   b   c", "."}, "a.b.c"},
		{"String contains letters with heading&trailing whitespaces", args{"   abcdef   ", "."}, "abcdef"},
		{"String contains letters and whitespaces with heading&trailing whitespaces", args{" \t  a \t b \v c  \n ", "."}, "a.b.c"},
		{"String contains substring", args{"1234567", "345"}, "1234567"},
		{"String and substring are the same", args{"12345678", "12345678"}, "12345678"},
		{"Separator string is empty", args{"   a  b  c   d   ", emptyString}, "abcd"},
		{"Separator string contains two chars", args{"   a  \n\t b\tc   d \n\n  e  ", "=-"}, "a=-b=-c=-d=-e"},
		{"String contains emoji chars", args{" \t\n🏖️💢\t❎\t💎🕳▶️🔛\t️🈹🕞 \t \f \n\v ", emptyString}, "🏖️💢❎💎🕳▶️🔛️🈹🕞"},
		{"String and substring contains emoji chars", args{" \t\f 🏖️\v\t\n🕞 \t \f \n🥰\v ", "💌"}, "🏖️💌🕞💌🥰"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shrink(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("Shrink() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkShrink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Shrink("\t    lone  \n ly devel\t\t  \n oper \v  \f ", "~~~")
	}
}

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
		{"String contains only ASCII", args{"BenchmarkTruncate-16", 9}, "Benchmark", false},
		{"String contains only non-ASCII n=0", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 0}, emptyString, false},
		{"String contains only non-ASCII n=1", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 1}, "す", false},
		{"String contains only non-ASCII n=7", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 7}, "すきやばし次郎", false},
		{"String contains only non-ASCII n=15 (less than actual)", args{"すきやばし次郎―生涯一鮨職人🍣", 15}, "すきやばし次郎―生涯一鮨職人🍣", false},
		{"String contains only non-ASCII n=16 (equal to actual)", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 16}, "すきやばし次郎―生涯一鮨職人🍣🍱", false},
		{"String contains only non-ASCII n=100 (more than actual)", args{"すきやばし次郎―生涯一鮨職人🍣🍱", 100}, "すきやばし次郎―生涯一鮨職人🍣🍱", false},
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
