package yhash

import (
	"testing"
)

var bytes4k = make([]byte, 0, 4096)

func init() {
	for i := 0; i < 1024; i++ {
		bytes4k = append(bytes4k, 0x46, 0x55, 0x43, 0x4b)
	}
}

func BenchmarkBytesMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesMD5(bytes4k)
	}
}

func BenchmarkBytesSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesSHA1(bytes4k)
	}
}

func BenchmarkBytesSHA224(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesSHA224(bytes4k)
	}
}

func BenchmarkBytesSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesSHA256(bytes4k)
	}
}

func BenchmarkBytesSHA384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesSHA384(bytes4k)
	}
}

func BenchmarkBytesSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = BytesSHA512(bytes4k)
	}
}

func TestStringBytes(t *testing.T) {
	type hashTestCase struct {
		name    string
		data    []byte
		wantStr string
		wantErr bool
	}
	tests := []struct {
		name   string
		method func(data []byte) (str string, err error)
		cases  []hashTestCase
	}{
		{
			name:   "MD5",
			method: BytesMD5,
			cases: []hashTestCase{
				{"nil", nil, "d41d8cd98f00b204e9800998ecf8427e", false},
				{"empty", []byte{}, "d41d8cd98f00b204e9800998ecf8427e", false},
				{"one zero", []byte{0}, "93b885adfe0da089cdf634904fd59f71", false},
				{"one byte", []byte{88}, "02129bb861061d1a052c592e2dc6b383", false},
				{"two bytes", []byte{88, 89}, "74c53bcd3dcb2bb79993b2fec37d362a", false},
				{"three bytes", []byte{88, 89, 90}, "e65075d550f9b5bf9992fa1d71a131be", false},
				{"4k bytes", bytes4k, "f57c8ef3cb002cb6069be7c805f83ae4", false},
			},
		},
		{
			name:   "SHA1",
			method: BytesSHA1,
			cases: []hashTestCase{
				{"nil", nil, "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
				{"empty", []byte{}, "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
				{"one zero", []byte{0}, "5ba93c9db0cff93f52b521d7420e43f6eda2784f", false},
				{"one byte", []byte{88}, "c032adc1ff629c9b66f22749ad667e6beadf144b", false},
				{"two bytes", []byte{88, 89}, "034f1965ccdbdf9e642feeb9858da5096b6d1a9a", false},
				{"three bytes", []byte{88, 89, 90}, "717c4ecc723910edc13dd2491b0fae91442619da", false},
				{"4k bytes", bytes4k, "695634e0d1baf3b99e3cae648414a7829d369f0d", false},
			},
		},
		{
			name:   "SHA224",
			method: BytesSHA224,
			cases: []hashTestCase{
				{"nil", nil, "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
				{"empty", []byte{}, "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
				{"one zero", []byte{0}, "fff9292b4201617bdc4d3053fce02734166a683d7d858a7f5f59b073", false},
				{"one byte", []byte{88}, "f00bdeb2cd9da240a57c951fdf1bcba509fd0cd83c5e5ad9a669de12", false},
				{"two bytes", []byte{88, 89}, "a3a149bd66cd66e971d8ca4c12394818f6c63bca01a0d8c6b730f0d7", false},
				{"three bytes", []byte{88, 89, 90}, "cc8660476871488742e0cac93a996a1b4fab7d3b7e3df10412cc0059", false},
				{"4k bytes", bytes4k, "082947382dc751e6f5bbc59224b2758b7ce99627210715ba0ab6bced", false},
			},
		},
		{
			name:   "SHA256",
			method: BytesSHA256,
			cases: []hashTestCase{
				{"nil", nil, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
				{"empty", []byte{}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
				{"one zero", []byte{0}, "6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d", false},
				{"one byte", []byte{88}, "4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015", false},
				{"two bytes", []byte{88, 89}, "c07a3de039fbc0914689549f041eae295d621de7f7f647fd863f6d2f8db2080e", false},
				{"three bytes", []byte{88, 89, 90}, "ade099751d2ea9f3393f0f32d20c6b980dd5d3b0989dea599b966ae0d3cd5a1e", false},
				{"4k bytes", bytes4k, "ba182851504af83589df0acd6ba850754d02cf61bff1ecd97ad810c34cfcdf79", false},
			},
		},
		{
			name:   "SHA384",
			method: BytesSHA384,
			cases: []hashTestCase{
				{"nil", nil, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", false},
				{"empty", []byte{}, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", false},
				{"one zero", []byte{0}, "bec021b4f368e3069134e012c2b4307083d3a9bdd206e24e5f0d86e13d6636655933ec2b413465966817a9c208a11717", false},
				{"one byte", []byte{88}, "754fe9beaa91bb7ae98bee55168e16c7b1f3c5aa54ccf83c28db3384633cace48639beee8cd005e3ebb6b95dd43c95b7", false},
				{"two bytes", []byte{88, 89}, "e82e4ac84aee08256eba503c33d3e3ed2b147c62074e2f3e6dd6a66709785463fbc2f49ec2f31d97fc9f1d2a65070e4c", false},
				{"three bytes", []byte{88, 89, 90}, "165f03f9bc00245fff1fa8febef2bc786eca3e11773b88f705d88ba3ccc26b63afb535029013bf682602ffc0eaaab482", false},
				{"4k bytes", bytes4k, "4f0f52037db97d7c3cdc3f2c58d479ea212f2e2456a9a64335922e6942ad4237bb79c18d4a6fe212810ad3019c6ef9ec", false},
			},
		},
		{
			name:   "SHA512",
			method: BytesSHA512,
			cases: []hashTestCase{
				{"nil", nil, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
				{"empty", []byte{}, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
				{"one zero", []byte{0}, "b8244d028981d693af7b456af8efa4cad63d282e19ff14942c246e50d9351d22704a802a71c3580b6370de4ceb293c324a8423342557d4e5c38438f0e36910ee", false},
				{"one byte", []byte{88}, "3173f0564ab9462b0978a765c1283f96f05ac9e9f8361ee1006dc905c153d85bf0e4c45622e5e990abcf48fb5192ad34722e8d6a723278b39fef9e4f9fc62378", false},
				{"two bytes", []byte{88, 89}, "be50868b08abc38408269a76330c39e12119dd7e1581ff0910addc8de4e8f7f47874b494e33b8972448bd6c2cd329278b5dc439c555b12d3fad29a1f2dc28571", false},
				{"three bytes", []byte{88, 89, 90}, "5f6bb0f20b1d682c099d6b7a487c83b2b399fa8587d3ce801909f91539589683a447fb382be91691362c7625cd909a0e07bf0389e2081f2c3af56ecdf93ea987", false},
				{"4k bytes", bytes4k, "846b1834529bb78aa3ec393c82d38e0bc4be39490f16a44f9d4a01a11e6bb064551dda5f1dc315af788e388bbf7da5d44ff675ec9049c069f6398dd20636ec69", false},
			},
		},
	}
	for _, ts := range tests {
		for _, tt := range ts.cases {
			t.Run(ts.name+" "+tt.name, func(t *testing.T) {
				gotStr, err := ts.method(tt.data)
				if (err != nil) != tt.wantErr {
					t.Errorf("Bytes%s() error = %v, wantErr %v", ts.name, err, tt.wantErr)
					return
				}
				if gotStr != tt.wantStr {
					t.Errorf("Bytes%s() gotStr = %v, want %v", ts.name, gotStr, tt.wantStr)
				}
			})
		}
	}
}
