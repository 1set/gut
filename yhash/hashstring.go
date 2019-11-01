package yhash

import (
	"crypto/md5"
	"crypto/sha1"
)

// StringMD5 returns MD5 checksum of the given string.
func StringMD5(content string) (str string, err error) {
	hash := md5.New()
	return calculateStringHash(&hash, content)
}

// StringSHA1 returns SHA1 checksum of the given string.
func StringSHA1(content string) (str string, err error) {
	hash := sha1.New()
	return calculateStringHash(&hash, content)
}
