package yhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// BytesMD5 returns MD5 checksum of the given bytes.
// N.B. MD5 is cryptographically broken and should not be used for secure applications.
func BytesMD5(data []byte) (str string, err error) {
	hash := md5.New()
	return calculateBytesHash(&hash, data)
}

// BytesSHA1 returns SHA-1 checksum of the given bytes.
// N.B. SHA-1 is cryptographically broken and should not be used for secure applications.
func BytesSHA1(data []byte) (str string, err error) {
	hash := sha1.New()
	return calculateBytesHash(&hash, data)
}

// BytesSHA224 returns SHA-224 checksum of the given bytes.
func BytesSHA224(data []byte) (str string, err error) {
	hash := sha256.New224()
	return calculateBytesHash(&hash, data)
}

// BytesSHA256 returns SHA-256 checksum of the given bytes.
func BytesSHA256(data []byte) (str string, err error) {
	hash := sha256.New()
	return calculateBytesHash(&hash, data)
}

// BytesSHA384 returns SHA-384 checksum of the given bytes.
func BytesSHA384(data []byte) (str string, err error) {
	hash := sha512.New384()
	return calculateBytesHash(&hash, data)
}

// BytesSHA512 returns SHA-512 checksum of the given bytes.
func BytesSHA512(data []byte) (str string, err error) {
	hash := sha512.New()
	return calculateBytesHash(&hash, data)
}

// BytesSHA512_224 returns SHA-512/224 checksum of the given bytes.
func BytesSHA512_224(data []byte) (str string, err error) {
	hash := sha512.New512_224()
	return calculateBytesHash(&hash, data)
}

// BytesSHA512_256 returns SHA-512/256 checksum of the given bytes.
func BytesSHA512_256(data []byte) (str string, err error) {
	hash := sha512.New512_256()
	return calculateBytesHash(&hash, data)
}
