package yhash

import (
	"crypto/md5"
)

// StringMD5 returns MD5 checksum of the given string.
func StringMD5(content string) (str string, err error) {
	hash := md5.New()
	return calculateStringHash(&hash, content)
}
