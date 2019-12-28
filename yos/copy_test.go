package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	EmptyString           string
	CaseCopyRoot          string
	CaseCopyOutputRoot    string
	CaseCopyBenchmarkRoot string
	CaseCopyFileMap       map[string]string
	CaseCopyDirMap        map[string]string

	CaseCopyDirRoot           string
	CaseCopyDirOutputRoot     string
	CaseCopyDirBenchmarkRoot  string
	CaseCopyDirSourceRoot     string
	CaseCopyDirSourceMap      map[string]string
	CaseCopyDirOutputMap      map[string]string
	CaseCopyDirDestinationMap map[string]string
)

func init() {
	testResourceRoot := os.Getenv("TESTRSSDIR")
	// testResourceRoot = "/var/folders/jy/cfbkpfvn6c9255yvvhfsdwzm0000gn/T/gut_test_resource"

	CaseCopyRoot = JoinPath(testResourceRoot, "yos", "copy")
	CaseCopyOutputRoot = JoinPath(CaseCopyRoot, "output")
	CaseCopyBenchmarkRoot = JoinPath(CaseCopyRoot, "benchmark")

	CaseCopyDirRoot = JoinPath(testResourceRoot, "yos", "copydir")
	CaseCopyDirSourceRoot = JoinPath(CaseCopyDirRoot, "source")
	CaseCopyDirOutputRoot = JoinPath(CaseCopyDirRoot, "output")
	CaseCopyDirBenchmarkRoot = JoinPath(CaseCopyDirRoot, "benchmark")

	CaseCopyFileMap = map[string]string{
		"SymlinkFile":        JoinPath(CaseCopyRoot, "soft-link.txt"),
		"SymlinkLink":        JoinPath(CaseCopyRoot, "soft-link2.txt"),
		"SymlinkDir":         JoinPath(CaseCopyRoot, "soft-link-dir"),
		"EmptyFile":          JoinPath(CaseCopyRoot, "empty-file.txt"),
		"SmallText":          JoinPath(CaseCopyRoot, "small-text.txt"),
		"LargeText":          JoinPath(CaseCopyRoot, "large-text.txt"),
		"PngImage":           JoinPath(CaseCopyRoot, "image.png"),
		"SvgImage":           JoinPath(CaseCopyRoot, "image.svg"),
		"SameName":           JoinPath(CaseCopyRoot, "same-name"),
		"SameName2":          JoinPath(CaseCopyRoot, "same-name2"),
		"NonePermission":     JoinPath(CaseCopyRoot, "none_perm.txt"),
		"Out_NonePermission": JoinPath(CaseCopyOutputRoot, "none_perm.txt"),
		"Out_ExistingFile":   JoinPath(CaseCopyOutputRoot, "existing-file.txt"),
		"Out_SameName2":      JoinPath(CaseCopyOutputRoot, "same-name2"),
	}
	CaseCopyDirMap = map[string]string{
		"EmptyDir":        JoinPath(CaseCopyRoot, "empty-folder"),
		"ContentDir":      JoinPath(CaseCopyRoot, "content-folder"),
		"Out_ExistingDir": JoinPath(CaseCopyOutputRoot, "existing-dir"),
	}

	CaseCopyDirSourceMap = map[string]string{
		"TextFile":        JoinPath(CaseCopyDirSourceRoot, "text.txt"),
		"Symlink":         JoinPath(CaseCopyDirSourceRoot, "link.txt"),
		"EmptyDir":        JoinPath(CaseCopyDirSourceRoot, "empty-dir"),
		"OnlyDirs":        JoinPath(CaseCopyDirSourceRoot, "only-dirs"),
		"OnlyFiles":       JoinPath(CaseCopyDirSourceRoot, "only-files"),
		"OnlySymlinks":    JoinPath(CaseCopyDirSourceRoot, "only-symlinks"),
		"NoPermDirs":      JoinPath(CaseCopyDirSourceRoot, "no-perm-dirs"),
		"NoPermFiles":     JoinPath(CaseCopyDirSourceRoot, "no-perm-files"),
		"BrokenSymlink":   JoinPath(CaseCopyDirSourceRoot, "broken-symlink"),
		"CircularSymlink": JoinPath(CaseCopyDirSourceRoot, "circular-symlink"),
		"MiscDir":         JoinPath(CaseCopyDirSourceRoot, "misc"),
	}
	CaseCopyDirOutputMap = map[string]string{
		"Nothing":         JoinPath(CaseCopyDirOutputRoot, "nothing"),
		"EmptyDir":        JoinPath(CaseCopyDirOutputRoot, "EmptyDir"),
		"OnlyDirs":        JoinPath(CaseCopyDirOutputRoot, "OnlyDirs"),
		"OnlyFiles":       JoinPath(CaseCopyDirOutputRoot, "OnlyFiles"),
		"OnlySymlinks":    JoinPath(CaseCopyDirOutputRoot, "OnlySymlinks"),
		"NoPermDirs":      JoinPath(CaseCopyDirOutputRoot, "NoPermDirs"),
		"NoPermFiles":     JoinPath(CaseCopyDirOutputRoot, "NoPermFiles"),
		"BrokenSymlink":   JoinPath(CaseCopyDirOutputRoot, "BrokenSymlink"),
		"CircularSymlink": JoinPath(CaseCopyDirOutputRoot, "CircularSymlink"),
		"MiscDir":         JoinPath(CaseCopyDirOutputRoot, "MiscDir"),
	}
}

