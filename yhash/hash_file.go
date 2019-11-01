package yhash

import (
	"crypto/md5"
)

// FileMD5 returns MD5 checksum of the named file.
func FileMD5(filePath string) (str string, err error) {
	hash := md5.New()
	return calculateFileHash(&hash, filePath)
}
