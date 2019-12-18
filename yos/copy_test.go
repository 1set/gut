package yos

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

var (
	TestCaseRootCopy   string
	TestCaseOutputCopy string
	TestFileMapCopy    map[string]string
)

func init() {
	TestCaseRootCopy = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "copy")
	TestCaseOutputCopy = JoinPath(TestCaseRootCopy, "output")
	TestFileMapCopy = map[string]string{
		"EmptyDir":   JoinPath(TestCaseRootCopy, "empty-folder"),
		"ContentDir": JoinPath(TestCaseRootCopy, "content-folder"),
		"EmptyFile":  JoinPath(TestCaseRootCopy, "empty-file.txt"),
		"SmallText":  JoinPath(TestCaseRootCopy, "small-text.txt"),
		"LargeText":  JoinPath(TestCaseRootCopy, "large-text.txt"),
		"PngImage":   JoinPath(TestCaseRootCopy, "image.png"),
		"SvgImage":   JoinPath(TestCaseRootCopy, "image.svg"),
	}
}

func compareFile(file1, file2 string) (bool, error) {
	f1s, err := os.Stat(file1)
	if err != nil {
		return false, err
	}
	f2s, err := os.Stat(file2)
	if err != nil {
		return false, err
	}

	if f1s.Size() != f2s.Size() {
		return false, nil
	}

	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}

	const chunckSize = 64 * 1024
	for {
		b1 := make([]byte, chunckSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunckSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, err1
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}

func TestCopyFileV4(t *testing.T) {

	var tests = []struct {
		name     string
		srcPath  string
		destPath string
		tarPath  string
		wantErr  bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFile(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := compareFile(tt.srcPath, tt.tarPath)
				if err != nil {
					t.Errorf("CopyFile() got error while comparing the files: %v, %v, error: %v", tt.srcPath, tt.tarPath, err)
				} else if !same {
					t.Errorf("CopyFile() the files are not the same: %v, %v", tt.srcPath, tt.tarPath)
					return
				}
			}
		})
	}
}
