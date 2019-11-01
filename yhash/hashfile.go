package yhash

import (
	"crypto/md5"
	"crypto/sha1"
)

// FileMD5 returns MD5 checksum of the named file.
func FileMD5(filePath string) (str string, err error) {
	hash := md5.New()
	return calculateFileHash(&hash, filePath)
}

// FileSHA1 returns SHA1 checksum of the named file.
func FileSHA1(filePath string) (str string, err error) {
	hash := sha1.New()
	return calculateFileHash(&hash, filePath)
}
