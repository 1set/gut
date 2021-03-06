package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	resourceListRoot          string
	resourceListFileInRoot    string
	resourceListSymlinkToRoot string
)

func init() {
	resourceListRoot = JoinPath(testResourceRoot, "yos", "list")
	resourceListFileInRoot = JoinPath(resourceListRoot, "no_ext_name_file")
	resourceListSymlinkToRoot = JoinPath(testResourceRoot, "yos", "link_list_dir")
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
	for _, path := range []string{"", "  ", "__not_found_folder__", resourceListFileInRoot} {
		if _, err := ListAll(path); err == nil {
			t.Errorf("ListAll(%q) got no error, diff from expected", path)
		} else {
			expectedErrorCheck(t, err)
		}
	}

	actual, err := ListAll(resourceListRoot)
	verifyTestResult(t, "ListAll", expectedResultMap["All"], actual, err)

	actual, err = ListAll(resourceListSymlinkToRoot)
	verifyTestResult(t, "ListAll(Symlink)", expectedResultMap["All"], actual, err)
}

func BenchmarkListAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListAll(resourceListRoot)
	}
}

func TestListFile(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__", resourceListFileInRoot} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
		} else {
			expectedErrorCheck(t, err)
		}
	}

	actual, err := ListFile(resourceListRoot)
	verifyTestResult(t, "ListFile", expectedResultMap["AllFiles"], actual, err)

	actual, err = ListFile(resourceListSymlinkToRoot)
	verifyTestResult(t, "ListFile(Symlink)", expectedResultMap["AllFiles"], actual, err)
}

func BenchmarkListFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListFile(resourceListRoot)
	}
}

func TestListSymlink(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__", resourceListFileInRoot} {
		if _, err := ListSymlink(path); err == nil {
			t.Errorf("ListSymlink(%q) got no error, diff from expected", path)
		} else {
			expectedErrorCheck(t, err)
		}
	}

	actual, err := ListSymlink(resourceListRoot)
	verifyTestResult(t, "ListSymlink", expectedResultMap["AllSymlinks"], actual, err)

	actual, err = ListSymlink(resourceListSymlinkToRoot)
	verifyTestResult(t, "ListSymlink(Symlink)", expectedResultMap["AllSymlinks"], actual, err)
}

func BenchmarkListSymlink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListSymlink(resourceListRoot)
	}
}

func TestListDir(t *testing.T) {
	for _, path := range []string{"", "  ", "__not_found_folder__", resourceListFileInRoot} {
		if _, err := ListFile(path); err == nil {
			t.Errorf("ListFile(%q) got no error, diff from expected", path)
		} else {
			expectedErrorCheck(t, err)
		}
	}

	actual, err := ListDir(resourceListRoot)
	verifyTestResult(t, "ListDir", expectedResultMap["AllDirs"], actual, err)

	actual, err = ListDir(resourceListSymlinkToRoot)
	verifyTestResult(t, "ListDir(Symlink)", expectedResultMap["AllDirs"], actual, err)
}

func BenchmarkListDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListDir(resourceListRoot)
	}
}

