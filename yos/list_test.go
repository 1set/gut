package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	CaseListRoot     string
	CaseListFileRoot string
)

func init() {
	CaseListRoot = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "list")
	CaseListFileRoot = JoinPath(CaseListRoot, "no_ext_name_file")
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
	t.Parallel()

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListAll(path); err == nil {
			t.Errorf("ListAll(%q) got no error, diff from expected", path)
			return
		}
	}

	actual, err := ListAll(CaseListFileRoot)
	verifyTestResult(t, "ListAll", expectedResultMap["Empty"], actual, err)

	actual, err = ListAll(CaseListRoot)
	verifyTestResult(t, "ListAll", expectedResultMap["All"], actual, err)
}

func BenchmarkListAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListAll(CaseListRoot)
	}
}

func TestListFile(t *testing.T) {
	t.Parallel()

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
	}

	actual, err := ListFile(CaseListFileRoot)
	verifyTestResult(t, "ListFile", expectedResultMap["Empty"], actual, err)

	actual, err = ListFile(CaseListRoot)
	verifyTestResult(t, "ListFile", expectedResultMap["AllFiles"], actual, err)
}

func BenchmarkListFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListFile(CaseListRoot)
	}
}

func TestListDir(t *testing.T) {
	t.Parallel()

	for _, path := range []string{"", "  ", "__not_found_folder__"} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
			return
		}
	}

	actual, err := ListDir(CaseListFileRoot)
	verifyTestResult(t, "ListDir", expectedResultMap["Empty"], actual, err)

	actual, err = ListDir(CaseListRoot)
	verifyTestResult(t, "ListDir", expectedResultMap["AllDirs"], actual, err)
}

func BenchmarkListDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListDir(CaseListRoot)
	}
}

func TestListMatch(t *testing.T) {
	t.Parallel()

	allEntriesPattern := []string{"*"}
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
		{"Empty root path", args{"", ListIncludeFile, allEntriesPattern}, expectedResultMap["Empty"], true},
		{"Root not exist", args{"__not_found_folder__", ListIncludeFile, allEntriesPattern}, expectedResultMap["Empty"], true},
		{"Root is a file", args{CaseListFileRoot, ListIncludeFile, allEntriesPattern}, expectedResultMap["Empty"], false},
		{"No Flag", args{CaseListRoot, 0, allEntriesPattern}, expectedResultMap["Empty"], false},
		{"Flag for file", args{CaseListRoot, ListIncludeFile, allEntriesPattern}, expectedResultMap["RootFiles"], false},
		{"Flag for dir", args{CaseListRoot, ListIncludeDir, allEntriesPattern}, expectedResultMap["RootDirs"], false},
		{"Flag for file & dir", args{CaseListRoot, ListIncludeFile | ListIncludeDir, allEntriesPattern}, expectedResultMap["RootAll"], false},
		{"Flag for recursive & file", args{CaseListRoot, ListRecursive | ListIncludeFile, allEntriesPattern}, expectedResultMap["AllFiles"], false},
		{"Flag for recursive & dir", args{CaseListRoot, ListRecursive | ListIncludeDir, allEntriesPattern}, expectedResultMap["AllDirs"], false},
		{"Flag with ToLower", args{CaseListRoot, ListIncludeFile | ListToLower, []string{"file*"}}, expectedResultMap["AllFile*Insensitive"], false},
		{"Flag without ToLower", args{CaseListRoot, ListIncludeFile, []string{"file*"}}, expectedResultMap["AllFile*Sensitive"], false},
		{"No pattern", args{CaseListRoot, ListIncludeFile, expectedResultMap["Empty"]}, expectedResultMap["Empty"], false},
		{"Broken pattern", args{CaseListRoot, ListIncludeFile, []string{"*[1-"}}, expectedResultMap["Empty"], true},
		{"Empty pattern", args{CaseListRoot, ListIncludeFile, []string{""}}, expectedResultMap["Empty"], false},
		{"Pattern for exact match", args{CaseListRoot, ListRecursive | ListIncludeFile, []string{"file1.txt"}}, expectedResultMap["All file1.txt"], false},
		{"Pattern for exclude", args{CaseListRoot, ListRecursive | ListIncludeFile, []string{"[^.]*"}}, expectedResultMap["AllFiles"], false},
		{"Pattern match none", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.pdf"}}, expectedResultMap["Empty"], false},
		{"Pattern match txt", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt"}}, expectedResultMap["All *.txt"], false},
		{"Pattern with slash", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"/*.txt"}}, expectedResultMap["Empty"], false},
		{"Pattern with case-sensitive match", args{CaseListRoot, ListIncludeFile, []string{"File*"}}, expectedResultMap["OnlyFile*"], false},
		{"Pattern with case-insensitive non-match", args{CaseListRoot, ListIncludeFile | ListToLower, []string{"File*"}}, expectedResultMap["Empty"], false},
		{"Duplicate patterns", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt", "*.txt", "*.txt"}}, expectedResultMap["All *.txt"], false},
		{"Multiple matched patterns", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt", "*.md"}}, expectedResultMap["All *.txt *.md"], false},
		{"Combine of match and non-match patterns", args{CaseListRoot, ListRecursive | ListIncludeFile | ListIncludeDir, []string{"*.txt", "*.pdf", "*.jpg"}}, expectedResultMap["All *.txt"], false},
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
		_, _ = ListMatch(CaseListRoot, ListRecursive|ListIncludeFile|ListIncludeDir, "*.txt", "deep*")
	}
}

var expectedResultMap = map[string][]string{
	"Empty": []string{},
	"All": []string{
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
	"RootFiles": []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/broken_symlink.wtf",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/no_ext_name_file",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"RootDirs": []string{
		"yos/list/deep_folder",
		"yos/list/empty_folder",
		"yos/list/folder_like_file.txt",
		"yos/list/nested_empty",
		"yos/list/simple_folder",
		"yos/list/white space",
	},
	"RootAll": []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/broken_symlink.wtf",
		"yos/list/deep_folder",
		"yos/list/empty_folder",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/folder_like_file.txt",
		"yos/list/nested_empty",
		"yos/list/no_ext_name_file",
		"yos/list/simple_folder",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
		"yos/list/white space",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"AllFile*Insensitive": []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
	},
	"AllFile*Sensitive": []string{
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
	},
	"All file1.txt": []string{
		"yos/list/file1.txt",
		"yos/list/simple_folder/file1.txt",
	},
	"All *.txt": []string{
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
	},
	"All *.txt *.md": []string{
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
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"OnlyFile*": []string{
		"yos/list/File0.txt",
		"yos/list/File4.txt",
	},
}
