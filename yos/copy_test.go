package yos

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

var (
	EmptyString           string
	TestCaseRootCopy      string
	TestCaseOutputCopy    string
	TestCaseBenchmarkCopy string
	TestFileMapCopy       map[string]string
	TestDirMapCopy        map[string]string
)

func init() {
	TestCaseRootCopy = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "copy")
	TestCaseOutputCopy = JoinPath(TestCaseRootCopy, "output")
	TestCaseBenchmarkCopy = JoinPath(TestCaseOutputCopy, "benchmark")
	TestFileMapCopy = map[string]string{
		"Symlink":          JoinPath(TestCaseRootCopy, "soft-link.txt"),
		"EmptyFile":        JoinPath(TestCaseRootCopy, "empty-file.txt"),
		"SmallText":        JoinPath(TestCaseRootCopy, "small-text.txt"),
		"LargeText":        JoinPath(TestCaseRootCopy, "large-text.txt"),
		"PngImage":         JoinPath(TestCaseRootCopy, "image.png"),
		"SvgImage":         JoinPath(TestCaseRootCopy, "image.svg"),
		"Out_ExistingFile": JoinPath(TestCaseOutputCopy, "existing-file.txt"),
	}
	TestDirMapCopy = map[string]string{
		"EmptyDir":        JoinPath(TestCaseRootCopy, "empty-folder"),
		"ContentDir":      JoinPath(TestCaseRootCopy, "content-folder"),
		"Out_ExistingDir": JoinPath(TestCaseOutputCopy, "existing-dir"),
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

func TestCopyFile(t *testing.T) {
	tests := []struct {
		name     string
		srcPath  string
		destPath string
		tarPath  string
		wantErr  bool
	}{
		{"Source is empty", EmptyString, TestCaseOutputCopy, EmptyString, true},
		{"Source file not exist", JoinPath(TestCaseRootCopy, "__not_exist__"), TestCaseOutputCopy, EmptyString, true},
		{"Source is a dir", TestDirMapCopy["ContentDir"], TestCaseOutputCopy, EmptyString, true},
		// {"Source is a symlink", TestFileMapCopy["Symlink"], TestCaseOutputCopy, TestCaseOutputCopy, false},
		{"Destination is empty", TestFileMapCopy["SmallText"], EmptyString, EmptyString, true},
		{"Destination is a dir", TestFileMapCopy["SmallText"], TestDirMapCopy["Out_ExistingDir"], JoinPath(TestDirMapCopy["Out_ExistingDir"], "small-text.txt"), false},
		{"Destination is a file", TestFileMapCopy["SmallText"], TestFileMapCopy["Out_ExistingFile"], TestFileMapCopy["Out_ExistingFile"], false},
		{"Destination file not exist", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "not-exist-file.txt"), JoinPath(TestCaseOutputCopy, "not-exist-file.txt"), false},
		{"Destination dir not exist", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "not-exist-dir", "not-exist-file.txt"), EmptyString, true},
		{"Copy empty file", TestFileMapCopy["EmptyFile"], JoinPath(TestCaseOutputCopy, "empty-file.txt"), JoinPath(TestCaseOutputCopy, "empty-file.txt"), false},
		{"Copy small text file", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "small-text.txt"), JoinPath(TestCaseOutputCopy, "small-text.txt"), false},
		{"Copy large text file", TestFileMapCopy["LargeText"], JoinPath(TestCaseOutputCopy, "large-text.txt"), JoinPath(TestCaseOutputCopy, "large-text.txt"), false},
		{"Copy png image file", TestFileMapCopy["PngImage"], JoinPath(TestCaseOutputCopy, "image.png"), JoinPath(TestCaseOutputCopy, "image.png"), false},
		{"Copy svg image file", TestFileMapCopy["SvgImage"], JoinPath(TestCaseOutputCopy, "image.svg"), JoinPath(TestCaseOutputCopy, "image.svg"), false},
		{"Source and destination are same", TestFileMapCopy["SmallText"], TestFileMapCopy["SmallText"], EmptyString, true},
		{"Source and destination root are same", TestFileMapCopy["SmallText"], TestCaseRootCopy, EmptyString, true},
	}

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

func BenchmarkCopyFile(b *testing.B) {
	for name, path := range TestFileMapCopy {
		if strings.HasPrefix(name, "Out_") {
			continue
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = CopyFile(path, TestCaseBenchmarkCopy)
			}
		})
	}
}
