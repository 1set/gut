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
		"LargeText":       JoinPath(resourceCopyFileRoot, "large-text.txt"),
		"XlargeText":      JoinPath(resourceCopyFileRoot, "xlarge-text.txt"),
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

func TestDirSize(t *testing.T) {
	t.Logf("root: %v", resourceSizeRoot)
	t.Logf("map: %v", resourceSizeSourceMap)

	tests := []struct {
		name     string
		path     string
		wantSize int64
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := DirSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSize != tt.wantSize {
				t.Errorf("DirSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestFileSize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantSize int64
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := FileSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSize() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := SymlinkSize(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("SymlinkSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSize != tt.wantSize {
				t.Errorf("SymlinkSize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}
