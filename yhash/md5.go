package yhash

import (
	"crypto/md5"
)

// FileMD5 returns MD5 checksum of the named file.
func FileMD5(filePath string) (str string, err error) {
	hash := md5.New()
	return calculateFileHash(&hash, filePath)
}

// StringMD5 returns MD5 checksum of the given string.
func StringMD5(content string) (str string, err error) {
	hash := md5.New()
	return calculateStringHash(&hash, content)
}

// BytesMD5 returns MD5 checksum of the given bytes.
func BytesMD5(data []byte) (str string, err error) {
	hash := md5.New()
	return calculateBytesHash(&hash, data)
}
