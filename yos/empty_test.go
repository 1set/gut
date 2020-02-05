package yos

import (
	"os"
	"testing"

	"github.com/1set/gut/ystring"
)

func TestIsFileEmpty(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantEmpty bool
		wantErr   bool
	}{
		{"Path is empty", emptyStr, false, true},
		{"Source is missing", "__not_exist_file__", false, true},
		{"Source is an empty file", resourceSizeSourceMap["EmptyFile"], true, false},
		{"Source is a small text file", resourceSizeSourceMap["TextFile"], false, false},
		{"Source is an image file", resourceSizeSourceMap["ImageFile"], false, false},
		{"Source is a large text file", resourceSizeSourceMap["LargeText"], false, false},
		{"Source is an extra large text file", resourceSizeSourceMap["XlargeText"], false, false},

		{"Source is a blank symlink", resourceSizeSourceMap["BlankSymlink"], false, true},
		{"Source is a broken symlink", resourceSizeSourceMap["BrokenSymlink"], false, true},
		{"Source is a circular symlink", resourceSizeSourceMap["CircularSymlink"], false, true},
		{"Source is a symlink to directory", resourceSizeSourceMap["DirSymlink"], false, true},
		{"Source is a symlink to file", resourceSizeSourceMap["FileSymlink"], false, false},
		{"Source is a symlink to empty file", resourceSizeSourceMap["EmptyFileSymlink"], true, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], false, true},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], false, true},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], false, true},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], false, true},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotEmpty, err := IsFileEmpty(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileEmpty() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}

			if gotEmpty != tt.wantEmpty {
				t.Errorf("IsFileEmpty() gotEmpty = %v, want %v", gotEmpty, tt.wantEmpty)
			}
		})
	}
}

func BenchmarkIsFileEmpty(b *testing.B) {
	files := []string{
		resourceSizeSourceMap["FileSymlink"],
		resourceSizeSourceMap["EmptyFileSymlink"],
		resourceSizeSourceMap["TextFile"],
		resourceSizeSourceMap["EmptyFile"],
	}
	for _, path := range files {
		name := ystring.TrimBeforeLast(path, string(os.PathSeparator))
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = IsFileEmpty(path)
			}
		})
	}
}

func TestIsDirEmpty(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantEmpty bool
		wantErr   bool
	}{
		{"Path is empty", emptyStr, false, true},
		{"Source is missing", "__not_exist_file__", false, true},
		{"Source is an empty file", resourceSizeSourceMap["EmptyFile"], false, true},
		{"Source is a small text file", resourceSizeSourceMap["TextFile"], false, true},
		{"Source is an image file", resourceSizeSourceMap["ImageFile"], false, true},
		{"Source is a large text file", resourceSizeSourceMap["LargeText"], false, true},
		{"Source is an extra large text file", resourceSizeSourceMap["XlargeText"], false, true},

		{"Source is a blank symlink", resourceSizeSourceMap["BlankSymlink"], false, true},
		{"Source is a broken symlink", resourceSizeSourceMap["BrokenSymlink"], false, true},
		{"Source is a circular symlink", resourceSizeSourceMap["CircularSymlink"], false, true},
		{"Source is a symlink to file", resourceSizeSourceMap["FileSymlink"], false, true},
		{"Source is a symlink to directory", resourceSizeSourceMap["DirSymlink"], false, false},
		{"Source is a symlink to empty directory", resourceSizeSourceMap["EmptyDirSymlink"], true, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], true, false},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], false, false},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], false, false},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], false, false},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotEmpty, err := IsDirEmpty(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDirEmpty() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}

			if gotEmpty != tt.wantEmpty {
				t.Errorf("IsDirEmpty() gotEmpty = %v, want %v", gotEmpty, tt.wantEmpty)
			}
		})
	}
}

func BenchmarkIsDirEmpty(b *testing.B) {
	dirs := []string{
		resourceSizeSourceMap["DirSymlink"],
		resourceSizeSourceMap["EmptyDirSymlink"],
		resourceSizeSourceMap["MiscDir"],
		resourceSizeSourceMap["EmptyDir"],
	}
	for _, path := range dirs {
		name := ystring.TrimBeforeLast(path, string(os.PathSeparator))
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = IsDirEmpty(path)
			}
		})
	}
}
