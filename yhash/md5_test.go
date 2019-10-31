package yhash

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

var filePathMap = make(map[string]string, 0)

func setup() {
	testFileContents := map[string]string{
		"empty file":         "",
		"one-line text file": "Hello World",
		"large text file":    strings.Repeat("Stop managing your time. Start managing your focus. ", 10000),
	}

	for title, content := range testFileContents {
		name := strings.ReplaceAll(title, " ", "_")
		if file, err := ioutil.TempFile("", name); err == nil {
			if _, err = file.WriteString(content); err == nil {
				filePathMap[title] = file.Name()
				_ = file.Close()
			}
		}
	}
}

func teardown() {
	for _, path := range filePathMap {
		os.Remove(path)
	}
}

func TestFileMD5(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{"file not found", "__FILE__NOT__EXIST__", "", true},
		{"empty file", "", "d41d8cd98f00b204e9800998ecf8427e", false},
		{"one-line text file", "", "b10a8db164e0754105b7a99be72e3fe5", false},
		{"large text file", "", "3094ffc905b6a832d68ca27c86d52dc0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := tt.filePath
			if len(tt.filePath) == 0 {
				var found bool
				if filePath, found = filePathMap[tt.name]; !found {
					t.Errorf("FileMD5() got no file for case '%v'", tt.name)
					return
				}
			}

			got, err := FileMD5(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileMD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileMD5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFileMD5(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileMD5() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FileMD5(path)
	}
}

func TestStringMD5(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantStr string
		wantErr bool
	}{
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e", false},
		{"one-char string", "A", "7fc56270e7a70fa81a5935b72eacbe29", false},
		{"str=123456", "12345678", "25d55ad283aa400af464c76d713c07ad", false},
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

func TestBytesMD5(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantStr string
		wantErr bool
	}{
		{"nil", nil, "d41d8cd98f00b204e9800998ecf8427e", false},
		{"empty", []byte{}, "d41d8cd98f00b204e9800998ecf8427e", false},
		{"one zero", []byte{0}, "93b885adfe0da089cdf634904fd59f71", false},
		{"one byte", []byte{88}, "02129bb861061d1a052c592e2dc6b383", false},
		{"two bytes", []byte{88, 89}, "74c53bcd3dcb2bb79993b2fec37d362a", false},
		{"three bytes", []byte{88, 89, 90}, "e65075d550f9b5bf9992fa1d71a131be", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := BytesMD5(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesMD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("BytesMD5() gotStr = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func BenchmarkBytesMD5(b *testing.B) {
	n := 1000
	data := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, byte(i%120))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BytesMD5(data)
	}
}
