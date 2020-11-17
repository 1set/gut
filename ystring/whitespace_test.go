package ystring

import (
	"testing"
)

func TestIsEmptyOrNot(t *testing.T) {
	t.Parallel()
	fallback := "this is fallback string"
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

			if tt.empty {
				if got := NotEmptyOrDefault(tt.s, fallback); got != fallback {
					t.Errorf("NotEmptyOrDefault() got = %q, want %q", got, fallback)
				}
			} else {
				if got := NotEmptyOrDefault(tt.s, fallback); got != tt.s {
					t.Errorf("NotEmptyOrDefault() got = %q, want %q", got, tt.s)
				}
			}
		})
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsEmpty("lovely developer")
	}
}

func BenchmarkIsNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNotEmpty("lovely developer")
	}
}

func BenchmarkNotEmptyOrDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NotEmptyOrDefault("lovely developer", "fallback default value")
	}
}

func TestIsBlankOrNot(t *testing.T) {
	t.Parallel()
	fallback := "this is fallback string"
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

			if tt.blank {
				if got := NotBlankOrDefault(tt.s, fallback); got != fallback {
					t.Errorf("NotBlankOrDefault() got = %q, want %q", got, fallback)
				}
			} else {
				if got := NotBlankOrDefault(tt.s, fallback); got != tt.s {
					t.Errorf("NotBlankOrDefault() got = %q, want %q", got, tt.s)
				}
			}
		})
	}
}

func BenchmarkIsBlank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsBlank("lovely developer")
	}
}

func BenchmarkIsNotBlank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNotBlank("lovely developer")
	}
}

func BenchmarkNotBlankOrDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NotBlankOrDefault("lovely developer", "fallback default value")
	}
}

func TestShrink(t *testing.T) {
	t.Parallel()
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
		_ = Shrink("\t    love  \n ly devel\t\t  \n oper \v  \f ", "~~~")
	}
}
