package yos

import (
	"os"
	"strings"
	"testing"
)

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
		{"Contains heading dot part", []string{".", "cd", "ef"}, strings.Join([]string{"cd", "ef"}, string(os.PathSeparator))},
		{"Contains trailing empty part", []string{"ab", "cd", ""}, strings.Join([]string{"ab", "cd"}, string(os.PathSeparator))},
		{"Contains trailing dot part", []string{"ab", "cd", "."}, strings.Join([]string{"ab", "cd"}, string(os.PathSeparator))},
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

func BenchmarkJoinPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JoinPath("", "ab", "/cd", "ef/", "gh", ".")
	}
}

func TestChExeDir(t *testing.T) {
	err := ChangeExeDir()
	if err != nil {
		t.Errorf("ChangeExeDir() got unexpected error: %v", err)
	}
}

func BenchmarkChExeDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ChangeExeDir()
	}
}

func TestMakeDir(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Path is empty", "", true},
		{"Path dir exists", resourceExistRoot, false},
		{"Path dir doesn't exist but its parent does", JoinPath(resourceExistRoot, "mkdir-new-dir1"), false},
		{"Path dir and its parent don't exist", JoinPath(resourceExistRoot, "mkdir-new-dir2", "mkdir-new-dir2b", "mkdir-new-dir2c"), false},
		{"Path file exists", JoinPath(resourceExistRoot, "origin_file.txt"), true},
		{"Path symlink exists", JoinPath(resourceExistRoot, "symlink.txt"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeDir(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("MakeDir() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				expectedErrorCheck(t, err)
			}
		})
	}
}

func BenchmarkMakeDir(b *testing.B) {
	path := JoinPath(resourceExistRoot, "mkdir-bench-new")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeDir(path)
	}
}
