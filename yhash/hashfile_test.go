package yhash

import (
	"os"
	"testing"

	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
)

var (
	TestCaseRootList  string
	TestFileBenchmark string
	FilePathMap       map[string]string
)

func init() {
	TestCaseRootList = yos.JoinPath(os.Getenv("TESTRSSDIR"), "yhash")
	FilePathMap = map[string]string{
		"empty file":           yos.JoinPath(TestCaseRootList, "empty_file"),
		"one-line text file":   yos.JoinPath(TestCaseRootList, "one-line_text.txt"),
		"large text file":      yos.JoinPath(TestCaseRootList, "large_text.txt"),
		"xlarge text file":     yos.JoinPath(TestCaseRootList, "xlarge_text.txt"),
		"small binary file":    yos.JoinPath(TestCaseRootList, "small_file.bin"),
		"another small binary": yos.JoinPath(TestCaseRootList, "another_small.bin"),
		"png image":            yos.JoinPath(TestCaseRootList, "image.png"),
		"jpg image":            yos.JoinPath(TestCaseRootList, "image.jpg"),
	}
	TestFileBenchmark = FilePathMap["png image"]
}

func BenchmarkFileMD5(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileMD5() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileMD5(TestFileBenchmark)
	}
}

func BenchmarkFileSHA1(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA1() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA1(TestFileBenchmark)
	}
}

func BenchmarkFileSHA224(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA224() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA224(TestFileBenchmark)
	}
}

func BenchmarkFileSHA256(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA256() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA256(TestFileBenchmark)
	}
}

func BenchmarkFileSHA384(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA384() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA384(TestFileBenchmark)
	}
}

func BenchmarkFileSHA512(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA512() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA512(TestFileBenchmark)
	}
}

func BenchmarkFileSHA512_224(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA512_224() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA512_224(TestFileBenchmark)
	}
}

