package yhash

import (
	"crypto/md5"
)

// BytesMD5 returns MD5 checksum of the given bytes.
func BytesMD5(data []byte) (str string, err error) {
	hash := md5.New()
	return calculateBytesHash(&hash, data)
}
