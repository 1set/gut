package yhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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

// StringSHA224 returns SHA-224 checksum of the given string.
func StringSHA224(content string) (str string, err error) {
	hash := sha256.New224()
	return calculateStringHash(&hash, content)
}

// StringSHA256 returns SHA-256 checksum of the given string.
func StringSHA256(content string) (str string, err error) {
	hash := sha256.New()
	return calculateStringHash(&hash, content)
}

// StringSHA384 returns SHA-384 checksum of the given string.
func StringSHA384(content string) (str string, err error) {
	hash := sha512.New384()
	return calculateStringHash(&hash, content)
}

// StringSHA512 returns SHA-512 checksum of the given string.
func StringSHA512(content string) (str string, err error) {
	hash := sha512.New()
	return calculateStringHash(&hash, content)
}

// StringSHA512_224 returns SHA-512/224 checksum of the given string.
func StringSHA512_224(content string) (str string, err error) {
	hash := sha512.New512_224()
	return calculateStringHash(&hash, content)
}

// StringSHA512_256 returns SHA-512/256 checksum of the given string.
func StringSHA512_256(content string) (str string, err error) {
	hash := sha512.New512_256()
	return calculateStringHash(&hash, content)
}
