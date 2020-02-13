package yos

import (
	"testing"
)

var resourceExistRoot string

func init() {
	resourceExistRoot = JoinPath(testResourceRoot, "yos", "exist")
}

func TestIsExistOrNot(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		exist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check text file", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Check current dir", ".", true},
		{"Check symlink origin", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Check symlink of file", JoinPath(resourceExistRoot, "symlink.txt"), true},
		{"Check symlink of dir", JoinPath(resourceExistRoot, "dir_link"), true},
		{"Check symlink of symlink", JoinPath(resourceExistRoot, "2symlink.txt"), true},
		{"Check broken symlink", JoinPath(resourceExistRoot, "broken.txt"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.path); got != tt.exist {
				t.Errorf("Exist() = %v, want %v", got, tt.exist)
			}
			if got := NotExist(tt.path); got != !tt.exist {
				t.Errorf("NotExist() = %v, want %v", got, !tt.exist)
			}
		})
	}
}

func BenchmarkIsExist(b *testing.B) {
	path := JoinPath(resourceExistRoot, "origin_file.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Exist(path)
	}
}

func BenchmarkIsNotExist(b *testing.B) {
	path := JoinPath(resourceExistRoot, "origin_file.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NotExist(path)
	}
}

func TestIsFileExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check text file", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Check current dir", ".", false},
		{"Check symlink dir", JoinPath(resourceExistRoot), false},
		{"Check symlink origin file", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Check symlink of file", JoinPath(resourceExistRoot, "symlink.txt"), true},
		{"Check symlink of symlink of file", JoinPath(resourceExistRoot, "2symlink.txt"), true},
		{"Check symlink origin dir", JoinPath(resourceExistRoot, "target_dir"), false},
		{"Check symlink of dir", JoinPath(resourceExistRoot, "dir_link"), false},
		{"Check symlink of symlink of dir", JoinPath(resourceExistRoot, "2dir_link"), false},
		{"Check broken file symlink", JoinPath(resourceExistRoot, "broken.txt"), false},
		{"Check broken dir symlink", JoinPath(resourceExistRoot, "broken2.txt"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist := ExistFile(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("ExistFile() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsFileExist(b *testing.B) {
	path := JoinPath(resourceExistRoot, "origin_file.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ExistFile(path)
	}
}

func TestIsDirExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check text file", JoinPath(resourceExistRoot, "origin_file.txt"), false},
		{"Check current dir", ".", true},
		{"Check symlink dir", JoinPath(resourceExistRoot), true},
		{"Check symlink origin file", JoinPath(resourceExistRoot, "origin_file.txt"), false},
		{"Check symlink of file", JoinPath(resourceExistRoot, "symlink.txt"), false},
		{"Check symlink of symlink of file", JoinPath(resourceExistRoot, "2symlink.txt"), false},
		{"Check symlink origin dir", JoinPath(resourceExistRoot, "target_dir"), true},
		{"Check symlink of dir", JoinPath(resourceExistRoot, "dir_link"), true},
		{"Check symlink of symlink of dir (non-Windows)", JoinPath(resourceExistRoot, "2dir_link"), true},
		{"Check broken file symlink", JoinPath(resourceExistRoot, "broken.txt"), false},
		{"Check broken dir symlink", JoinPath(resourceExistRoot, "broken2.txt"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preconditionCheck(t, tt.name)

			gotExist := ExistDir(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("ExistDir() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsDirExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ExistDir(resourceExistRoot)
	}
}

func TestIsSymlinkExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check text file", JoinPath(resourceExistRoot, "origin_file.txt"), false},
		{"Check current dir", ".", false},
		{"Check symlink dir", JoinPath(resourceExistRoot), false},
		{"Check symlink origin file", JoinPath(resourceExistRoot, "origin_file.txt"), false},
		{"Check symlink of file", JoinPath(resourceExistRoot, "symlink.txt"), true},
		{"Check symlink of symlink of file", JoinPath(resourceExistRoot, "2symlink.txt"), true},
		{"Check symlink origin dir", JoinPath(resourceExistRoot, "target_dir"), false},
		{"Check symlink of dir", JoinPath(resourceExistRoot, "dir_link"), true},
		{"Check symlink of symlink of dir", JoinPath(resourceExistRoot, "2dir_link"), true},
		{"Check broken file symlink", JoinPath(resourceExistRoot, "broken.txt"), true},
		{"Check broken dir symlink", JoinPath(resourceExistRoot, "broken2.txt"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist := ExistSymlink(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("ExistSymlink() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsSymlinkExist(b *testing.B) {
	path := JoinPath(resourceExistRoot, "symlink.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ExistSymlink(path)
	}
}
