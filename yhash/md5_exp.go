package yhash

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func MD5v1(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}

	defer file.Close()
	text, err := ioutil.ReadAll(file)

	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func MD5v2(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

const bufferSize = 65536

func MD5v3(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

func MD5v4(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func MD5v5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
