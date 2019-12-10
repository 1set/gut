package yos

import (
	"os"
	"strings"
	"testing"
)

var TestCaseRootList string

func init() {
	TestCaseRootList = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "list")
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
		suffix := expected[idx]
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
	expected := []string{
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
	}

	actual, err := ListAll(TestCaseRootList)
	verifyTestResult(t, "ListAll", expected, actual, err)
}

func BenchmarkListAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListAll(TestCaseRootList)
	}
}

func TestListFile(t *testing.T) {
	expected := []string{
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
	}

	actual, err := ListFile(TestCaseRootList)
	verifyTestResult(t, "ListFile", expected, actual, err)
}

func BenchmarkListFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListFile(TestCaseRootList)
	}
}

func TestListDir(t *testing.T) {
	expected := []string{
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
	}

	actual, err := ListDir(TestCaseRootList)
	verifyTestResult(t, "ListDir", expected, actual, err)
}

func BenchmarkListDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListDir(TestCaseRootList)
	}
}
