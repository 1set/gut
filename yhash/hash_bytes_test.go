package yhash

import (
	"testing"
)

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
