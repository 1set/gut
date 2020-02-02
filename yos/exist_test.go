package yos

import (
	"os"
	"strings"
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
		{"Check doc file", "doc.go", true},
		{"Check current dir", ".", true},
		{"Check symlink origin", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Check symlink of file", JoinPath(resourceExistRoot, "symlink.txt"), true},
		{"Check symlink of dir", JoinPath(resourceExistRoot, "dir_link"), true},
		{"Check symlink of symlink", JoinPath(resourceExistRoot, "2symlink.txt"), true},
		{"Check broken symlink", JoinPath(resourceExistRoot, "broken.txt"), false},
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
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check doc file", "doc.go", true},
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
			gotExist := IsFileExist(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("IsFileExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsFileExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsFileExist("doc.go")
	}
}

func TestIsDirExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check doc file", "doc.go", false},
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

			gotExist := IsDirExist(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("IsDirExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsDirExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsDirExist(resourceExistRoot)
	}
}

func TestIsSymlinkExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
	}{
		{"Check missing", "__do_not_exist__", false},
		{"Check doc file", "doc.go", false},
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
			gotExist := IsSymlinkExist(tt.path)
			if gotExist != tt.wantExist {
				t.Errorf("IsSymlinkExist() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func BenchmarkIsSymlinkExist(b *testing.B) {
	path := JoinPath(resourceExistRoot, "symlink.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsSymlinkExist(path)
	}
}

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name string
		elem []string
		want string
	}{
		{"Nil", nil, ""},
		{"Empty", []string{}, ""},
		{"Single part", []string{"abc"}, "abc"},
		{"Two parts", []string{"ab", "cd"}, strings.Join([]string{"ab", "cd"}, string(os.PathSeparator))},
		{"Three parts", []string{"ab", "cd", "ef"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"Contains heading empty part", []string{"", "cd", "ef"}, strings.Join([]string{"cd", "ef"}, string(os.PathSeparator))},
		{"Contains trailing empty part", []string{"ab", "cd", ""}, strings.Join([]string{"ab", "cd"}, string(os.PathSeparator))},
		{"Contains empty part in the middle", []string{"abc", "", "ef"}, strings.Join([]string{"abc", "ef"}, string(os.PathSeparator))},
		{"Contains trailing slash", []string{"ab/", "cd/", "ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"Contains heading slash", []string{"ab", "/cd", "/ef"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"Contains heading & trailing slash", []string{"ab/", "/cd/", "/ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
		{"Contains extra slash", []string{"ab//", "//cd//", "//ef/"}, strings.Join([]string{"ab", "cd", "ef"}, string(os.PathSeparator))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPath(tt.elem...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
