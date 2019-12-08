package yos

import (
	"os"
	"strings"
	"testing"
)

func TestIsExistOrNot(t *testing.T) {
	tests := []struct {
		name string
		path string
		exist bool
	}{
		{"check missing", "__do_not_exist__", false},
		{"check doc file", "doc.go", true},
		{"check current dir", ".", true},
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

func TestIsFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsFileExist(tt.args.path)
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

func TestIsDirExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsDirExist(tt.args.path)
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

func TestIsSymlinkExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := IsSymlinkExist(tt.args.path)
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

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name string
		elem []string
		want string
	}{
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