package yos

import (
	"os"
	"strings"
	"testing"
)

var TestCaseRootSymlink string

func init() {
	TestCaseRootSymlink = "/Users/vej/Desktop/temp-test-symlink"
}

func TestIsExistOrNot(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		exist bool
	}{
		{"check missing", "__do_not_exist__", false},
		{"check doc file", "doc.go", true},
		{"check current dir", ".", true},
		{"check symlink origin", JoinPath(TestCaseRootSymlink, "origin_file.txt"), true},
		{"check symlink of file", JoinPath(TestCaseRootSymlink, "symlink.txt"), true},
		{"check symlink of dir", JoinPath(TestCaseRootSymlink, "dir_link"), true},
		{"check symlink of symlink", JoinPath(TestCaseRootSymlink, "2symlink.txt"), true},
		{"check broken symlink", JoinPath(TestCaseRootSymlink, "broken.txt"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.path); got != tt.exist {
				t.Errorf("IsExist() = %v, want %v", got, tt.exist)
			}
			if got := IsNotExist(tt.path); got != !tt.exist {
				t.Errorf("IsNotExist() = %v, want %v", got, !tt.exist)
			}
		})
	}
}

func BenchmarkIsExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsExist("doc.go")
	}
}

func BenchmarkIsNotExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNotExist("doc.go")
	}
}

func TestIsFileExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
		wantErr   bool
	}{
		{"check missing", "__do_not_exist__", false, true},
		{"check doc file", "doc.go", true, false},
		{"check current dir", ".", false, true},
		{"check symlink dir", JoinPath(TestCaseRootSymlink), false, true},
		{"check symlink origin file", JoinPath(TestCaseRootSymlink, "origin_file.txt"), true, false},
		{"check symlink of file", JoinPath(TestCaseRootSymlink, "symlink.txt"), true, false},
		{"check symlink of symlink of file", JoinPath(TestCaseRootSymlink, "2symlink.txt"), true, false},
		{"check symlink origin dir", JoinPath(TestCaseRootSymlink, "target_dir"), false, true},
		{"check symlink of dir", JoinPath(TestCaseRootSymlink, "dir_link"), false, true},
		{"check symlink of symlink of dir", JoinPath(TestCaseRootSymlink, "2dir_link"), false, true},
		{"check broken file symlink", JoinPath(TestCaseRootSymlink, "broken.txt"), false, true},
		{"check broken dir symlink", JoinPath(TestCaseRootSymlink, "broken2.txt"), false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsFileExist(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExist != tt.wantExist {
				t.Errorf("IsFileExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsFileExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = IsFileExist("doc.go")
	}
}

func TestIsDirExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
		wantErr   bool
	}{
		{"check missing", "__do_not_exist__", false, true},
		{"check doc file", "doc.go", false, true},
		{"check current dir", ".", true, false},
		{"check symlink dir", JoinPath(TestCaseRootSymlink), true, false},
		{"check symlink origin file", JoinPath(TestCaseRootSymlink, "origin_file.txt"), false, true},
		{"check symlink of file", JoinPath(TestCaseRootSymlink, "symlink.txt"), false, true},
		{"check symlink of symlink of file", JoinPath(TestCaseRootSymlink, "2symlink.txt"), false, true},
		{"check symlink origin dir", JoinPath(TestCaseRootSymlink, "target_dir"), true, false},
		{"check symlink of dir", JoinPath(TestCaseRootSymlink, "dir_link"), true, false},
		{"check symlink of symlink of dir", JoinPath(TestCaseRootSymlink, "2dir_link"), true, false},
		{"check broken file symlink", JoinPath(TestCaseRootSymlink, "broken.txt"), false, true},
		{"check broken dir symlink", JoinPath(TestCaseRootSymlink, "broken2.txt"), false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsDirExist(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDirExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExist != tt.wantExist {
				t.Errorf("IsDirExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsDirExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = IsDirExist(".")
	}
}

func TestIsSymlinkExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
		wantErr   bool
	}{
		{"check missing", "__do_not_exist__", false, true},
		{"check doc file", "doc.go", false, true},
		{"check current dir", ".", false, true},
		{"check symlink dir", JoinPath(TestCaseRootSymlink), false, true},
		{"check symlink origin file", JoinPath(TestCaseRootSymlink, "origin_file.txt"), false, true},
		{"check symlink of file", JoinPath(TestCaseRootSymlink, "symlink.txt"), true, false},
		{"check symlink of symlink of file", JoinPath(TestCaseRootSymlink, "2symlink.txt"), true, false},
		{"check symlink origin dir", JoinPath(TestCaseRootSymlink, "target_dir"), false, true},
		{"check symlink of dir", JoinPath(TestCaseRootSymlink, "dir_link"), true, false},
		{"check symlink of symlink of dir", JoinPath(TestCaseRootSymlink, "2dir_link"), true, false},
		{"check broken file symlink", JoinPath(TestCaseRootSymlink, "broken.txt"), true, false},
		{"check broken dir symlink", JoinPath(TestCaseRootSymlink, "broken2.txt"), true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsSymlinkExist(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSymlinkExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExist != tt.wantExist {
				t.Errorf("IsSymlinkExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsSymlinkExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = IsSymlinkExist("doc.go")
	}
}

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name string
		elem []string
		want string
	}{
		{"nil", nil, ""},
		{"empty", []string{}, ""},
		{"single part", []string{"abc"}, "abc"},
		{"two parts", []string{"ab", "cd"}, strings.Join([]string{"ab", "cd"}, string(os.PathSeparator))},
		{"three parts", []string{"ab", "cd", "ef"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"contains trailing slash", []string{"ab/", "cd/", "ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"contains heading slash", []string{"ab", "/cd", "/ef"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"contains heading & trailing slash", []string{"ab/", "/cd/", "/ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"contains extra slash", []string{"ab//", "//cd//", "//ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPath(tt.elem...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
