package yos

import (
	"os"
	"testing"

	"github.com/1set/gut/ystring"
)

var (
	resourceSizeRoot      string
	resourceSizeSourceMap map[string]string
)

func init() {
	resourceSizeRoot = JoinPath(testResourceRoot, "yos", "size")
	resourceSizeSourceMap = map[string]string{
		"EmptyFile":        JoinPath(resourceSizeRoot, "empty.txt"),
		"TextFile":         JoinPath(resourceSizeRoot, "text.txt"),
		"ImageFile":        JoinPath(resourceCopyFileRoot, "image.png"),
		"LargeText":        JoinPath(resourceCopyFileRoot, "large-text.txt"),
		"XlargeText":       JoinPath(resourceCopyFileRoot, "xlarge-text.txt"),
		"BlankSymlink":     JoinPath(resourceSizeRoot, "lonely-link"),
		"BrokenSymlink":    JoinPath(resourceSizeRoot, "link-broken"),
		"CircularSymlink":  JoinPath(resourceSizeRoot, "link-circular"),
		"FileSymlink":      JoinPath(resourceSizeRoot, "link.txt"),
		"DirSymlink":       JoinPath(resourceSizeRoot, "link-dir"),
		"EmptyFileSymlink": JoinPath(resourceSizeRoot, "link-empty.txt"),
		"EmptyDirSymlink":  JoinPath(resourceSizeRoot, "link-empty-dir"),
		"LinkFileSymlink":  JoinPath(resourceSizeRoot, "link2.txt"),
		"LinkDirSymlink":   JoinPath(resourceSizeRoot, "link2-dir"),
		"EmptyDir":         JoinPath(resourceSizeRoot, "empty-dir"),
		"OneFileDir":       JoinPath(resourceSizeRoot, "one-file-dir"),
		"DirsDir":          JoinPath(resourceSizeRoot, "only-dirs"),
		"SymlinksDir":      JoinPath(resourceSizeRoot, "only-symlinks"),
		"MiscDir":          JoinPath(resourceSizeRoot, "misc"),
	}
}

func TestGetFileSize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantSize int64
		wantErr  bool
	}{
		{"Path is empty", emptyStr, 0, true},
		{"Source is missing", "__not_exist_file__", 0, true},
		{"Source is an empty file", resourceSizeSourceMap["EmptyFile"], 0, false},
		{"Source is a small text file", resourceSizeSourceMap["TextFile"], 99, false},
		{"Source is an image file", resourceSizeSourceMap["ImageFile"], 631513, false},
		{"Source is a large text file", resourceSizeSourceMap["LargeText"], 63000000, false},
		{"Source is an extra large text file", resourceSizeSourceMap["XlargeText"], 94500000, false},

		{"Source is a blank symlink", resourceSizeSourceMap["BlankSymlink"], 0, true},
		{"Source is a broken symlink", resourceSizeSourceMap["BrokenSymlink"], 0, true},
		{"Source is a circular symlink", resourceSizeSourceMap["CircularSymlink"], 0, true},
		{"Source is a symlink to file", resourceSizeSourceMap["FileSymlink"], 99, false},
		{"Source is a symlink to directory", resourceSizeSourceMap["DirSymlink"], 0, true},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], 0, true},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], 0, true},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], 0, true},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], 0, true},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotSize, err := GetFileSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileSize() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}

			if gotSize != tt.wantSize {
				t.Errorf("GetFileSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func BenchmarkGetFileSize(b *testing.B) {
	files := []string{
		resourceSizeSourceMap["FileSymlink"],
		resourceSizeSourceMap["EmptyFileSymlink"],
		resourceSizeSourceMap["TextFile"],
		resourceSizeSourceMap["EmptyFile"],
		resourceSizeSourceMap["ImageFile"],
		resourceSizeSourceMap["LargeText"],
	}
	for _, path := range files {
		name := ystring.TrimBeforeLast(path, string(os.PathSeparator))
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = GetFileSize(path)
			}
		})
	}
}

