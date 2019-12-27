package yos

import (
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
		"SymlinkFile":        JoinPath(TestCaseRootCopy, "soft-link.txt"),
		"SymlinkLink":        JoinPath(TestCaseRootCopy, "soft-link2.txt"),
		"SymlinkDir":         JoinPath(TestCaseRootCopy, "soft-link-dir"),
		"EmptyFile":          JoinPath(TestCaseRootCopy, "empty-file.txt"),
		"SmallText":          JoinPath(TestCaseRootCopy, "small-text.txt"),
		"LargeText":          JoinPath(TestCaseRootCopy, "large-text.txt"),
		"PngImage":           JoinPath(TestCaseRootCopy, "image.png"),
		"SvgImage":           JoinPath(TestCaseRootCopy, "image.svg"),
		"SameName":           JoinPath(TestCaseRootCopy, "same-name"),
		"SameName2":          JoinPath(TestCaseRootCopy, "same-name2"),
		"NonePermission":     JoinPath(TestCaseRootCopy, "none_perm.txt"),
		"Out_NonePermission": JoinPath(TestCaseOutputCopy, "none_perm.txt"),
		"Out_ExistingFile":   JoinPath(TestCaseOutputCopy, "existing-file.txt"),
		"Out_SameName2":      JoinPath(TestCaseOutputCopy, "same-name2"),
	}
	TestDirMapCopy = map[string]string{
		"EmptyDir":        JoinPath(TestCaseRootCopy, "empty-folder"),
		"ContentDir":      JoinPath(TestCaseRootCopy, "content-folder"),
		"Out_ExistingDir": JoinPath(TestCaseOutputCopy, "existing-dir"),
	}
}

func TestCopyFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		inputPath  string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", EmptyString, TestCaseOutputCopy, EmptyString, EmptyString, true},
		{"Source got permission denied", TestFileMapCopy["NonePermission"], JoinPath(TestCaseOutputCopy, "whatever.txt"), EmptyString, EmptyString, true},
		{"Source file not exist", JoinPath(TestCaseRootCopy, "__not_exist__"), TestCaseOutputCopy, EmptyString, EmptyString, true},
		{"Source is a dir", TestDirMapCopy["ContentDir"], TestCaseOutputCopy, EmptyString, EmptyString, true},
		{"Source is a symlink to file", TestFileMapCopy["SymlinkFile"], TestCaseOutputCopy, TestFileMapCopy["LargeText"], JoinPath(TestCaseOutputCopy, "soft-link.txt"), false},
		{"Source is a symlink to symlink", TestFileMapCopy["SymlinkLink"], TestCaseOutputCopy, TestFileMapCopy["LargeText"], JoinPath(TestCaseOutputCopy, "soft-link.txt"), false},
		{"Source is a symlink to dir", TestFileMapCopy["SymlinkDir"], TestCaseOutputCopy, EmptyString, EmptyString, true},
		{"Destination is empty", TestFileMapCopy["SmallText"], EmptyString, EmptyString, EmptyString, true},
		{"Destination is a dir", TestFileMapCopy["SmallText"], TestDirMapCopy["Out_ExistingDir"], TestFileMapCopy["SmallText"], JoinPath(TestDirMapCopy["Out_ExistingDir"], "small-text.txt"), false},
		{"Destination is a file", TestFileMapCopy["SmallText"], TestFileMapCopy["Out_ExistingFile"], TestFileMapCopy["SmallText"], TestFileMapCopy["Out_ExistingFile"], false},
		{"Destination got permission denied", TestFileMapCopy["SmallText"], TestFileMapCopy["Out_NonePermission"], EmptyString, EmptyString, true},
		{"Destination file not exist", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "not-exist-file.txt"), TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "not-exist-file.txt"), false},
		{"Destination dir not exist", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "not-exist-dir", "not-exist-file.txt"), EmptyString, EmptyString, true},
		{"Copy empty file", TestFileMapCopy["EmptyFile"], JoinPath(TestCaseOutputCopy, "empty-file.txt"), TestFileMapCopy["EmptyFile"], JoinPath(TestCaseOutputCopy, "empty-file.txt"), false},
		{"Copy small text file", TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "small-text.txt"), TestFileMapCopy["SmallText"], JoinPath(TestCaseOutputCopy, "small-text.txt"), false},
		{"Copy large text file", TestFileMapCopy["LargeText"], JoinPath(TestCaseOutputCopy, "large-text.txt"), TestFileMapCopy["LargeText"], JoinPath(TestCaseOutputCopy, "large-text.txt"), false},
		{"Copy png image file", TestFileMapCopy["PngImage"], JoinPath(TestCaseOutputCopy, "image.png"), TestFileMapCopy["PngImage"], JoinPath(TestCaseOutputCopy, "image.png"), false},
		{"Copy svg image file", TestFileMapCopy["SvgImage"], JoinPath(TestCaseOutputCopy, "image.svg"), TestFileMapCopy["SvgImage"], JoinPath(TestCaseOutputCopy, "image.svg"), false},
		{"Source and destination are exactly the same", TestFileMapCopy["SmallText"], TestFileMapCopy["SmallText"], EmptyString, EmptyString, true},
		{"Source and destination are actually the same", TestFileMapCopy["SmallText"], TestCaseRootCopy, EmptyString, EmptyString, true},
		{"Source and inferred destination(dir) use the same name: can't overwrite dir", TestFileMapCopy["SameName"], TestCaseOutputCopy, EmptyString, EmptyString, true},
		{"Source and inferred destination(file) use the same name: overwrite the file", TestFileMapCopy["SameName2"], TestCaseOutputCopy, TestFileMapCopy["SameName2"], TestFileMapCopy["Out_SameName2"], false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "permission") && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}

			if err := CopyFile(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameContent(tt.inputPath, tt.outputPath)
				if err != nil {
					t.Errorf("CopyFile() got error while comparing the files: %v, %v, error: %v", tt.inputPath, tt.outputPath, err)
				} else if !same {
					t.Errorf("CopyFile() the files are not the same: %v, %v", tt.inputPath, tt.outputPath)
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
