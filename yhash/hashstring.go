package yhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

// StringMD5 returns MD5 checksum of the given string.
// N.B. MD5 is cryptographically broken and should not be used for secure applications.
func StringMD5(content string) (str string, err error) {
	hash := md5.New()
	return calculateStringHash(&hash, content)
}

// StringSHA1 returns SHA-1 checksum of the given string.
// N.B. SHA-1 is cryptographically broken and should not be used for secure applications.
func StringSHA1(content string) (str string, err error) {
	hash := sha1.New()
	return calculateStringHash(&hash, content)
}

// StringSHA224 returns SHA224 checksum of the given string.
func StringSHA224(content string) (str string, err error) {
	hash := sha256.New224()
	return calculateStringHash(&hash, content)
}

// StringSHA256 returns SHA256 checksum of the given string.
func StringSHA256(content string) (str string, err error) {
	hash := sha256.New()
	return calculateStringHash(&hash, content)
}
