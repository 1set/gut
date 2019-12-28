package yos

import (
	"os"
	"strings"
	"testing"
)

var CaseSymlinkRoot string

func init() {
	CaseSymlinkRoot = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "symlink")
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
		{"Check symlink origin", JoinPath(CaseSymlinkRoot, "origin_file.txt"), true},
		{"Check symlink of file", JoinPath(CaseSymlinkRoot, "symlink.txt"), true},
		{"Check symlink of dir", JoinPath(CaseSymlinkRoot, "dir_link"), true},
		{"Check symlink of symlink", JoinPath(CaseSymlinkRoot, "2symlink.txt"), true},
		{"Check broken symlink", JoinPath(CaseSymlinkRoot, "broken.txt"), false},
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
		{"Check missing", "__do_not_exist__", false, true},
		{"Check doc file", "doc.go", true, false},
		{"Check current dir", ".", false, true},
		{"Check symlink dir", JoinPath(CaseSymlinkRoot), false, true},
		{"Check symlink origin file", JoinPath(CaseSymlinkRoot, "origin_file.txt"), true, false},
		{"Check symlink of file", JoinPath(CaseSymlinkRoot, "symlink.txt"), true, false},
		{"Check symlink of symlink of file", JoinPath(CaseSymlinkRoot, "2symlink.txt"), true, false},
		{"Check symlink origin dir", JoinPath(CaseSymlinkRoot, "target_dir"), false, true},
		{"Check symlink of dir", JoinPath(CaseSymlinkRoot, "dir_link"), false, true},
		{"Check symlink of symlink of dir", JoinPath(CaseSymlinkRoot, "2dir_link"), false, true},
		{"Check broken file symlink", JoinPath(CaseSymlinkRoot, "broken.txt"), false, true},
		{"Check broken dir symlink", JoinPath(CaseSymlinkRoot, "broken2.txt"), false, true},
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
		{"Check missing", "__do_not_exist__", false, true},
		{"Check doc file", "doc.go", false, true},
		{"Check current dir", ".", true, false},
		{"Check symlink dir", JoinPath(CaseSymlinkRoot), true, false},
		{"Check symlink origin file", JoinPath(CaseSymlinkRoot, "origin_file.txt"), false, true},
		{"Check symlink of file", JoinPath(CaseSymlinkRoot, "symlink.txt"), false, true},
		{"Check symlink of symlink of file", JoinPath(CaseSymlinkRoot, "2symlink.txt"), false, true},
		{"Check symlink origin dir", JoinPath(CaseSymlinkRoot, "target_dir"), true, false},
		{"Check symlink of dir", JoinPath(CaseSymlinkRoot, "dir_link"), true, false},
		// {"Check symlink of symlink of dir", JoinPath(CaseSymlinkRoot, "2dir_link"), true, false},
		{"Check broken file symlink", JoinPath(CaseSymlinkRoot, "broken.txt"), false, true},
		{"Check broken dir symlink", JoinPath(CaseSymlinkRoot, "broken2.txt"), false, true},
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
		_, _ = IsDirExist(CaseSymlinkRoot)
	}
}

func TestIsSymlinkExist(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantExist bool
		wantErr   bool
	}{
		{"Check missing", "__do_not_exist__", false, true},
		{"Check doc file", "doc.go", false, true},
		{"Check current dir", ".", false, true},
		{"Check symlink dir", JoinPath(CaseSymlinkRoot), false, true},
		{"Check symlink origin file", JoinPath(CaseSymlinkRoot, "origin_file.txt"), false, true},
		{"Check symlink of file", JoinPath(CaseSymlinkRoot, "symlink.txt"), true, false},
		{"Check symlink of symlink of file", JoinPath(CaseSymlinkRoot, "2symlink.txt"), true, false},
		{"Check symlink origin dir", JoinPath(CaseSymlinkRoot, "target_dir"), false, true},
		{"Check symlink of dir", JoinPath(CaseSymlinkRoot, "dir_link"), true, false},
		{"Check symlink of symlink of dir", JoinPath(CaseSymlinkRoot, "2dir_link"), true, false},
		{"Check broken file symlink", JoinPath(CaseSymlinkRoot, "broken.txt"), true, false},
		{"Check broken dir symlink", JoinPath(CaseSymlinkRoot, "broken2.txt"), true, false},
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
	path := JoinPath(CaseSymlinkRoot, "symlink.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = IsSymlinkExist(path)
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
