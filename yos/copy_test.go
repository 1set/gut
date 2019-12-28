package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	emptyStr                  string
	resourceCopyRoot          string
	resourceCopyOutputRoot    string
	resourceCopyBenchmarkRoot string
	resourceCopyFileMap       map[string]string
	resourceCopyDirMap        map[string]string

	resourceCopyDirRoot          string
	resourceCopyDirOutputRoot    string
	resourceCopyDirBenchmarkRoot string
	resourceCopyDirSourceMap     map[string]string
)

func init() {
	testResourceRoot := os.Getenv("TESTRSSDIR")
	// testResourceRoot = "/var/folders/jy/cfbkpfvn6c9255yvvhfsdwzm0000gn/T/gut_test_resource"

	resourceCopyRoot = JoinPath(testResourceRoot, "yos", "copy")
	resourceCopyOutputRoot = JoinPath(resourceCopyRoot, "output")
	resourceCopyBenchmarkRoot = JoinPath(resourceCopyRoot, "benchmark")
	resourceCopyFileMap = map[string]string{
		"SymlinkFile":        JoinPath(resourceCopyRoot, "soft-link.txt"),
		"SymlinkLink":        JoinPath(resourceCopyRoot, "soft-link2.txt"),
		"SymlinkDir":         JoinPath(resourceCopyRoot, "soft-link-dir"),
		"EmptyFile":          JoinPath(resourceCopyRoot, "empty-file.txt"),
		"SmallText":          JoinPath(resourceCopyRoot, "small-text.txt"),
		"LargeText":          JoinPath(resourceCopyRoot, "large-text.txt"),
		"PngImage":           JoinPath(resourceCopyRoot, "image.png"),
		"SvgImage":           JoinPath(resourceCopyRoot, "image.svg"),
		"SameName":           JoinPath(resourceCopyRoot, "same-name"),
		"SameName2":          JoinPath(resourceCopyRoot, "same-name2"),
		"NonePermission":     JoinPath(resourceCopyRoot, "none_perm.txt"),
		"Out_NonePermission": JoinPath(resourceCopyOutputRoot, "none_perm.txt"),
		"Out_ExistingFile":   JoinPath(resourceCopyOutputRoot, "existing-file.txt"),
		"Out_SameName2":      JoinPath(resourceCopyOutputRoot, "same-name2"),
	}
	resourceCopyDirMap = map[string]string{
		"EmptyDir":        JoinPath(resourceCopyRoot, "empty-folder"),
		"ContentDir":      JoinPath(resourceCopyRoot, "content-folder"),
		"Out_ExistingDir": JoinPath(resourceCopyOutputRoot, "existing-dir"),
	}

	resourceCopyDirRoot = JoinPath(testResourceRoot, "yos", "copydir")
	resourceCopyDirOutputRoot = JoinPath(resourceCopyDirRoot, "output")
	resourceCopyDirBenchmarkRoot = JoinPath(resourceCopyDirRoot, "benchmark")
	resourceCopyDirSourceRoot := JoinPath(resourceCopyDirRoot, "source")
	resourceCopyDirSourceMap = map[string]string{
		"TextFile":        JoinPath(resourceCopyDirSourceRoot, "text.txt"),
		"Symlink":         JoinPath(resourceCopyDirSourceRoot, "link.txt"),
		"EmptyDir":        JoinPath(resourceCopyDirSourceRoot, "empty-dir"),
		"OnlyDirs":        JoinPath(resourceCopyDirSourceRoot, "only-dirs"),
		"OnlyFiles":       JoinPath(resourceCopyDirSourceRoot, "only-files"),
		"OnlySymlinks":    JoinPath(resourceCopyDirSourceRoot, "only-symlinks"),
		"NoPermDirs":      JoinPath(resourceCopyDirSourceRoot, "no-perm-dirs"),
		"NoPermFiles":     JoinPath(resourceCopyDirSourceRoot, "no-perm-files"),
		"BrokenSymlink":   JoinPath(resourceCopyDirSourceRoot, "broken-symlink"),
		"CircularSymlink": JoinPath(resourceCopyDirSourceRoot, "circular-symlink"),
		"MiscDir":         JoinPath(resourceCopyDirSourceRoot, "misc"),
		"OneFileDir":      JoinPath(resourceCopyDirSourceRoot, "one-file-dir"),
	}
}

