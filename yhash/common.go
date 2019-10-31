package yhash

import (
	"encoding/hex"
	"hash"
	"io"
	"os"
)

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

func calculateStringHash(algo *hash.Hash, content string) (str string, err error) {
	if _, err = io.WriteString(*algo, content); err == nil {
		str = hex.EncodeToString((*algo).Sum(nil))
	}
	return
}

func calculateBytesHash(algo *hash.Hash, data []byte) (str string, err error) {
	if _, err = (*algo).Write(data); err == nil {
		str = hex.EncodeToString((*algo).Sum(nil))
	}
	return
}
