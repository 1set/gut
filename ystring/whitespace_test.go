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
		{"Empty string", "", true},
		{"String contains one whitespace", " ", false},
		{"String contains one tab", "\t", false},
		{"String contains whitespaces", " \t\n \t   ", false},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
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
		{"Empty string", "", true},
		{"String contains one whitespace", " ", true},
		{"String contains one tab", "\t", true},
		{"String contains whitespaces", " \t\n \t \f \n\v ", true},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.s); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
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
		{"Empty string", args{"", "."}, ""},
		{"String contains one whitespace", args{" ", "."}, ""},
		{"String contains one tab", args{"\t", "."}, ""},
		{"String contains only whitespaces", args{" \t\n \t \f \n\v ", "."}, ""},
		{"String contains letters and whitespaces", args{"a   b   c", "."}, "a.b.c"},
		{"String contains letters with heading&trailing whitespaces", args{"   abcdef   ", "."}, "abcdef"},
		{"String contains letters and whitespaces with heading&trailing whitespaces", args{" \t  a \t b \v c  \n ", "."}, "a.b.c"},
		{"Separator string is empty", args{"   a  b  c   d   ", ""}, "abcd"},
		{"Separator string contains two chars", args{"   a  \n\t b\tc   d \n\n  e  ", "=-"}, "a=-b=-c=-d=-e"},
		{"String contains emoji chars", args{" \t\n🏖️💢\t❎\t💎🕳▶️🔛\t️🈹🕞 \t \f \n\v ", ""}, "🏖️💢❎💎🕳▶️🔛️🈹🕞"},
		{"String and substring contains emoji chars", args{" \t\f 🏖️\v\t\n🕞 \t \f \n🥰\v ", "💌"}, "🏖️💌🕞💌🥰"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shrink(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("Shrink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkShrink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Shrink("\t    lone  \n ly devel\t\t  \n oper \v  \f ", "~~~")
	}
}