func TestGetSymlinkSize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantSize int64
		wantErr  bool
	}{
		{"Path is empty", emptyStr, 0, true},
		{"Source is missing", "__not_exist_link__", 0, true},
		{"Source is an empty file", resourceSizeSourceMap["EmptyFile"], 0, true},
		{"Source is a small text file", resourceSizeSourceMap["TextFile"], 0, true},
		{"Source is an image file", resourceSizeSourceMap["ImageFile"], 0, true},
		{"Source is a large text file", resourceSizeSourceMap["LargeText"], 0, true},
		{"Source is an extra large text file", resourceSizeSourceMap["XlargeText"], 0, true},

		{"Source is a blank symlink (non-Windows)", resourceSizeSourceMap["BlankSymlink"], 1, false},
		{"Source is a broken symlink (non-Windows)", resourceSizeSourceMap["BrokenSymlink"], 7, false},
		{"Source is a circular symlink (non-Windows)", resourceSizeSourceMap["CircularSymlink"], 13, false},
		{"Source is a symlink to file (non-Windows)", resourceSizeSourceMap["FileSymlink"], 8, false},
		{"Source is a symlink to directory (non-Windows)", resourceSizeSourceMap["DirSymlink"], 4, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], 0, true},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], 0, true},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], 0, true},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], 0, true},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotSize, err := GetSymlinkSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymlinkSize() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}

			if gotSize != tt.wantSize {
				t.Errorf("GetSymlinkSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func BenchmarkGetSymlinkSize(b *testing.B) {
	links := []string{
		resourceSizeSourceMap["BlankSymlink"],
		resourceSizeSourceMap["BrokenSymlink"],
		resourceSizeSourceMap["CircularSymlink"],
		resourceSizeSourceMap["FileSymlink"],
		resourceSizeSourceMap["DirSymlink"],
	}
	for _, path := range links {
		name := ystring.TrimBeforeLast(path, string(os.PathSeparator))
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = GetSymlinkSize(path)
			}
		})
	}
}

func TestGetDirSize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantSize int64
		wantErr  bool
	}{
		{"Path is empty", emptyStr, 0, true},
		{"Source is missing", "__not_exist_link__", 0, true},
		{"Source is an empty file", resourceSizeSourceMap["EmptyFile"], 0, true},
		{"Source is a small text file", resourceSizeSourceMap["TextFile"], 0, true},
		{"Source is an image file", resourceSizeSourceMap["ImageFile"], 0, true},
		{"Source is a large text file", resourceSizeSourceMap["LargeText"], 0, true},
		{"Source is an extra large text file", resourceSizeSourceMap["XlargeText"], 0, true},

		{"Source is a blank symlink", resourceSizeSourceMap["BlankSymlink"], 0, true},
		{"Source is a broken symlink", resourceSizeSourceMap["BrokenSymlink"], 0, true},
		{"Source is a circular symlink", resourceSizeSourceMap["CircularSymlink"], 0, true},
		{"Source is a symlink to file", resourceSizeSourceMap["FileSymlink"], 0, true},
		{"Source is a symlink to directory (non-Windows)", resourceSizeSourceMap["DirSymlink"], 309, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], 0, false},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], 3, false},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], 0, false},
		{"Source is a directory containing only symlinks (non-Windows)", resourceSizeSourceMap["SymlinksDir"], 29, false},
		{"Source is a directory containing files, symlinks and directories (non-Windows)", resourceSizeSourceMap["MiscDir"], 309, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotSize, err := GetDirSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDirSize() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}

			if gotSize != tt.wantSize {
				t.Errorf("GetDirSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func BenchmarkGetDirSize(b *testing.B) {
	dirs := []string{
		resourceSizeSourceMap["EmptyDir"],
		resourceSizeSourceMap["OneFileDir"],
		resourceSizeSourceMap["DirsDir"],
		resourceSizeSourceMap["SymlinksDir"],
		resourceSizeSourceMap["MiscDir"],
		resourceSizeSourceMap["DirSymlink"],
		resourceSizeSourceMap["EmptyDirSymlink"],
	}
	for _, path := range dirs {
		name := ystring.TrimBeforeLast(path, string(os.PathSeparator))
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = GetDirSize(path)
			}
		})
	}
}
