package yhash

import (
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// Returns hash of a given file with specific algorithm as a hexadecimal number.
func calculateFileHash(algo *hash.Hash, filePath string) (str string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	if _, err = io.Copy(*algo, file); err == nil {
		str = hex.EncodeToString((*algo).Sum(nil))
	}
	return
}

// Returns hash of a given string with specific algorithm as a hexadecimal number.
func calculateStringHash(algo *hash.Hash, content string) (str string, err error) {
	if _, err = io.WriteString(*algo, content); err == nil {
		str = hex.EncodeToString((*algo).Sum(nil))
	}
	return
}

// Returns hash of a slice of bytes with specific algorithm as a hexadecimal number.
func calculateBytesHash(algo *hash.Hash, data []byte) (str string, err error) {
	if _, err = (*algo).Write(data); err == nil {
		str = hex.EncodeToString((*algo).Sum(nil))
	}
	return
}
