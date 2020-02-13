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

func TestChExeDir(t *testing.T) {
	err := ChExeDir()
	if err != nil {
		t.Errorf("ChExeDir() got unexpected error: %v", err)
	}
}
