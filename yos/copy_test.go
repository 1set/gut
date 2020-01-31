package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	emptyStr                      string
	resourceCopyFileRoot          string
	resourceCopyFileOutputRoot    string
	resourceCopyFileBenchmarkRoot string
	resourceCopyFileFileMap       map[string]string
	resourceCopyFileDirMap        map[string]string

	resourceCopyDirRoot          string
	resourceCopyDirOutputRoot    string
	resourceCopyDirBenchmarkRoot string
	resourceCopyDirSourceRoot    string
	resourceCopyDirSourceMap     map[string]string

	resourceCopySymlinkRoot          string
	resourceCopySymlinkSourceRoot    string
	resourceCopySymlinkOutputRoot    string
	resourceCopySymlinkBenchmarkRoot string
)

func init() {
	testResourceRoot := os.Getenv("TESTRSSDIR")

	resourceCopyFileRoot = JoinPath(testResourceRoot, "yos", "copy_file")
	resourceCopyFileOutputRoot = JoinPath(resourceCopyFileRoot, "output")
	resourceCopyFileBenchmarkRoot = JoinPath(resourceCopyFileRoot, "benchmark")
	resourceCopyFileFileMap = map[string]string{
		"SymlinkFile":        JoinPath(resourceCopyFileRoot, "soft-link.txt"),
		"SymlinkLink":        JoinPath(resourceCopyFileRoot, "soft-link2.txt"),
		"SymlinkDir":         JoinPath(resourceCopyFileRoot, "soft-link-dir"),
		"SymlinkBroken":      JoinPath(resourceCopyFileRoot, "soft-link-broken"),
		"SymlinkCircular":    JoinPath(resourceCopyFileRoot, "soft-link-circular"),
		"EmptyFile":          JoinPath(resourceCopyFileRoot, "empty-file.txt"),
		"SmallText":          JoinPath(resourceCopyFileRoot, "small-text.txt"),
		"LargeText":          JoinPath(resourceCopyFileRoot, "large-text.txt"),
		"PngImage":           JoinPath(resourceCopyFileRoot, "image.png"),
		"SvgImage":           JoinPath(resourceCopyFileRoot, "image.svg"),
		"SameName":           JoinPath(resourceCopyFileRoot, "same-name"),
		"SameName2":          JoinPath(resourceCopyFileRoot, "same-name2"),
		"NonePermission":     JoinPath(resourceCopyFileRoot, "none_perm.txt"),
		"Out_NonePermission": JoinPath(resourceCopyFileOutputRoot, "none_perm.txt"),
		"Out_ExistingFile":   JoinPath(resourceCopyFileOutputRoot, "existing-file.txt"),
		"Out_SameName2":      JoinPath(resourceCopyFileOutputRoot, "same-name2"),
	}
	resourceCopyFileDirMap = map[string]string{
		"EmptyDir":        JoinPath(resourceCopyFileRoot, "empty-folder"),
		"ContentDir":      JoinPath(resourceCopyFileRoot, "content-folder"),
		"Out_ExistingDir": JoinPath(resourceCopyFileOutputRoot, "existing-dir"),
	}

	resourceCopyDirRoot = JoinPath(testResourceRoot, "yos", "copy_dir")
	resourceCopyDirOutputRoot = JoinPath(resourceCopyDirRoot, "output")
	resourceCopyDirBenchmarkRoot = JoinPath(resourceCopyDirRoot, "benchmark")
	resourceCopyDirSourceRoot = JoinPath(resourceCopyDirRoot, "source")
	resourceCopyDirSourceMap = map[string]string{
		"TextFile":        JoinPath(resourceCopyDirSourceRoot, "text.txt"),
		"SymlinkDir":      JoinPath(resourceCopyDirSourceRoot, "link-dir"),
		"SymlinkFile":     JoinPath(resourceCopyDirSourceRoot, "link.txt"),
		"SymlinkBroken":   JoinPath(resourceCopyDirSourceRoot, "link-broken"),
		"SymlinkCircular": JoinPath(resourceCopyDirSourceRoot, "link-circular"),
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

	resourceCopySymlinkRoot = JoinPath(testResourceRoot, "yos", "copy_link")
	resourceCopySymlinkOutputRoot = JoinPath(resourceCopySymlinkRoot, "output")
	resourceCopySymlinkSourceRoot = JoinPath(resourceCopySymlinkRoot, "source")
	resourceCopySymlinkBenchmarkRoot = JoinPath(resourceCopySymlinkRoot, "benchmark")
}

func TestCopyFile(t *testing.T) {
	outputRoot := resourceCopyFileOutputRoot

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		inputPath  string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, outputRoot, emptyStr, emptyStr, true},
		{"Source got permission denied", resourceCopyFileFileMap["NonePermission"], JoinPath(outputRoot, "whatever.txt"), emptyStr, emptyStr, true},
		{"Source file not exist", JoinPath(resourceCopyFileRoot, "__not_exist__"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a dir", resourceCopyFileDirMap["ContentDir"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a symlink to file", resourceCopyFileFileMap["SymlinkFile"], outputRoot, resourceCopyFileFileMap["LargeText"], JoinPath(outputRoot, "soft-link.txt"), false},
		{"Source is a symlink to symlink", resourceCopyFileFileMap["SymlinkLink"], outputRoot, resourceCopyFileFileMap["LargeText"], JoinPath(outputRoot, "soft-link.txt"), false},
		{"Source is a symlink to dir", resourceCopyFileFileMap["SymlinkDir"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a broken symlink", resourceCopyFileFileMap["SymlinkBroken"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a circular symlink", resourceCopyFileFileMap["SymlinkCircular"], outputRoot, emptyStr, emptyStr, true},

		{"Destination is empty", resourceCopyFileFileMap["SmallText"], emptyStr, emptyStr, emptyStr, true},
		{"Destination is a dir", resourceCopyFileFileMap["SmallText"], resourceCopyFileDirMap["Out_ExistingDir"], resourceCopyFileFileMap["SmallText"], JoinPath(resourceCopyFileDirMap["Out_ExistingDir"], "small-text.txt"), false},
		{"Destination is a file", resourceCopyFileFileMap["SmallText"], resourceCopyFileFileMap["Out_ExistingFile"], resourceCopyFileFileMap["SmallText"], resourceCopyFileFileMap["Out_ExistingFile"], false},
		{"Destination is a symlink to file", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "link-file.txt"), resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "text-file.txt"), false},
		{"Destination is a symlink to directory", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "link-dir"), resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "directory", "small-text.txt"), false},
		{"Destination is a broken symlink", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "link-broken"), resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "wtf-broken-file"), false},
		{"Destination is a circular symlink", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "link-circular"), emptyStr, emptyStr, true},
		{"Destination got permission denied", resourceCopyFileFileMap["SmallText"], resourceCopyFileFileMap["Out_NonePermission"], emptyStr, emptyStr, true},
		{"Destination file not exist", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "not-exist-file.txt"), resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "not-exist-file.txt"), false},
		{"Destination dir not exist", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "not-exist-dir", "not-exist-file.txt"), emptyStr, emptyStr, true},

		{"Copy empty file", resourceCopyFileFileMap["EmptyFile"], JoinPath(outputRoot, "empty-file.txt"), resourceCopyFileFileMap["EmptyFile"], JoinPath(outputRoot, "empty-file.txt"), false},
		{"Copy small text file", resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "small-text.txt"), resourceCopyFileFileMap["SmallText"], JoinPath(outputRoot, "small-text.txt"), false},
		{"Copy large text file", resourceCopyFileFileMap["LargeText"], JoinPath(outputRoot, "large-text.txt"), resourceCopyFileFileMap["LargeText"], JoinPath(outputRoot, "large-text.txt"), false},
		{"Copy png image file", resourceCopyFileFileMap["PngImage"], JoinPath(outputRoot, "image.png"), resourceCopyFileFileMap["PngImage"], JoinPath(outputRoot, "image.png"), false},
		{"Copy svg image file", resourceCopyFileFileMap["SvgImage"], JoinPath(outputRoot, "image.svg"), resourceCopyFileFileMap["SvgImage"], JoinPath(outputRoot, "image.svg"), false},
		{"Source and destination are exactly the same", resourceCopyFileFileMap["SmallText"], resourceCopyFileFileMap["SmallText"], emptyStr, emptyStr, true},
		{"Source and destination are actually the same", resourceCopyFileFileMap["SmallText"], resourceCopyFileRoot, emptyStr, emptyStr, true},
		{"Source and inferred destination(dir) use the same name: can't overwrite dir", resourceCopyFileFileMap["SameName"], outputRoot, emptyStr, emptyStr, true},
		{"Source and inferred destination(file) use the same name: overwrite the file", resourceCopyFileFileMap["SameName2"], outputRoot, resourceCopyFileFileMap["SameName2"], resourceCopyFileFileMap["Out_SameName2"], false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			if err := CopyFile(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameFileContent(tt.inputPath, tt.outputPath)
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
	for name, path := range resourceCopyFileFileMap {
		if strings.HasPrefix(name, "Out_") {
			continue
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = CopyFile(path, resourceCopyFileBenchmarkRoot)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	outputRoot := resourceCopyDirOutputRoot
	expectedOutputRoot := JoinPath(resourceCopyDirRoot, "destination")

	tests := []struct {
		name         string
		srcPath      string
		destPath     string
		expectedPath string
		actualPath   string
		wantErr      bool
	}{
		{"Source is empty", emptyStr, outputRoot, emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(resourceCopyDirSourceMap["EmptyDir"], "__not_found__"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a file", resourceCopyDirSourceMap["TextFile"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a symlink to directory", resourceCopyDirSourceMap["SymlinkDir"], JoinPath(outputRoot, "out-link-dir"), resourceCopyDirSourceMap["MiscDir"], JoinPath(outputRoot, "out-link-dir"), false},
		{"Source is a symlink to file", resourceCopyDirSourceMap["SymlinkFile"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a broken symlink", resourceCopyDirSourceMap["SymlinkBroken"], outputRoot, emptyStr, emptyStr, true},
		{"Source is a circular symlink", resourceCopyDirSourceMap["SymlinkCircular"], outputRoot, emptyStr, emptyStr, true},
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
		{"Destination is a symlink to directory", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "link-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "actual-dir", "one-file-dir"), false},
		{"Destination is a symlink to file", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "existing-link.txt"), emptyStr, emptyStr, true},
		{"Destination is a broken symlink", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "link-broken"), emptyStr, emptyStr, true},
		{"Destination is a circular symlink", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "link-circular"), emptyStr, emptyStr, true},
		{"Destination and its parent don't exist", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "non-exist", "non-exist-nested"), emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent does", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "nested-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "nested-dir"), false},
		{"Destination directory exists, and it's empty", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "empty-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist", "empty-dir", "one-file-dir"), false},
		{"Destination directory exists and already contains files", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-other"), JoinPath(expectedOutputRoot, "exist-other"), JoinPath(outputRoot, "exist-other", "one-file-dir"), false},
		{"Destination directory exists and already contains the same source", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-same"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-same", "one-file-dir"), false},
		{"Destination directory exists and contains a file with the same name and no permissions", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-no-perm-file"), emptyStr, emptyStr, true},
		{"Destination directory exists and contains a directory with the same name and no permissions", resourceCopyDirSourceMap["MiscDir"], JoinPath(outputRoot, "exist-no-perm-dir"), emptyStr, emptyStr, true},
		{"Destination directory exists and contains a symlink with the same name", resourceCopyDirSourceMap["OnlySymlinks"], JoinPath(outputRoot, "exist-symlink"), resourceCopyDirSourceMap["OnlySymlinks"], JoinPath(outputRoot, "exist-symlink", "only-symlinks"), false},

		{"Source and destination are exactly the same", resourceCopyDirSourceMap["OneFileDir"], resourceCopyDirSourceMap["OneFileDir"], emptyStr, emptyStr, true},
		{"Source and destination are actually the same", resourceCopyDirSourceMap["OneFileDir"], resourceCopyDirSourceRoot, emptyStr, emptyStr, true},
		{"Source and inferred destination(file) use the same name: can't overwrite file", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-file"), emptyStr, emptyStr, true},
		{"Source and inferred destination(dir) use the same name: overwrite the dir", resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-dir"), resourceCopyDirSourceMap["OneFileDir"], JoinPath(outputRoot, "exist-dir", "one-file-dir"), false},
		{"Source is the parent of destination directory", JoinPath(outputRoot, "infinite1"), JoinPath(outputRoot, "infinite1", "infinite2"), emptyStr, emptyStr, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			if err := CopyDir(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameDirEntries(tt.expectedPath, tt.actualPath)
				if err != nil {
					t.Errorf("CopyDir() got error while comparing the directories: %v, %v, error: %v", tt.expectedPath, tt.actualPath, err)
				} else if !same {
					t.Errorf("CopyDir() the directories are not the same: %v, %v", tt.expectedPath, tt.actualPath)
					return
				}
			}
		})
	}
}

func BenchmarkCopyDir(b *testing.B) {
	for name, path := range resourceCopyDirSourceMap {
		outputPath := JoinPath(resourceCopyDirBenchmarkRoot, name)
		if err := os.MkdirAll(outputPath, defaultDirectoryFileMode); err != nil {
			b.Errorf("failed to create the directory for output: %v, error: %v", outputPath, err)
			continue
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = CopyDir(path, outputPath)
			}
		})
	}
}

func TestCopySymlink(t *testing.T) {
	outputRoot := resourceCopySymlinkOutputRoot
	existRoot := JoinPath(resourceCopySymlinkOutputRoot, "exist")
	sourceRoot := resourceCopySymlinkSourceRoot

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		inputPath  string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, outputRoot, emptyStr, emptyStr, true},
		{"Source path is inferred", joinPathNoClean(sourceRoot, "..", "source", "link-file.txt"), JoinPath(outputRoot, "link1.txt"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "link1.txt"), false},
		{"Source got permission denied", JoinPath(sourceRoot, "no_perm_file"), outputRoot, emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(sourceRoot, "__not_exist__"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a file", JoinPath(sourceRoot, "text-file.txt"), outputRoot, emptyStr, emptyStr, true},
		{"Source is an empty directory", JoinPath(sourceRoot, "empty-dir"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a directory with content", JoinPath(sourceRoot, "one-file-dir"), outputRoot, emptyStr, emptyStr, true},
		{"Source is a symlink to file", JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "link2.txt"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "link2.txt"), false},
		{"Source is a symlink to directory", JoinPath(sourceRoot, "link-dir"), JoinPath(outputRoot, "link3.txt"), JoinPath(sourceRoot, "link-dir"), JoinPath(outputRoot, "link3.txt"), false},
		{"Source is a circular symlink", JoinPath(sourceRoot, "link-circular"), JoinPath(outputRoot, "link4.txt"), JoinPath(sourceRoot, "link-circular"), JoinPath(outputRoot, "link4.txt"), false},
		{"Source is a broken symlink", JoinPath(sourceRoot, "link-broken"), JoinPath(outputRoot, "link5.txt"), JoinPath(sourceRoot, "link-broken"), JoinPath(outputRoot, "link5.txt"), false},
		{"Source and destination path are the same", JoinPath(sourceRoot, "link-file.txt"), JoinPath(sourceRoot, "link-file.txt"), emptyStr, emptyStr, true},

		{"Destination is empty", JoinPath(sourceRoot, "link-file.txt"), emptyStr, emptyStr, emptyStr, true},
		{"Destination path is inferred", JoinPath(sourceRoot, "link-file.txt"), joinPathNoClean(outputRoot, "..", "output", "out1.txt"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "out1.txt"), false},
		{"Destination got permission denied", JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Destination exists and it's a file", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "text.txt"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "text.txt"), false},
		{"Destination exists and it's an empty directory", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "empty-dir"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "empty-dir", "link-file.txt"), false},
		{"Destination exists and it's a directory with content", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "one-file-dir"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "one-file-dir", "link-file.txt"), false},
		{"Destination exists and it's a symlink to file", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-file.txt"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-file.txt"), false},
		{"Destination exists and it's a symlink to directory", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-dir"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-dir"), false},
		{"Destination exists and it's a broken symlink", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-broken"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-broken"), false},
		{"Destination exists and it's a circular symlink", JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-circular"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(existRoot, "link-circular"), false},
		{"Destination doesn't exist but its parent does", JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "target.lnk"), JoinPath(sourceRoot, "link-file.txt"), JoinPath(outputRoot, "target.lnk"), false},
		{"Destination and its parent don't exist", JoinPath(sourceRoot, "link-file.txt"), JoinPath(resourceCopySymlinkRoot, "missing1", "missing2"), emptyStr, emptyStr, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			if err := CopySymlink(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopySymlink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameSymlinkContent(tt.inputPath, tt.outputPath)
				if err != nil {
					t.Errorf("CopySymlink() got error while comparing the files: %v, %v, error: %v", tt.inputPath, tt.outputPath, err)
				} else if !same {
					t.Errorf("CopySymlink() the files are not the same: %v, %v", tt.inputPath, tt.outputPath)
					return
				}
			}
		})
	}
}

func BenchmarkCopySymlink(b *testing.B) {
	inputPath, outputPath := JoinPath(resourceCopySymlinkSourceRoot, "link-file.txt"), resourceCopySymlinkBenchmarkRoot
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CopySymlink(inputPath, outputPath)
	}
}
