package yhash

import (
	"encoding/base64"
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
	tempFileContents := map[string]string{
		"empty file":           "",
		"one-line text file":   "Hello World",
		"large text file":      strings.Repeat("Stop managing your time. Start managing your focus. ", 10000),
		"xlarge text file":     strings.Repeat("Do or do not, there is no try. ", 1000000),
		"small binary file":    "R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
		"another small binary": "VGhlIHF1aWNrIGJyb3duIPCfpooganVtcHMgb3ZlciAxMyBsYXp5IPCfkLYuIOKAnFRoZXJlIHdhcyBhIHNtYWxsIG51bWJlciBvZiByZWFsbHkgc21hcnQsIHJlYWxseSB5b3VuZyBjb2RlcnMgd2hvIHByb2R1Y2VkIGEgbG90IG9mIHZlcnkgY2xldmVyIGNvZGUgdGhhdCBvbmx5IHRoZXkgY291bGQgdW5kZXJzdGFuZCzigJ0gc2FpZCB2YW4gUm9zc3VtLiDigJxUaGF0IGlzIHByb2JhYmx5IHRoZSByaWdodCBhdHRpdHVkZSB0byBoYXZlIHdoZW4geW91J3JlIGEgcmVhbGx5IHNtYWxsIHN0YXJ0dXAu4oCd",
	}
	for title, content := range tempFileContents {
		name := strings.ReplaceAll(title, " ", "_")
		if file, err := ioutil.TempFile("", name); err == nil {
			if strings.Contains(title, "binary") {
				data, err := base64.StdEncoding.DecodeString(content)
				if err != nil {
					continue
				}
				if _, err = file.Write(data); err != nil {
					continue
				}
			} else if _, err = file.WriteString(content); err != nil {
				continue
			}

			filePathMap[title] = file.Name()
			_ = file.Close()
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
		{"xlarge text file", "", "0d9d7d9349c970fbf71b46698c5d1165", false},
		{"small binary file", "", "f837aa60b6fe83458f790db60d529fc9", false},
		{"another small binary", "", "89e19cf9c9680994d75adfac30887b31", false},
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

func TestFileSHA1(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{"file not found", "__FILE__NOT__EXIST__", "", true},
		{"empty file", "", "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
		{"one-line text file", "", "0a4d55a8d778e5022fab701977c5d840bbc486d0", false},
		{"large text file", "", "22c538b3b8ba9a9f817eb05d463a0cb7ba6a9625", false},
		{"xlarge text file", "", "3912642ec8f1430ae3f1f870d2279a26a0a02297", false},
		{"small binary file", "", "14af87ccec7f81bb28d53c84da2fd5a9d5925cda", false},
		{"another small binary", "", "cd2ad24ee27178115ff6440bb3f996a142888838", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := tt.filePath
			if len(tt.filePath) == 0 {
				var found bool
				if filePath, found = filePathMap[tt.name]; !found {
					t.Errorf("FileSHA1() got no file for case '%v'", tt.name)
					return
				}
			}

			got, err := FileSHA1(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSHA1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileSHA1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFileSHA1(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileSHA1() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FileSHA1(path)
	}
}
