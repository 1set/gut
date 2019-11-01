package yhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

// FileMD5 returns MD5 checksum of the named file.
// N.B. MD5 is cryptographically broken and should not be used for secure applications.
func FileMD5(filePath string) (str string, err error) {
	hash := md5.New()
	return calculateFileHash(&hash, filePath)
}

// FileSHA1 returns SHA-1 checksum of the named file.
// N.B. SHA-1 is cryptographically broken and should not be used for secure applications.
func FileSHA1(filePath string) (str string, err error) {
	hash := sha1.New()
	return calculateFileHash(&hash, filePath)
}

// FileSHA256 returns SHA256 checksum of the named file.
func FileSHA256(filePath string) (str string, err error) {
	hash := sha256.New()
	return calculateFileHash(&hash, filePath)
}
