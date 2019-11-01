package yhash

import (
	"strings"
	"testing"
)

func TestStringMD5(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantStr string
		wantErr bool
	}{
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e", false},
		{"one-char string", "A", "7fc56270e7a70fa81a5935b72eacbe29", false},
		{"str=12345678", "12345678", "25d55ad283aa400af464c76d713c07ad", false},
		{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "142e80fd38631675e5f19dcc3e81dc11", false},
		{"long string", strings.Repeat("Good", 60), "7bbd7b0e70e71acbe1c4f0e67e59817e", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := StringMD5(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringMD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("StringMD5() gotStr = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func BenchmarkStringMD5(b *testing.B) {
	var content = strings.Repeat("Angel", 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringMD5(content)
	}
}

func TestStringSHA1(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantStr string
		wantErr bool
	}{
		{"empty string", "", "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
		{"one-char string", "B", "ae4f281df5a5d0ff3cad6371f76d5c29b6d953ec", false},
		{"str=123456789", "123456789", "f7c3bc1d808e04732adf679965ccc34ca7ae3441", false},
		{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "0780f2ed8873cc7ceff6d7925eea9992f6249b0f", false},
		{"long string", strings.Repeat("Good", 60), "c8fee3477eb127f23cb7dcc9d96bd6cf96987c93", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := StringSHA1(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSHA1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("StringSHA1() gotStr = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func BenchmarkStringSHA1(b *testing.B) {
	var content = strings.Repeat("Angel", 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringSHA1(content)
	}
}

func TestStringSHA224(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantStr string
		wantErr bool
	}{
		{"empty string", "", "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
		{"one-char string", "C", "484d52691fcadbfabec5a318d1cf9692c7f293cbc8c1d5f22b2d839b", false},
		{"str=123456789", "123456789", "9b3e61bf29f17c75572fae2e86e17809a4513d07c8a18152acf34521", false},
		{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "8cffb7f5b819a9131f42c67dbf8ab2f79c585e0a7d53c8948eccc435", false},
		{"long string", strings.Repeat("Good", 70), "ac664b538d59ad4df4768830da3bbaf3a71eabed4aaa86c1074a8015", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := StringSHA224(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSHA224() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("StringSHA224() gotStr = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func BenchmarkStringSHA224(b *testing.B) {
	var content = strings.Repeat("Angel", 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringSHA224(content)
	}
}