func BenchmarkFileSHA512_256(b *testing.B) {
	if ystring.IsBlank(TestFileBenchmark) {
		b.Errorf("FileSHA512_256() got no file for benchmark")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FileSHA512_256(TestFileBenchmark)
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
				{"png image", "", "ee42b0615d9d8219eb152f32297ecf20", false},
				{"jpg image", "", "ad352cef1774d58c7198a275c401b6c9", false},
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
				{"png image", "", "eebf5f2c4f6e4081c41899f959de2fd2711df9c1", false},
				{"jpg image", "", "c576e9fd3617e4c5992c5fcc147ca7ad2b719ef1", false},
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
				{"png image", "", "26c35aa68c695cf4a7086740f5f05303202531885e332b2c80d29185", false},
				{"jpg image", "", "09768296afddab047c9b0183e74cfc0dfd128e835f4e2616cc12c911", false},
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
				{"png image", "", "dfeab6a38cd16c70b1ac79201f5d1ff7fbb59f40351daae990c27f0a5ba16c22", false},
				{"jpg image", "", "05875dcc88abb73037cae1e090ca01dbfb272983f0c9055392c2f9e2e2451c39", false},
			},
		},
		{
			name:   "SHA384",
			method: FileSHA384,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", false},
				{"one-line text file", "", "99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f", false},
				{"large text file", "", "b72c19d514ca52d6ed0bea994c705523c0e2de0eca61a1f0cfd2589f06d12436c69b5b26b83aacb5217626e6c7a2fc98", false},
				{"xlarge text file", "", "929e8e1eff9b533ac203dc042a40cb54a63eda04e2b6430903daee5d1bab4206b520b0a57c31303955d98cb36e7906d1", false},
				{"small binary file", "", "2e76b983134df83f43fbacb576992f07f87d8cd0620892ba19f8dde2a94ed904abda6d1fac5c5c7dda32dd99c387eb39", false},
				{"another small binary", "", "1e0b8234559cbc8658851b6414810ee3a0b84222e3d49675a89eca50534419dccd3703410dedf13a6a0d9fde91451ed4", false},
				{"png image", "", "f8f090850bea75f43206978c2da711754d43209be2b3267254e38fc4ce93e48d01610d5e68c42d37bef9b03f814ec184", false},
				{"jpg image", "", "b6d6647539c9e23de6d2ebe0c7bf0ccf76949919168286308e734e55af4185d31bf07fe86652502188a8f4b3a98a22dd", false},
			},
		},
		{
			name:   "SHA512",
			method: FileSHA512,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
				{"one-line text file", "", "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b", false},
				{"large text file", "", "ebbf86826b75d9283f4ad559d4a9a5f61e9de13095279dfda43a5c83b78251fa5eb642dfecb732f21052789912e6efe3e1b2b89ebb725820ed9148818f536f8a", false},
				{"xlarge text file", "", "3aa8d4b6bcaeb74f4d0d7bdc0fe227865efac852a2b8edcb629230c9a8bb01eefc35de318a0d9d09f0a4bb1f8718e2fcaa511ec44e02d4f8d2354dc3edf045e0", false},
				{"small binary file", "", "a85e09c3b5dbb560f4e03ba880047dbc8b4999a64c1f54fbfbca17ee0bcbed3bc6708d699190b56668e464a59358d6b534c3963a1329ba01db21075ef5bedace", false},
				{"another small binary", "", "025eee01d2ab71d80d20c9aa461f83f6413cd9bf20d9ce9ff201d025b43f7df10609ef8d207fa31c8aa708653650bd80a3af5830f495f114e1d5d3cd909bb4d7", false},
				{"png image", "", "f96c0db10b08874254fbfa73d6576ecb41e20a95b1c287692da48adea10ac3feee6eb910bbcc293207ce777f861c1acc55b0b79fafe05bc8b4951a70cbbbb89c", false},
				{"jpg image", "", "3a3add8d5f8b4e71d712e37ce5f15e2b2ff5d44b760e5c4f7f8754077c36185a1e54984712adee26afc28b27f8a633898c87f20e7a27b583305a4b78be8c88dc", false},
			},
		},
		{
			name:   "SHA512_224",
			method: FileSHA512_224,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4", false},
				{"one-line text file", "", "feca41095c80a571ae782f96bcef9ab81bdf0182509a6844f32c4c17", false},
				{"large text file", "", "0bf0af89471ebec215018e6596d38dce46ac66f000ca1786aad03490", false},
				{"xlarge text file", "", "08374eff5a420bc96e7dc3fd596c601cae8ccbec1b380e094a392574", false},
				{"small binary file", "", "3ccd36a41c16eb62949ee655b6b02109e2ae3d6e5dbe8508f10d5a13", false},
				{"another small binary", "", "5a31ba516b588cbdb6d76ff1c760b3a11691049ef6c36d071ba9b8e5", false},
				{"png image", "", "774f87d77bac2d3461fdb5e498cf9d5930027c66e8407bb76d885638", false},
				{"jpg image", "", "a6e8c0a166ad7d9303dc1862209931cd9b71d192ec9b3e971737252b", false},
			},
		},
		{
			name:   "SHA512_256",
			method: FileSHA512_256,
			cases: []hashTestCase{
				{"file not found", "__FILE__NOT__EXIST__", "", true},
				{"empty file", "", "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a", false},
				{"one-line text file", "", "ff20018851481c25bfc2e5d0c1e1fa57dac2a237a1a96192f99a10da47aa5442", false},
				{"large text file", "", "0dbf2b5e21defb850acd8088eb025349e7aa2c8638d0eac4a221cb0103e0a82c", false},
				{"xlarge text file", "", "912567be7c2b0b9dfdc71cd086672689a67bb53ee0c7c32b4b45a331f01ba195", false},
				{"small binary file", "", "36dd5d7e2ec22d11864bcf18e27f3f2ee26ae027e123f4eed0280c8e30bd6694", false},
				{"another small binary", "", "6b47bf188df953bed6353a90872b3fae8451d96843a0879706df58598779b124", false},
				{"png image", "", "1da258ed4c48b6f853e5534f3053271c9a53bf23d6bfca02c814b3ddd763afae", false},
				{"jpg image", "", "04248ed1401a94c25ab372fce592d7c6c92c54c8b55dbd2d878ed4fae9088405", false},
			},
		},
	}
	for _, ts := range tests {
		for _, tt := range ts.cases {
			t.Run(ts.name+" "+tt.name, func(t *testing.T) {
				filePath := tt.filePath
				if ystring.IsBlank(tt.filePath) {
					var found bool
					if filePath, found = FilePathMap[tt.name]; !found {
						t.Errorf("File%s() got no file for case '%v'", ts.name, tt.name)
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