func TestListMatch(t *testing.T) {
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
		{"Root is a file", args{resourceListFileInRoot, ListIncludeFile, allEntriesPattern}, expectedResultMap["Empty"], true},
		{"Root is a symlink to directory", args{resourceListSymlinkToRoot, ListIncludeFile, allEntriesPattern}, expectedResultMap["RootFiles"], false},
		{"No Flag", args{resourceListRoot, 0, allEntriesPattern}, expectedResultMap["Empty"], false},
		{"Flag for file", args{resourceListRoot, ListIncludeFile, allEntriesPattern}, expectedResultMap["RootFiles"], false},
		{"Flag for dir", args{resourceListRoot, ListIncludeDir, allEntriesPattern}, expectedResultMap["RootDirs"], false},
		{"Flag for link", args{resourceListRoot, ListIncludeSymlink, allEntriesPattern}, expectedResultMap["RootSymlinks"], false},
		{"Flag for all", args{resourceListRoot, ListIncludeAll, allEntriesPattern}, expectedResultMap["RootAll"], false},
		{"Flag for recursive & file", args{resourceListRoot, ListRecursive | ListIncludeFile, allEntriesPattern}, expectedResultMap["AllFiles"], false},
		{"Flag for recursive & link", args{resourceListRoot, ListRecursive | ListIncludeSymlink, allEntriesPattern}, expectedResultMap["AllSymlinks"], false},
		{"Flag for recursive & dir", args{resourceListRoot, ListRecursive | ListIncludeDir, allEntriesPattern}, expectedResultMap["AllDirs"], false},
		{"Flag with ToLower", args{resourceListRoot, ListIncludeFile | ListToLower, []string{"file*"}}, expectedResultMap["AllFile*Insensitive"], false},
		{"Flag without ToLower", args{resourceListRoot, ListIncludeFile, []string{"file*"}}, expectedResultMap["AllFile*Sensitive"], false},
		{"No pattern", args{resourceListRoot, ListIncludeFile, expectedResultMap["Empty"]}, expectedResultMap["Empty"], false},
		{"Broken pattern", args{resourceListRoot, ListIncludeFile, []string{"*[1-"}}, expectedResultMap["Empty"], true},
		{"Broken regexp pattern", args{resourceListRoot, ListIncludeFile | ListUseRegExp, []string{"[a"}}, expectedResultMap["Empty"], true},
		{"Empty pattern", args{resourceListRoot, ListIncludeFile, []string{""}}, expectedResultMap["Empty"], false},
		{"Pattern for exact match", args{resourceListRoot, ListRecursive | ListIncludeFile, []string{"file1.txt"}}, expectedResultMap["All file1.txt"], false},
		{"Pattern for exclude", args{resourceListRoot, ListRecursive | ListIncludeFile, []string{"[^.]*"}}, expectedResultMap["AllFiles"], false},
		{"Pattern match none", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"*.pdf"}}, expectedResultMap["Empty"], false},
		{"Pattern match txt", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"*.txt"}}, expectedResultMap["All *.txt"], false},
		{"Pattern with slash", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"/*.txt"}}, expectedResultMap["Empty"], false},
		{"Pattern with case-sensitive match", args{resourceListRoot, ListIncludeFile, []string{"File*"}}, expectedResultMap["OnlyFile*"], false},
		{"Pattern with case-insensitive non-match", args{resourceListRoot, ListIncludeFile | ListToLower, []string{"File*"}}, expectedResultMap["Empty"], false},
		{"Regexp pattern: ^[fF]ile(.*).txt", args{resourceListRoot, ListRecursive | ListIncludeAll | ListUseRegExp, []string{"^[fF]ile(.*).txt"}}, expectedResultMap["^[fF]ile(.*).txt"], false},
		{"Regexp pattern: ^file(.*).txt", args{resourceListRoot, ListRecursive | ListIncludeAll | ListUseRegExp, []string{"^file(.*).txt"}}, expectedResultMap["^file(.*).txt"], false},
		{"Duplicate patterns", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"*.txt", "*.txt", "*.txt"}}, expectedResultMap["All *.txt"], false},
		{"Multiple matched patterns", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"*.txt", "*.md"}}, expectedResultMap["All *.txt *.md"], false},
		{"Multiple matched regexp patterns", args{resourceListRoot, ListRecursive | ListIncludeAll | ListUseRegExp, []string{".txt$", ".md$"}}, expectedResultMap["All *.txt *.md"], false},
		{"Combine of match and non-match patterns", args{resourceListRoot, ListRecursive | ListIncludeAll, []string{"*.txt", "*.pdf", "*.jpg"}}, expectedResultMap["All *.txt"], false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ListMatch(tt.args.root, tt.args.flag, tt.args.patterns...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatch() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				verifyTestResult(t, "ListMatch", tt.wantSuffix, actual, err)
			} else {
				expectedErrorCheck(t, err)
			}
		})
	}
}

func BenchmarkListMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListMatch(resourceListRoot, ListRecursive|ListIncludeFile|ListIncludeDir, "*.txt", "deep*")
	}
}

var expectedResultMap = map[string][]string{
	"Empty": {},
	"All": {
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
	"AllFiles": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/no_ext_name_file",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/white space/only one file",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"AllSymlinks": {
		"yos/list/broken_symlink.wtf",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
	},
	"AllDirs": {
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
	"RootFiles": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/no_ext_name_file",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"RootSymlinks": {
		"yos/list/broken_symlink.wtf",
		"yos/list/symlink_to_dir",
		"yos/list/symlink_to_file.txt",
	},
	"RootDirs": {
		"yos/list/deep_folder",
		"yos/list/empty_folder",
		"yos/list/folder_like_file.txt",
		"yos/list/nested_empty",
		"yos/list/simple_folder",
		"yos/list/white space",
	},
	"RootAll": {
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
	"AllFile*Insensitive": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
	},
	"^[fF]ile(.*).txt": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
	},
	"^file(.*).txt": {
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
	},
	"AllFile*Sensitive": {
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
	},
	"All file1.txt": {
		"yos/list/file1.txt",
		"yos/list/simple_folder/file1.txt",
	},
	"All *.txt": {
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
	"All *.txt *.md": {
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
	"OnlyFile*": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
	},
	"SortByModTime": {
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/white space.txt",
		"yos/list/white space/only one file",
		"yos/list/no_ext_name_file",
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
	},
	"SortByName": {
		"yos/list/File0.txt",
		"yos/list/File4.txt",
		"yos/list/file1.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/file2.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/no_ext_name_file",
		"yos/list/white space/only one file",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/white space.txt",
		"yos/list/测试文件.md",
		"yos/list/🤙🏝️.md",
	},
	"SortBySize": {
		"yos/list/white space/only one file",
		"yos/list/file1.txt",
		"yos/list/file2.txt",
		"yos/list/file3.txt",
		"yos/list/simple_folder/file1.txt",
		"yos/list/simple_folder/file2.txt",
		"yos/list/simple_folder/file3.txt",
		"yos/list/File4.txt",
		"yos/list/File0.txt",
		"yos/list/deep_folder/deep/deeper/deepest/text_file.txt",
		"yos/list/no_ext_name_file",
		"yos/list/🤙🏝️.md",
		"yos/list/测试文件.md",
		"yos/list/white space.txt",
	},
}
