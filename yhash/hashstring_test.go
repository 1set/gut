package yhash

import (
	"strings"
	"testing"
)

var benchmarkContent = strings.Repeat("Angel", 1000)

func BenchmarkStringMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringMD5(benchmarkContent)
	}
}

func BenchmarkStringSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA1(benchmarkContent)
	}
}

func BenchmarkStringSHA224(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA224(benchmarkContent)
	}
}

func BenchmarkStringSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA256(benchmarkContent)
	}
}

func BenchmarkStringSHA384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA384(benchmarkContent)
	}
}

func BenchmarkStringSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA512(benchmarkContent)
	}
}

func BenchmarkStringSHA512_224(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA512_224(benchmarkContent)
	}
}

func BenchmarkStringSHA512_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = StringSHA512_256(benchmarkContent)
	}
}

func TestStringHash(t *testing.T) {
	type hashTestCase struct {
		name    string
		content string
		wantStr string
		wantErr bool
	}
	tests := []struct {
		name   string
		method func(content string) (str string, err error)
		cases  []hashTestCase
	}{
		{
			name:   "MD5",
			method: StringMD5,
			cases: []hashTestCase{
				{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e", false},
				{"one-char string", "A", "7fc56270e7a70fa81a5935b72eacbe29", false},
				{"str=12345678", "12345678", "25d55ad283aa400af464c76d713c07ad", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "142e80fd38631675e5f19dcc3e81dc11", false},
				{"long string", strings.Repeat("Good", 60), "7bbd7b0e70e71acbe1c4f0e67e59817e", false},
			},
		},
		{
			name:   "SHA1",
			method: StringSHA1,
			cases: []hashTestCase{
				{"empty string", "", "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
				{"one-char string", "B", "ae4f281df5a5d0ff3cad6371f76d5c29b6d953ec", false},
				{"str=123456789", "123456789", "f7c3bc1d808e04732adf679965ccc34ca7ae3441", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "0780f2ed8873cc7ceff6d7925eea9992f6249b0f", false},
				{"long string", strings.Repeat("Good", 60), "c8fee3477eb127f23cb7dcc9d96bd6cf96987c93", false},
			},
		},
		{
			name:   "SHA224",
			method: StringSHA224,
			cases: []hashTestCase{
				{"empty string", "", "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
				{"one-char string", "C", "484d52691fcadbfabec5a318d1cf9692c7f293cbc8c1d5f22b2d839b", false},
				{"str=123456789", "123456789", "9b3e61bf29f17c75572fae2e86e17809a4513d07c8a18152acf34521", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "8cffb7f5b819a9131f42c67dbf8ab2f79c585e0a7d53c8948eccc435", false},
				{"long string", strings.Repeat("Good", 70), "ac664b538d59ad4df4768830da3bbaf3a71eabed4aaa86c1074a8015", false},
			},
		},
		{
			name:   "SHA256",
			method: StringSHA256,
			cases: []hashTestCase{
				{"empty string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
				{"one-char string", "C", "6b23c0d5f35d1b11f9b683f0b0a617355deb11277d91ae091d399c655b87940d", false},
				{"str=123456789", "123456789", "15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "140340b962fada1ba3ed7fbc28f14f58fad155d450ca7d9c6b82f9817c9f7275", false},
				{"long string", strings.Repeat("Good", 70), "8b2428387ef532986b095492ce1afb949f797a125af101d655a5684c9fda6e8e", false},
			},
		},
		{
			name:   "SHA384",
			method: StringSHA384,
			cases: []hashTestCase{
				{"empty string", "", "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", false},
				{"one-char string", "C", "7860d388ac9e470c83d65c4b0b66bdd00e6c96fbadc78882174e020fab9793a6221724b3df9a2ec99f9395d9a410b9ed", false},
				{"str=123456789", "123456789", "eb455d56d2c1a69de64e832011f3393d45f3fa31d6842f21af92d2fe469c499da5e3179847334a18479c8d1dedea1be3", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "7727682eb706834077dbe0b5c2497f32510dbfef2d82dec305a6241c14bcabbeeef18b908c426315a5e9671f91e7108d", false},
				{"long string", strings.Repeat("Good", 70), "f7541651e7b249c7a8f99566ad119b2d303d63b1a447989f8c552ef7a1dd1e204421fb5b7bde30b61b00a31958786d60", false},
			},
		},
		{
			name:   "SHA512",
			method: StringSHA512,
			cases: []hashTestCase{
				{"empty string", "", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
				{"one-char string", "C", "3d637ae63d59522dd3cb1b81c1ad67e56d46185b0971e0bc7dd2d8ad3b26090acb634c252fc6a63b3766934314ea1a6e59fa0c8c2bc027a7b6a460b291cd4dfb", false},
				{"str=123456789", "123456789", "d9e6762dd1c8eaf6d61b3c6192fc408d4d6d5f1176d0c29169bc24e71c3f274ad27fcd5811b313d681f7e55ec02d73d499c95455b6b5bb503acf574fba8ffe85", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "e65feb8ccef3f84215b9605b5cc3b50c5e447a09e95c2c8da3eac14a2641b0e8cf8cb7995db54145ccd9655f77203a62c1de4ae6dff60bd752eb2f3f8df25464", false},
				{"long string", strings.Repeat("Good", 70), "3d396d883ec6ebfce3296910e69d11bed38b7f430a0b765015c5a88a6a50c9f015f83fb35f32dfdd591e05ad583bb62694c4c938f7c71d27018eb4b1b018213e", false},
			},
		},
		{
			name:   "SHA512_224",
			method: StringSHA512_224,
			cases: []hashTestCase{
				{"empty string", "", "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4", false},
				{"one-char string", "C", "3d3cf3c31a76ad3b7c8d99f325e75a740c8878971cefd3659016129e", false},
				{"str=123456789", "123456789", "f2a68a474bcbea375e9fc62eaab7b81fefbda64bb1c72d72e7c27314", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "bfe1cd9624cf97d9de3710e1bd9a0d52f158fccc99f4e9cfeb94eac7", false},
				{"long string", strings.Repeat("Good", 70), "e695ec53df09548c3c64933a704615a30bd756a6750cecfdc53c35f5", false},
			},
		},
		{
			name:   "SHA512_256",
			method: StringSHA512_256,
			cases: []hashTestCase{
				{"empty string", "", "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a", false},
				{"one-char string", "C", "34b99f8dde1ba273c0a28cf5b2e4dbe497f8cb2453de0c8ba6d578c9431a62cb", false},
				{"str=123456789", "123456789", "1877345237853a31ad79e14c1fcb0ddcd3df9973b61af7f906e4b4d052cc9416", false},
				{"str=你好(*´▽｀)ノノ", "你好(*´▽｀)ノノ", "69cd771a348ac00c2866c444a8daa32ceab35683a8171dadc3858baac14adabc", false},
				{"long string", strings.Repeat("Good", 70), "47655077f958003f7fd26b2ed1e95d4d889f7de4ae7915aba6578b64fe4f0837", false},
			},
		},
	}
	for _, ts := range tests {
		for _, tt := range ts.cases {
			t.Run(ts.name+" "+tt.name, func(t *testing.T) {
				gotStr, err := ts.method(tt.content)
				if (err != nil) != tt.wantErr {
					t.Errorf("String%s() error = %v, wantErr %v", ts.name, err, tt.wantErr)
					return
				}
				if gotStr != tt.wantStr {
					t.Errorf("String%s() gotStr = %v, want %v", ts.name, gotStr, tt.wantStr)
				}
			})
		}
	}
}
