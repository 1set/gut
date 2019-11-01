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

func BenchmarkFileMD5(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileMD5() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileMD5(path)
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
		_, _ = FileSHA1(path)
	}
}

func BenchmarkFileSHA224(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileSHA224() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA224(path)
	}
}

func BenchmarkFileSHA256(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileSHA256() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA256(path)
	}
}

func BenchmarkFileSHA384(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileSHA384() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA384(path)
	}
}

func BenchmarkFileSHA512(b *testing.B) {
	path, found := "", false
	if path, found = filePathMap["large text file"]; !found {
		b.Errorf("FileSHA512() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA512(path)
	}
}

func TestFileHash(t *testing.T) {
	type hashTestCase struct {
		name     string
		filePath string
		wantStr  string
		wantErr  bool
	}
	tests := []struct {
		name   string
		method func(filePath string) (str string, err error)
		cases  []hashTestCase
	}{
		{
			name:   "MD5",
			method: FileMD5,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "d41d8cd98f00b204e9800998ecf8427e", false},
				{"one-line text file", "", "b10a8db164e0754105b7a99be72e3fe5", false},
				{"large text file", "", "3094ffc905b6a832d68ca27c86d52dc0", false},
				{"xlarge text file", "", "0d9d7d9349c970fbf71b46698c5d1165", false},
				{"small binary file", "", "f837aa60b6fe83458f790db60d529fc9", false},
				{"another small binary", "", "89e19cf9c9680994d75adfac30887b31", false},
			},
		},
		{
			name:   "SHA1",
			method: FileSHA1,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
				{"one-line text file", "", "0a4d55a8d778e5022fab701977c5d840bbc486d0", false},
				{"large text file", "", "22c538b3b8ba9a9f817eb05d463a0cb7ba6a9625", false},
				{"xlarge text file", "", "3912642ec8f1430ae3f1f870d2279a26a0a02297", false},
				{"small binary file", "", "14af87ccec7f81bb28d53c84da2fd5a9d5925cda", false},
				{"another small binary", "", "cd2ad24ee27178115ff6440bb3f996a142888838", false},
			},
		},
		{
			name:   "SHA224",
			method: FileSHA224,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
				{"one-line text file", "", "c4890faffdb0105d991a461e668e276685401b02eab1ef4372795047", false},
				{"large text file", "", "88810e2ea214ca516e4f68b6d5c62ee7247f9aa34d5ddd082707f73c", false},
				{"xlarge text file", "", "45f94373b06ec25a35bb3d09f92648fd6a459adc727c1c4144a45b67", false},
				{"small binary file", "", "f8bd06da0e66c71e85ffc6ca6a6ebffbf6eaf5bf97e1054148dde87c", false},
				{"another small binary", "", "f6099a9ecf7273b5d0c8ce376da0306d9ad77b8e8c958676c6df1a8a", false},
			},
		},
		{
			name:   "SHA256",
			method: FileSHA256,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
				{"one-line text file", "", "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e", false},
				{"large text file", "", "0b05758ca33d8a49752a05b695824785b0ba56689d478250eb9f6b7b9057e6f8", false},
				{"xlarge text file", "", "79ba34c7b43e2b6e4262ad966e2ba599ff53d553bd10fd73bbce096fd3ffa28f", false},
				{"small binary file", "", "dcecab1355b5c2b9ecef281322bf265ac5840b4688748586e9632b473a5fe56b", false},
				{"another small binary", "", "4d16089410a483860214a39730859e6b5a8a8b8e970911c79dd44ff331edde40", false},
			},
		},
	}
	for _, ts := range tests {
		for _, tt := range ts.cases {
			t.Run(ts.name+" "+tt.name, func(t *testing.T) {
				filePath := tt.filePath
				if len(tt.filePath) == 0 {
					var found bool
					if filePath, found = filePathMap[tt.name]; !found {
						t.Errorf("FileMD5() got no file for case '%v'", tt.name)
						return
					}
				}
				gotStr, err := ts.method(filePath)
				if (err != nil) != tt.wantErr {
					t.Errorf("File%s() error = %v, wantErr %v", ts.name, err, tt.wantErr)
					return
				}
				if gotStr != tt.wantStr {
					t.Errorf("File%s() gotStr = %v, want %v", ts.name, gotStr, tt.wantStr)
				}
			})
		}
	}
}