func TestCopyFile(t *testing.T) {
	//t.Parallel()
	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		inputPath  string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", EmptyString, CaseCopyOutputRoot, EmptyString, EmptyString, true},
		{"Source got permission denied", CaseCopyFileMap["NonePermission"], JoinPath(CaseCopyOutputRoot, "whatever.txt"), EmptyString, EmptyString, true},
		{"Source file not exist", JoinPath(CaseCopyRoot, "__not_exist__"), CaseCopyOutputRoot, EmptyString, EmptyString, true},
		{"Source is a dir", CaseCopyDirMap["ContentDir"], CaseCopyOutputRoot, EmptyString, EmptyString, true},
		{"Source is a symlink to file", CaseCopyFileMap["SymlinkFile"], CaseCopyOutputRoot, CaseCopyFileMap["LargeText"], JoinPath(CaseCopyOutputRoot, "soft-link.txt"), false},
		{"Source is a symlink to symlink", CaseCopyFileMap["SymlinkLink"], CaseCopyOutputRoot, CaseCopyFileMap["LargeText"], JoinPath(CaseCopyOutputRoot, "soft-link.txt"), false},
		{"Source is a symlink to dir", CaseCopyFileMap["SymlinkDir"], CaseCopyOutputRoot, EmptyString, EmptyString, true},
		{"Destination is empty", CaseCopyFileMap["SmallText"], EmptyString, EmptyString, EmptyString, true},
		{"Destination is a dir", CaseCopyFileMap["SmallText"], CaseCopyDirMap["Out_ExistingDir"], CaseCopyFileMap["SmallText"], JoinPath(CaseCopyDirMap["Out_ExistingDir"], "small-text.txt"), false},
		{"Destination is a file", CaseCopyFileMap["SmallText"], CaseCopyFileMap["Out_ExistingFile"], CaseCopyFileMap["SmallText"], CaseCopyFileMap["Out_ExistingFile"], false},
		{"Destination got permission denied", CaseCopyFileMap["SmallText"], CaseCopyFileMap["Out_NonePermission"], EmptyString, EmptyString, true},
		{"Destination file not exist", CaseCopyFileMap["SmallText"], JoinPath(CaseCopyOutputRoot, "not-exist-file.txt"), CaseCopyFileMap["SmallText"], JoinPath(CaseCopyOutputRoot, "not-exist-file.txt"), false},
		{"Destination dir not exist", CaseCopyFileMap["SmallText"], JoinPath(CaseCopyOutputRoot, "not-exist-dir", "not-exist-file.txt"), EmptyString, EmptyString, true},
		{"Copy empty file", CaseCopyFileMap["EmptyFile"], JoinPath(CaseCopyOutputRoot, "empty-file.txt"), CaseCopyFileMap["EmptyFile"], JoinPath(CaseCopyOutputRoot, "empty-file.txt"), false},
		{"Copy small text file", CaseCopyFileMap["SmallText"], JoinPath(CaseCopyOutputRoot, "small-text.txt"), CaseCopyFileMap["SmallText"], JoinPath(CaseCopyOutputRoot, "small-text.txt"), false},
		{"Copy large text file", CaseCopyFileMap["LargeText"], JoinPath(CaseCopyOutputRoot, "large-text.txt"), CaseCopyFileMap["LargeText"], JoinPath(CaseCopyOutputRoot, "large-text.txt"), false},
		{"Copy png image file", CaseCopyFileMap["PngImage"], JoinPath(CaseCopyOutputRoot, "image.png"), CaseCopyFileMap["PngImage"], JoinPath(CaseCopyOutputRoot, "image.png"), false},
		{"Copy svg image file", CaseCopyFileMap["SvgImage"], JoinPath(CaseCopyOutputRoot, "image.svg"), CaseCopyFileMap["SvgImage"], JoinPath(CaseCopyOutputRoot, "image.svg"), false},
		{"Source and destination are exactly the same", CaseCopyFileMap["SmallText"], CaseCopyFileMap["SmallText"], EmptyString, EmptyString, true},
		{"Source and destination are actually the same", CaseCopyFileMap["SmallText"], CaseCopyRoot, EmptyString, EmptyString, true},
		{"Source and inferred destination(dir) use the same name: can't overwrite dir", CaseCopyFileMap["SameName"], CaseCopyOutputRoot, EmptyString, EmptyString, true},
		{"Source and inferred destination(file) use the same name: overwrite the file", CaseCopyFileMap["SameName2"], CaseCopyOutputRoot, CaseCopyFileMap["SameName2"], CaseCopyFileMap["Out_SameName2"], false},
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
	for name, path := range CaseCopyFileMap {
		if strings.HasPrefix(name, "Out_") {
			continue
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = CopyFile(path, CaseCopyBenchmarkRoot)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	tests := []struct {
		name         string
		srcPath      string
		destPath     string
		actualPath   string
		expectedPath string
		wantErr      bool
	}{
		{"Source is empty", EmptyString, CaseCopyDirOutputMap["Nothing"], EmptyString, EmptyString, true},
		{"Source doesn't exist", JoinPath(CaseCopyDirSourceRoot, "__not_found__"), CaseCopyDirOutputMap["Nothing"], EmptyString, EmptyString, true},
		{"Source is a file", CaseCopyDirSourceMap["TextFile"], CaseCopyDirOutputMap["Nothing"], EmptyString, EmptyString, true},
		{"Source is a symlink", CaseCopyDirSourceMap["Symlink"], CaseCopyDirOutputMap["Nothing"], EmptyString, EmptyString, true},
		{"Source directory is empty", CaseCopyDirSourceMap["EmptyDir"], CaseCopyDirOutputMap["EmptyDir"], CaseCopyDirSourceMap["EmptyDir"], JoinPath(CaseCopyDirOutputMap["EmptyDir"], "empty-dir"), false},
		{"Source directory contains only directories", CaseCopyDirSourceMap["OnlyDirs"], CaseCopyDirOutputMap["OnlyDirs"], CaseCopyDirSourceMap["OnlyDirs"], JoinPath(CaseCopyDirOutputMap["OnlyDirs"], "only-dirs"), false},
		{"Source directory contains only files", CaseCopyDirSourceMap["OnlyFiles"], CaseCopyDirOutputMap["OnlyFiles"], CaseCopyDirSourceMap["OnlyFiles"], JoinPath(CaseCopyDirOutputMap["OnlyFiles"], "only-files"), false},
		{"Source directory contains only symlinks", CaseCopyDirSourceMap["OnlySymlinks"], CaseCopyDirOutputMap["OnlySymlinks"], CaseCopyDirSourceMap["OnlySymlinks"], JoinPath(CaseCopyDirOutputMap["OnlySymlinks"], "only-symlinks"), false},
		{"Source directory contains a file with no permissions", CaseCopyDirSourceMap["NoPermDirs"], CaseCopyDirOutputMap["NoPermDirs"], EmptyString, EmptyString, true},
		{"Source directory contains a directory with no permissions", CaseCopyDirSourceMap["NoPermFiles"], CaseCopyDirOutputMap["NoPermFiles"], EmptyString, EmptyString, true},
		{"Source directory contains a broken symlink", CaseCopyDirSourceMap["BrokenSymlink"], CaseCopyDirOutputMap["BrokenSymlink"], CaseCopyDirSourceMap["BrokenSymlink"], JoinPath(CaseCopyDirOutputMap["BrokenSymlink"], "broken-symlink"), false},
		{"Source directory contains a circular symlink", CaseCopyDirSourceMap["CircularSymlink"], CaseCopyDirOutputMap["CircularSymlink"], CaseCopyDirSourceMap["CircularSymlink"], JoinPath(CaseCopyDirOutputMap["CircularSymlink"], "circular-symlink"), false},
		{"Source directory contains files, symlinks and directories", CaseCopyDirSourceMap["MiscDir"], CaseCopyDirOutputMap["MiscDir"], CaseCopyDirSourceMap["MiscDir"], JoinPath(CaseCopyDirOutputMap["MiscDir"], "misc"), false},
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
