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
		{"Empty string", args{emptyString, threeCharsString}, emptyString},
		{"Empty substring", args{threeCharsString, emptyString}, threeCharsString},
		{"Empty string and substring", args{emptyString, emptyString}, emptyString},
		{"Same string and substring", args{threeCharsString, threeCharsString}, emptyString},
		{"Equal length with diff content", args{"ABC", "ABD"}, "ABC"},
		{"Substring contains string", args{"ABC", "ABCDE"}, "ABC"},
		{"String starts with substring", args{"AppleComputer", "Apple"}, emptyString},
		{"String ends with substring", args{"AppleComputer", "Computer"}, "Apple"},
		{"String doesn't contain substring", args{"AppleComputer", "Banana"}, "AppleComputer"},
		{"String contains single substring", args{"What a wonderful world!", "wonderful"}, "What a "},
		{"String contains one-char substring", args{"abc.def.ghi.txt", "."}, "abc"},
		{"String contains multiple substring", args{"Long, long ago, long ago, long ago, long ago.", "long ago"}, "Long, "},
		{"String and substring contains non-ASCII", args{"我真的非常非常感谢你🤙", "非常"}, "我真的"},
		{"String and substring are full of emojis", args{"💟🤙⏭️✔️🔠🏖️💢❎💎🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒", "💎"}, "💟🤙⏭️✔️🔠🏖️💢❎"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAfterFirst(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimAfterFirst() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimAfterFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimAfterFirst("Long, long ago, long ago, long ago, long ago.", "long ago")
	}
}

func TestTrimAfterLast(t *testing.T) {
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
		{"String contains single substring", args{"What a wonderful world!", "wonderful"}, "What a "},
		{"String contains one-char substring", args{"abc.def.ghi.txt", "."}, "abc.def.ghi"},
		{"String contains multiple substring", args{"Long, long ago, long ago, long ago, long ago.", "long ago"}, "Long, long ago, long ago, long ago, "},
		{"String and substring contains non-ASCII", args{"我真的非常非常感谢你🤙", "非常"}, "我真的非常"},
		{"String and substring are full of emojis", args{"💟🤙⏭️✔️🔠🏖️💢❎💎🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒", "💎"}, "💟🤙⏭️✔️🔠🏖️💢❎"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAfterLast(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimAfterFirst() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimAfterLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimAfterLast("Long, long ago, long ago, long ago, long ago.", "long ago")
	}
}

func TestTrimBeforeFirst(t *testing.T) {
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
		{"String starts with substring", args{"AppleComputer", "Apple"}, "Computer"},
		{"String ends with substring", args{"AppleComputer", "Computer"}, emptyString},
		{"String doesn't contain substring", args{"AppleComputer", "Banana"}, "AppleComputer"},
		{"String contains single substring", args{"What a wonderful world!", "wonderful"}, " world!"},
		{"String contains one-char substring", args{"abc.def.ghi.txt", "."}, "def.ghi.txt"},
		{"String contains multiple substring", args{"Long, long ago, long ago, long ago, long ago.", "long ago"}, ", long ago, long ago, long ago."},
		{"String and substring contains non-ASCII", args{"我真的非常非常感谢你🤙", "非常"}, "非常感谢你🤙"},
		{"String and substring are full of emojis", args{"💟🤙⏭️✔️🔠🏖️💢❎💎🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒", "💎"}, "🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimBeforeFirst(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimBeforeFirst() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimBeforeFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimBeforeFirst("Long, long ago, long ago, long ago, long ago.", "long ago")
	}
}

func TestTrimBeforeLast(t *testing.T) {
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
		{"String starts with substring", args{"AppleComputer", "Apple"}, "Computer"},
		{"String ends with substring", args{"AppleComputer", "Computer"}, emptyString},
		{"String doesn't contain substring", args{"AppleComputer", "Banana"}, "AppleComputer"},
		{"String contains single substring", args{"What a wonderful world!", "wonderful"}, " world!"},
		{"String contains one-char substring", args{"abc.def.ghi.txt", "."}, "txt"},
		{"String contains multiple substring", args{"Long, long ago, long ago, long ago, long ago.", "long ago"}, "."},
		{"String and substring contains non-ASCII", args{"我真的非常非常感谢你🤙", "非常"}, "感谢你🤙"},
		{"String and substring are full of emojis", args{"💟🤙⏭️✔️🔠🏖️💢❎💎🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒", "💎"}, "🕳▶️🔛️🈹🕞🇧🇪🆎🔉☑️🚫⏏️💠🍒"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimBeforeLast(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimBeforeLast() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimBeforeLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimBeforeLast("Long, long ago, long ago, long ago, long ago.", "long ago")
	}
}