func TestCopyFile(t *testing.T) {
	//t.Parallel()
	outputRoot := resourceCopyOutputRoot

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		inputPath  string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, outputRoot, emptyStr, emptyStr, true},
		{"Source got permission denied", resourceCopyFileMap["NonePermission"], JoinPath(outputRoot, "whatever.txt"), emptyStr, emptyStr, true},
		{"Source file not exist", JoinPath(resourceCopyRoot, "__not_exist__"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a dir", resourceCopyDirMap["ContentDir"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a symlink to file", resourceCopyFileMap["SymlinkFile"], outputRoot, resourceCopyFileMap["LargeText"], JoinPath(outputRoot, "soft-link.txt"), false},
		{"Source is a symlink to symlink", resourceCopyFileMap["SymlinkLink"], outputRoot, resourceCopyFileMap["LargeText"], JoinPath(outputRoot, "soft-link.txt"), false},
		{"Source is a symlink to dir", resourceCopyFileMap["SymlinkDir"], outputRoot, emptyStr, emptyStr, true},
		{"Destination is empty", resourceCopyFileMap["SmallText"], emptyStr, emptyStr, emptyStr, true},
		{"Destination is a dir", resourceCopyFileMap["SmallText"], resourceCopyDirMap["Out_ExistingDir"], resourceCopyFileMap["SmallText"], JoinPath(resourceCopyDirMap["Out_ExistingDir"], "small-text.txt"), false},
		{"Destination is a file", resourceCopyFileMap["SmallText"], resourceCopyFileMap["Out_ExistingFile"], resourceCopyFileMap["SmallText"], resourceCopyFileMap["Out_ExistingFile"], false},
		{"Destination got permission denied", resourceCopyFileMap["SmallText"], resourceCopyFileMap["Out_NonePermission"], emptyStr, emptyStr, true},
		{"Destination file not exist", resourceCopyFileMap["SmallText"], JoinPath(outputRoot, "not-exist-file.txt"), resourceCopyFileMap["SmallText"], JoinPath(outputRoot, "not-exist-file.txt"), false},
		{"Destination dir not exist", resourceCopyFileMap["SmallText"], JoinPath(outputRoot, "not-exist-dir", "not-exist-file.txt"), emptyStr, emptyStr, true},
		{"Copy empty file", resourceCopyFileMap["EmptyFile"], JoinPath(outputRoot, "empty-file.txt"), resourceCopyFileMap["EmptyFile"], JoinPath(outputRoot, "empty-file.txt"), false},
		{"Copy small text file", resourceCopyFileMap["SmallText"], JoinPath(outputRoot, "small-text.txt"), resourceCopyFileMap["SmallText"], JoinPath(outputRoot, "small-text.txt"), false},
		{"Copy large text file", resourceCopyFileMap["LargeText"], JoinPath(outputRoot, "large-text.txt"), resourceCopyFileMap["LargeText"], JoinPath(outputRoot, "large-text.txt"), false},
		{"Copy png image file", resourceCopyFileMap["PngImage"], JoinPath(outputRoot, "image.png"), resourceCopyFileMap["PngImage"], JoinPath(outputRoot, "image.png"), false},
		{"Copy svg image file", resourceCopyFileMap["SvgImage"], JoinPath(outputRoot, "image.svg"), resourceCopyFileMap["SvgImage"], JoinPath(outputRoot, "image.svg"), false},
		{"Source and destination are exactly the same", resourceCopyFileMap["SmallText"], resourceCopyFileMap["SmallText"], emptyStr, emptyStr, true},
		{"Source and destination are actually the same", resourceCopyFileMap["SmallText"], resourceCopyRoot, emptyStr, emptyStr, true},
		{"Source and inferred destination(dir) use the same name: can't overwrite dir", resourceCopyFileMap["SameName"], outputRoot, emptyStr, emptyStr, true},
		{"Source and inferred destination(file) use the same name: overwrite the file", resourceCopyFileMap["SameName2"], outputRoot, resourceCopyFileMap["SameName2"], resourceCopyFileMap["Out_SameName2"], false},
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
	for name, path := range resourceCopyFileMap {
		if strings.HasPrefix(name, "Out_") {
			continue
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = CopyFile(path, resourceCopyBenchmarkRoot)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	outputRoot := resourceCopyDirOutputRoot
	//expectedOutputRoot := JoinPath(resourceCopyDirRoot, "destination")

	tests := []struct {
		name         string
		srcPath      string
		destPath     string
		actualPath   string
		expectedPath string
		wantErr      bool
	}{
		{"Source is empty", emptyStr, outputRoot, emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(resourceCopyDirSourceMap["EmptyDir"], "__not_found__"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a file", resourceCopyDirSourceMap["TextFile"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a symlink", resourceCopyDirSourceMap["Symlink"], outputRoot, emptyStr, emptyStr, true},
		{"Source directory is empty", resourceCopyDirSourceMap["EmptyDir"], outputRoot, resourceCopyDirSourceMap["EmptyDir"], JoinPath(outputRoot, "empty-dir"), false},
		{"Source directory contains only directories", resourceCopyDirSourceMap["OnlyDirs"], outputRoot, resourceCopyDirSourceMap["OnlyDirs"], JoinPath(outputRoot, "only-dirs"), false},
		{"Source directory contains only files", resourceCopyDirSourceMap["OnlyFiles"], outputRoot, resourceCopyDirSourceMap["OnlyFiles"], JoinPath(outputRoot, "only-files"), false},
		{"Source directory contains only symlinks", resourceCopyDirSourceMap["OnlySymlinks"], outputRoot, resourceCopyDirSourceMap["OnlySymlinks"], JoinPath(outputRoot, "only-symlinks"), false},
		{"Source directory contains a file with no permissions", resourceCopyDirSourceMap["NoPermDirs"], outputRoot, emptyStr, emptyStr, true},
		{"Source directory contains a directory with no permissions", resourceCopyDirSourceMap["NoPermFiles"], outputRoot, emptyStr, emptyStr, true},
		{"Source directory contains a broken symlink", resourceCopyDirSourceMap["BrokenSymlink"], outputRoot, resourceCopyDirSourceMap["BrokenSymlink"], JoinPath(outputRoot, "broken-symlink"), false},
		{"Source directory contains a circular symlink", resourceCopyDirSourceMap["CircularSymlink"], outputRoot, resourceCopyDirSourceMap["CircularSymlink"], JoinPath(outputRoot, "circular-symlink"), false},
		{"Source directory contains files, symlinks and directories", resourceCopyDirSourceMap["MiscDir"], outputRoot, resourceCopyDirSourceMap["MiscDir"], JoinPath(outputRoot, "misc"), false},

		{"Destination is empty", resourceCopyDirSourceMap["OneFileDir"], emptyStr, emptyStr, emptyStr, true},
		{"Destination is a file", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "existing-file.txt"), emptyStr, emptyStr, true},
		{"Destination is a symlink", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "existing-link.txt"), emptyStr, emptyStr, true},
		{"Destination and its parent don't exist", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "non-exist", "non-exist-nested"), emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent does", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "nested-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "nested-dir"), false},
		{"Destination directory exists, and it's empty", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "empty-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "empty-dir", "one-file-dir"), false},

		//{ "Destination directory exists and already contains files", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
		//{ "Destination directory exists and already contains the same source", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
		//{ "Destination directory exists and contains a file with the same name and no permissions", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
		//{ "Destination directory exists and contains a directory with the same name and no permissions", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
		//{ "Destination directory exists and contains a symlink with the same name", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
		//{ "Destination directory exists and contains a symlink with the same name and no permissions", resourceCopyDirSourceMap["EmptyDir"], emptyStr, emptyStr, emptyStr, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "permission") && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}

			if err := CopyDir(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				ae, _ := IsDirExist(tt.actualPath)
				ee, _ := IsDirExist(tt.expectedPath)
				t.Logf("actual: %v, exist: %v", tt.actualPath, ae)
				t.Logf("expected: %v, exist: %v", tt.expectedPath, ee)
				if !(ae && ee) {
					t.Errorf("failed copy")
					return
				}
			}
		})
	}
}
