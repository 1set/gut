package yline

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type LineFunc func(line string) (err error)

var (
	QuitRead = errors.New("yline: quit read for")
)

func ReadFor(reader io.Reader, callback LineFunc) (err error) {
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		if err = callback(sc.Text()); err != nil {
			break
		}
	}

	if err == QuitRead {
		err = nil
	}
	if err == nil {
		err = sc.Err()
	}
	return
}

func ReadFileFor(filename string, callback LineFunc) (err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()
	return ReadFor(file, callback)
}

func ReadAll(reader io.Reader) (lines []string, err error) {
	err = ReadFor(reader, func(line string) (err error) {
		lines = append(lines, line)
		return
	})
	return
}

func ReadFileAll(filename string) (lines []string, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()
	return ReadAll(file)
}
