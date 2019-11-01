package yhash

import (
	"testing"
)

var bytes3k = make([]byte, 0, 3000)

func init() {
	for i := 0; i < 1000; i++ {
		bytes3k = append(bytes3k, 0x41, 0x42, 0x43)
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
		{"3k bytes", bytes3k, "0c101c2013f996d63f1f01450a4c73e9", false},
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
	for i := 0; i < b.N; i++ {
		BytesMD5(bytes3k)
	}
}

func TestBytesSHA1(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantStr string
		wantErr bool
	}{
		{"nil", nil, "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
		{"empty", []byte{}, "da39a3ee5e6b4b0d3255bfef95601890afd80709", false},
		{"one zero", []byte{0}, "5ba93c9db0cff93f52b521d7420e43f6eda2784f", false},
		{"one byte", []byte{88}, "c032adc1ff629c9b66f22749ad667e6beadf144b", false},
		{"two bytes", []byte{88, 89}, "034f1965ccdbdf9e642feeb9858da5096b6d1a9a", false},
		{"three bytes", []byte{88, 89, 90}, "717c4ecc723910edc13dd2491b0fae91442619da", false},
		{"3k bytes", bytes3k, "7cdb527beffa8ad0ffea8b0d6d6a997aa505cd27", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := BytesSHA1(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesSHA1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("BytesSHA1() gotStr = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func BenchmarkBytesSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesSHA1(bytes3k)
	}
}
