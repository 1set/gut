package yhash

import (
	"crypto/md5"
	"crypto/sha1"
)

// BytesMD5 returns MD5 checksum of the given bytes.
func BytesMD5(data []byte) (str string, err error) {
	hash := md5.New()
	return calculateBytesHash(&hash, data)
}

// BytesSHA1 returns SHA1 checksum of the given bytes.
func BytesSHA1(data []byte) (str string, err error) {
	hash := sha1.New()
	return calculateBytesHash(&hash, data)
}
