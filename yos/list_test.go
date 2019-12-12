package yos

import (
	"os"
	"strings"
	"testing"
)

var TestCaseRootList string

func init() {
	TestCaseRootList = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "list")
	//TestCaseRootList = `/Users/vej/go/src/github.com/1set/gut/local/test_resource/yos/list`
}

func verifyTestResult(t *testing.T, name string, expected []string, actual []*FilePathInfo, err error) {
	if err != nil {
		t.Errorf("%s() got error = %v, wantErr %v", name, err, false)
		return
	}
	if len(actual) != len(expected) {
		t.Errorf("%s() got length = %v, want = %v", name, len(actual), len(expected))
		return
	}

	for idx, item := range actual {
		suffix := strings.Replace(expected[idx], `/`, string(os.PathSeparator), -1)
		if !strings.HasSuffix(item.Path, suffix) {
			t.Errorf("%s() got #%d path = %q, want suffix = %q", name, idx, item.Path, suffix)
			return
		}
		fileName := item.Info.Name()
		if !strings.HasSuffix(suffix, fileName) {
			t.Errorf("%s() got #%d suffix = %q, want name = %q", name, idx, suffix, fileName)
			return
		}
	}
}

func TestListAll(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListAll(path); err == nil {
			t.Errorf("ListAll(%q) got no error, diff from expected", path)
			return
		}
	}
	actual, err := ListAll(TestCaseRootList)
	verifyTestResult(t, "ListAll", expectedResultMap["All"], actual, err)
}

func BenchmarkListAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListAll(TestCaseRootList)
	}
}

func TestListFile(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
	}
	actual, err := ListFile(TestCaseRootList)
	verifyTestResult(t, "ListFile", expectedResultMap["AllFiles"], actual, err)
}

func BenchmarkListFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListFile(TestCaseRootList)
	}
}

func TestListDir(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
	}
	actual, err := ListDir(TestCaseRootList)
	verifyTestResult(t, "ListDir", expectedResultMap["AllDirs"], actual, err)
}

func BenchmarkListDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListDir(TestCaseRootList)
	}
}

func TestListMatch(t *testing.T) {
	type args struct {
		root     string
		flag     int
		patterns []string
	}
	tests := []struct {
		name       string
		args       args
		wantSuffix []string
		wantErr    bool
	}{
		// TODO: fill these tests
		{"Empty root", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Root not exist", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"No Flag", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for file", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for dir", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for file & dir", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for no recursive", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for recursive", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for ToLower", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Flag for no ToLower", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"No pattern", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Broken pattern", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"No pattern", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Empty pattern", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern for *", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern for exact match", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern match none", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern with slash", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern with case-sensitive match", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
		{"Pattern with case-insensitive match", args{TestCaseRootList, 0, []string{}}, expectedResultMap["Empty"], false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ListMatch(tt.args.root, tt.args.flag, tt.args.patterns...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				verifyTestResult(t, "ListMatch", tt.wantSuffix, actual, err)
			}
		})
	}
}

func BenchmarkListMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListMatch(TestCaseRootList, ListRecursive|ListIncludeFile|ListIncludeDir, "*.txt", "deep*")
	}
}

var expectedResultMap = map[string][]string{
	"Empty": []string{},
	"All": []string{
		"yos/list",
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/broken_symlink.wtf",
		"yos/list/deep_folder",
		"yos/list/deep_folder/deep",
		"yos/list/deep_folder/deep/deeper",
		"yos/list/deep_folder/deep/deeper/deepest",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/empty_folder",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/folder_like_file.txt",
		"yos/list/nested_empty",
		"yos/list/nested_empty/empty1",
		"yos/list/nested_empty/empty1/empty2",
		"yos/list/nested_empty/empty1/empty2/empty3",
		"yos/list/nested_empty/empty1/empty2/empty3/empty4",
		"yos/list/nested_empty/empty1/empty2/empty3/empty4/empty5",
		"yos/list/no_ext_name_file",
		"yos/list/simple_folder",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space",
		"yos/list/white space/only one file",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"AllFiles": []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/broken_symlink.wtf",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/no_ext_name_file",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space/only one file",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"AllDirs": []string{
		"yos/list",
		"yos/list/deep_folder",
		"yos/list/deep_folder/deep",
		"yos/list/deep_folder/deep/deeper",
		"yos/list/deep_folder/deep/deeper/deepest",
		"yos/list/empty_folder",
		"yos/list/folder_like_file.txt",
		"yos/list/nested_empty",
		"yos/list/nested_empty/empty1",
		"yos/list/nested_empty/empty1/empty2",
		"yos/list/nested_empty/empty1/empty2/empty3",
		"yos/list/nested_empty/empty1/empty2/empty3/empty4",
		"yos/list/nested_empty/empty1/empty2/empty3/empty4/empty5",
		"yos/list/simple_folder",
		"yos/list/white space",
	},
}
