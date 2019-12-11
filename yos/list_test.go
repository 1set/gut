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

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListAll(path); err == nil {
			t.Errorf("ListAll(%q) got no error, diff from expected", path)
			return
		}
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

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
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

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
	}
	actual, err := ListDir(TestCaseRootList)
	verifyTestResult(t, "ListDir", expected, actual, err)
}

func BenchmarkListDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListDir(TestCaseRootList)
	}
}

func TestListMatchAll(t *testing.T) {
	expectedAll := []string{
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
	expectedTxt := []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/folder_like_file.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space.txt",
	}
	expectedMd := []string{
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	}
	expectedTxtMd := append(expectedTxt, expectedMd...)
	emptyStringList := []string{}

	type args struct {
		root     string
		patterns []string
	}
	tests := []struct {
		name       string
		args       args
		wantSuffix []string
		wantErr    bool
	}{
		{"no pattern", args{TestCaseRootList, emptyStringList}, emptyStringList, false},
		{"empty pattern", args{TestCaseRootList, []string{""}}, emptyStringList, false},
		{"broken pattern", args{TestCaseRootList, []string{"*[1-b"}}, emptyStringList, true},
		{"pattern for none", args{TestCaseRootList, []string{"z*"}}, emptyStringList, false},
		{"pattern for all", args{TestCaseRootList, []string{"*"}}, expectedAll, false},
		{"pattern case-sensitive", args{TestCaseRootList, []string{"*.TXT"}}, emptyStringList, false},
		{"single pattern txt", args{TestCaseRootList, []string{"*.txt"}}, expectedTxt, false},
		{"single pattern md", args{TestCaseRootList, []string{"*.md"}}, expectedMd, false},
		{"patterns for txt & md", args{TestCaseRootList, []string{"*.txt", "*.md"}}, expectedTxtMd, false},
		{"patterns for md & txt", args{TestCaseRootList, []string{"*.md", "*.txt"}}, expectedTxtMd, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ListMatchAll(tt.args.root, tt.args.patterns...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatchAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				verifyTestResult(t, "ListMatchAll", tt.wantSuffix, actual, err)
			}
		})
	}
}

func BenchmarkListMatchAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListMatchAll(TestCaseRootList, "*.md", "*.txt")
	}
}

func TestListMatchFile(t *testing.T) {
	expectedAll := []string{
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
	expectedTxt := []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space.txt",
	}
	expectedMd := []string{
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	}
	expectedTxtMd := append(expectedTxt, expectedMd...)
	emptyStringList := []string{}

	type args struct {
		root     string
		patterns []string
	}
	tests := []struct {
		name       string
		args       args
		wantSuffix []string
		wantErr    bool
	}{
		{"no pattern", args{TestCaseRootList, emptyStringList}, emptyStringList, false},
		{"empty pattern", args{TestCaseRootList, []string{""}}, emptyStringList, false},
		{"broken pattern", args{TestCaseRootList, []string{"*[1-b"}}, emptyStringList, true},
		{"pattern for none", args{TestCaseRootList, []string{"z*"}}, emptyStringList, false},
		{"pattern for all", args{TestCaseRootList, []string{"*"}}, expectedAll, false},
		{"pattern case-sensitive", args{TestCaseRootList, []string{"*.TXT"}}, emptyStringList, false},
		{"single pattern txt", args{TestCaseRootList, []string{"*.txt"}}, expectedTxt, false},
		{"single pattern md", args{TestCaseRootList, []string{"*.md"}}, expectedMd, false},
		{"patterns for txt & md", args{TestCaseRootList, []string{"*.txt", "*.md"}}, expectedTxtMd, false},
		{"patterns for md & txt", args{TestCaseRootList, []string{"*.md", "*.txt"}}, expectedTxtMd, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ListMatchFile(tt.args.root, tt.args.patterns...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatchFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				verifyTestResult(t, "ListMatchFile", tt.wantSuffix, actual, err)
			}
		})
	}
}

func BenchmarkListMatchFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListMatchFile(TestCaseRootList, "*.md", "*.txt")
	}
}

func TestListMatchDir(t *testing.T) {
	expectedAll := []string{
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
	expectedTxt := []string{
		"yos/list/folder_like_file.txt",
	}
	expectedDeep := []string{
		"yos/list/deep_folder",
		"yos/list/deep_folder/deep",
		"yos/list/deep_folder/deep/deeper",
		"yos/list/deep_folder/deep/deeper/deepest",
	}
	expectedDeepTxt := append(expectedDeep, expectedTxt...)
	emptyStringList := []string{}

	type args struct {
		root     string
		patterns []string
	}
	tests := []struct {
		name       string
		args       args
		wantSuffix []string
		wantErr    bool
	}{
		{"no pattern", args{TestCaseRootList, emptyStringList}, emptyStringList, false},
		{"empty pattern", args{TestCaseRootList, []string{""}}, emptyStringList, false},
		{"broken pattern", args{TestCaseRootList, []string{"*[1-b"}}, emptyStringList, true},
		{"pattern for none", args{TestCaseRootList, []string{"z*"}}, emptyStringList, false},
		{"pattern for all", args{TestCaseRootList, []string{"*"}}, expectedAll, false},
		{"pattern case-sensitive", args{TestCaseRootList, []string{"*.TXT"}}, emptyStringList, false},
		{"single pattern txt", args{TestCaseRootList, []string{"*.txt"}}, expectedTxt, false},
		{"single pattern deep", args{TestCaseRootList, []string{"deep*"}}, expectedDeep, false},
		{"patterns for txt & deep", args{TestCaseRootList, []string{"*.txt", "deep*"}}, expectedDeepTxt, false},
		{"patterns for deep & txt", args{TestCaseRootList, []string{"deep*", "*.txt"}}, expectedDeepTxt, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ListMatchDir(tt.args.root, tt.args.patterns...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatchDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				verifyTestResult(t, "ListMatchDir", tt.wantSuffix, actual, err)
			}
		})
	}
}

func BenchmarkListMatchDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListMatchDir(TestCaseRootList, "*.txt", "deep*")
	}
}

func TestListMatch(t *testing.T) {
	expectedAll := []string{
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
	expectedLowerFileTxt := []string{
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
	}
	expectedAllFileTxt := []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
	}
	expectedTxt := []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/folder_like_file.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space.txt",
	}
	expectedMd := []string{
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	}
	expectedTxtMd := append(expectedTxt, expectedMd...)
	emptyStringList := []string{}

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
		{"no pattern", args{TestCaseRootList, 0, emptyStringList}, emptyStringList, false},
		{"empty pattern", args{TestCaseRootList, 0, []string{""}}, emptyStringList, false},
		{"broken pattern", args{TestCaseRootList, ListIncludeDir, []string{"*[1-b"}}, emptyStringList, true},
		{"pattern for none", args{TestCaseRootList, ListRecursive | ListIncludeDir | ListIncludeFile, []string{"z*"}}, emptyStringList, false},
		{"pattern for all", args{TestCaseRootList, ListRecursive | ListIncludeDir | ListIncludeFile, []string{"*"}}, expectedAll, false},
		{"pattern case-sensitive", args{TestCaseRootList, ListRecursive | ListIncludeFile, []string{"file*"}}, expectedLowerFileTxt, false},
		{"pattern case-insensitive", args{TestCaseRootList, ListRecursive | ListIncludeFile | ListToLower, []string{"file*"}}, expectedAllFileTxt, false},
		{"single pattern txt", args{TestCaseRootList, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt"}}, expectedTxt, false},
		{"single pattern md", args{TestCaseRootList, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.md"}}, expectedMd, false},
		{"patterns for txt & md", args{TestCaseRootList, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt", "*.md"}}, expectedTxtMd, false},
		{"patterns for md & txt", args{TestCaseRootList, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.md", "*.txt"}}, expectedTxtMd, false},
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
