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

var filePathMap = make(map[string]string)

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
