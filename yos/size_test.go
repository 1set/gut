package yos

import (
	"testing"
)

var (
	resourceSizeRoot      string
	resourceSizeSourceMap map[string]string
)

func init() {
	resourceSizeRoot = JoinPath(testResourceRoot, "yos", "size")
	resourceSizeSourceMap = map[string]string{
		"EmptyFile":       JoinPath(resourceSizeRoot, "empty.txt"),
		"TextFile":        JoinPath(resourceSizeRoot, "text.txt"),
		"ImageFile":       JoinPath(resourceCopyFileRoot, "image.png"),
		"LargeText":       JoinPath(resourceCopyFileRoot, "large-text.txt"),
		"XlargeText":      JoinPath(resourceCopyFileRoot, "xlarge-text.txt"),
		"BlankSymlink":    JoinPath(resourceSizeRoot, "lonely-link"),
		"BrokenSymlink":   JoinPath(resourceSizeRoot, "link-broken"),
		"CircularSymlink": JoinPath(resourceSizeRoot, "link-circular"),
		"FileSymlink":     JoinPath(resourceSizeRoot, "link.txt"),
		"DirSymlink":      JoinPath(resourceSizeRoot, "link-dir"),
		"EmptyDir":        JoinPath(resourceSizeRoot, "empty-dir"),
		"OneFileDir":      JoinPath(resourceSizeRoot, "one-file-dir"),
		"DirsDir":         JoinPath(resourceSizeRoot, "only-dirs"),
		"SymlinksDir":     JoinPath(resourceSizeRoot, "only-symlinks"),
		"MiscDir":         JoinPath(resourceSizeRoot, "misc"),
	}
}

func TestFileSize(t *testing.T) {
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
			gotSize, err := FileSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				expectedErrorCheck(t, err)
			}
			if gotSize != tt.wantSize {
				t.Errorf("FileSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestSymlinkSize(t *testing.T) {
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

		{"Source is a blank symlink", resourceSizeSourceMap["BlankSymlink"], 1, false},
		{"Source is a broken symlink", resourceSizeSourceMap["BrokenSymlink"], 7, false},
		{"Source is a circular symlink", resourceSizeSourceMap["CircularSymlink"], 13, false},
		{"Source is a symlink to file", resourceSizeSourceMap["FileSymlink"], 8, false},
		{"Source is a symlink to directory", resourceSizeSourceMap["DirSymlink"], 4, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], 0, true},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], 0, true},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], 0, true},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], 0, true},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := SymlinkSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("SymlinkSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				expectedErrorCheck(t, err)
			}
			if gotSize != tt.wantSize {
				t.Errorf("SymlinkSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestDirSize(t *testing.T) {
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
		{"Source is a symlink to directory", resourceSizeSourceMap["DirSymlink"], 309, false},

		{"Source is an empty directory", resourceSizeSourceMap["EmptyDir"], 0, false},
		{"Source is a directory containing one file", resourceSizeSourceMap["OneFileDir"], 3, false},
		{"Source is a directory containing only directories", resourceSizeSourceMap["DirsDir"], 0, false},
		{"Source is a directory containing only symlinks", resourceSizeSourceMap["SymlinksDir"], 29, false},
		{"Source is a directory containing files, symlinks and directories", resourceSizeSourceMap["MiscDir"], 309, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := DirSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				expectedErrorCheck(t, err)
			}
			if gotSize != tt.wantSize {
				t.Errorf("DirSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}
