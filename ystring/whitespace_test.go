package ystring

import (
	"testing"
)

func TestIsEmptyOrNot(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		empty bool
	}{
		{"Empty string", emptyString, true},
		{"String contains one whitespace", oneWhitespaceString, false},
		{"String contains one tab", "\t", false},
		{"String contains whitespaces", " \t\n \t   ", false},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.s); got != tt.empty {
				t.Errorf("IsEmpty() got = %v, want %v", got, tt.empty)
			}
			if got := IsNotEmpty(tt.s); got != !tt.empty {
				t.Errorf("IsNotEmpty() got = %v, want %v", got, !tt.empty)
			}
		})
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsEmpty("lonely developer")
	}
}

func TestIsBlankOrNot(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		blank bool
	}{
		{"Empty string", emptyString, true},
		{"String contains one tab", "\t", true},
		{"String contains one whitespace", oneWhitespaceString, true},
		{"String contains whitespaces", " \t\n \t \f \n\v ", true},
		{"String contains letters", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.s); got != tt.blank {
				t.Errorf("IsBlank() got = %v, want %v", got, tt.blank)
			}
			if got := IsNotBlank(tt.s); got != !tt.blank {
				t.Errorf("IsNotBlank() got = %v, want %v", got, !tt.blank)
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
