package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	TestCaseRootSame     string
	TestCaseRootSameLink string
	TestFileMapSet1      map[string]string
	TestFileMapSet2      map[string]string
)

func init() {
	TestCaseRootSame = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "same")
	TestCaseRootSameLink = JoinPath(TestCaseRootSame, "link")
	TestFileMapSet1 = map[string]string{
		"EmptyDir":      JoinPath(TestCaseRootSame, "set1", "empty-folder"),
		"EmptyFile":     JoinPath(TestCaseRootSame, "set1", "empty-file.txt"),
		"SmallText":     JoinPath(TestCaseRootSame, "set1", "small-text.txt"),
		"LargeText":     JoinPath(TestCaseRootSame, "set1", "large-text.txt"),
		"PngImage":      JoinPath(TestCaseRootSame, "set1", "image.png"),
		"SvgImage":      JoinPath(TestCaseRootSame, "set1", "image.svg"),
		"BrokenSymlink": JoinPath(TestCaseRootSame, "set1", "broken_symlink.txt"),
	}
	TestFileMapSet2 = map[string]string{
		"EmptyDir":      JoinPath(TestCaseRootSame, "set2", "empty-folder"),
		"EmptyFile":     JoinPath(TestCaseRootSame, "set2", "empty-file.txt"),
		"SmallText":     JoinPath(TestCaseRootSame, "set2", "small-text.txt"),
		"SmallTextExe":  JoinPath(TestCaseRootSame, "set2", "small-text.exe"),
		"SmallTextV2":   JoinPath(TestCaseRootSame, "set2", "small-text-v2.txt"),
		"SmallTextV3":   JoinPath(TestCaseRootSame, "set2", "small-text-v3.txt"),
		"LargeText":     JoinPath(TestCaseRootSame, "set2", "large-text.txt"),
		"LargeTextV2":   JoinPath(TestCaseRootSame, "set2", "large-text-v2.txt"),
		"PngImage":      JoinPath(TestCaseRootSame, "set2", "image.png"),
		"SvgImage":      JoinPath(TestCaseRootSame, "set2", "image.svg"),
		"BrokenSymlink": JoinPath(TestCaseRootSame, "set1", "broken_symlink.txt"),
	}
}

func joinPathNoClean(elem ...string) string {
	return strings.Join(elem, string(os.PathSeparator))
}

func TestSameContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		path1    string
		path2    string
		wantSame bool
		wantErr  bool
	}{
		{"Path1 is empty", EmptyString, TestFileMapSet2["SmallText"], false, true},
		{"Path2 is empty", TestFileMapSet1["SmallText"], EmptyString, false, true},
		{"Path1 is not found", "__not_found_file__", TestFileMapSet2["SmallText"], false, true},
		{"Path2 is not found", TestFileMapSet1["SmallText"], "__not_found_file__", false, true},
		{"Path1 is a directory", TestFileMapSet1["EmptyDir"], TestFileMapSet2["SmallText"], false, true},
		{"Path2 is a directory", TestFileMapSet1["SmallText"], TestFileMapSet2["EmptyDir"], false, true},
		{"Path1 is a broken symlink", TestFileMapSet1["BrokenSymlink"], TestFileMapSet2["SmallText"], false, true},
		{"Path2 is a broken symlink", TestFileMapSet1["SmallText"], TestFileMapSet2["BrokenSymlink"], false, true},
		{"Path1 and path2 are exactly the same file", TestFileMapSet1["SmallText"], TestFileMapSet1["SmallText"], true, false},
		{"Path1 and path2 are actually the same file", TestFileMapSet1["SmallText"], joinPathNoClean(TestCaseRootSame, "set1", "..", "set1", "small-text.txt"), true, false},
		{"Path1 and path2 are files with same content", TestFileMapSet1["SmallText"], TestFileMapSet2["SmallText"], true, false},
		{"Path1 and path2 are files with same content and different permissions", TestFileMapSet1["SmallText"], TestFileMapSet2["SmallTextExe"], true, false},
		{"Path1 and path2 are empty files", TestFileMapSet1["EmptyFile"], TestFileMapSet2["EmptyFile"], true, false},
		{"Path1 and path2 are different files (whitespace)", TestFileMapSet1["SmallText"], TestFileMapSet2["SmallTextV2"], false, false},
		{"Path1 and path2 are different files (newline)", TestFileMapSet1["SmallText"], TestFileMapSet2["SmallTextV3"], false, false},
		{"Path1 and path2 are different files with same size", TestFileMapSet1["LargeText"], TestFileMapSet2["LargeTextV2"], false, false},
		{"Path1 and path2 are symlinks to the same file", JoinPath(TestCaseRootSameLink, "link_content1.txt"), JoinPath(TestCaseRootSameLink, "link2_content1.txt"), true, false},
		{"Path1 and path2 are symlinks to files with same content", JoinPath(TestCaseRootSameLink, "link_content1.txt"), JoinPath(TestCaseRootSameLink, "link_content2.txt"), true, false},
		{"Path1 is a symlink to a directory", JoinPath(TestCaseRootSameLink, "link_folder"), TestFileMapSet2["SmallText"], false, true},
		{"Path1 is a symlink to a file and path2 is the file", JoinPath(TestCaseRootSameLink, "link_content1.txt"), JoinPath(TestCaseRootSameLink, "content1.txt"), true, false},
		{"Path1 is a symlink to a file and path2 is a file with same content", JoinPath(TestCaseRootSameLink, "link_content1.txt"), JoinPath(TestCaseRootSameLink, "content2.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a file", JoinPath(TestCaseRootSameLink, "link_link_content1.txt"), JoinPath(TestCaseRootSameLink, "link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a directory", JoinPath(TestCaseRootSameLink, "link_link_folder"), JoinPath(TestCaseRootSameLink, "link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to path1", JoinPath(TestCaseRootSameLink, "circle_link1"), JoinPath(TestCaseRootSameLink, "circle_link2"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to itself", JoinPath(TestCaseRootSameLink, "link_self_link"), JoinPath(TestCaseRootSameLink, "self_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink which is broken", JoinPath(TestCaseRootSameLink, "link_broken_link"), JoinPath(TestCaseRootSameLink, "broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink which is broken", JoinPath(TestCaseRootSameLink, "link_link_broken_link"), JoinPath(TestCaseRootSameLink, "link_broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a file", JoinPath(TestCaseRootSameLink, "link_link_link_content1.txt"), JoinPath(TestCaseRootSameLink, "link_link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a directory", JoinPath(TestCaseRootSameLink, "link_link_link_folder"), JoinPath(TestCaseRootSameLink, "link_link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to path1", JoinPath(TestCaseRootSameLink, "triple_link1"), JoinPath(TestCaseRootSameLink, "triple_link2"), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSame, err := SameContent(tt.path1, tt.path2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SameContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSame != tt.wantSame {
				t.Errorf("SameContent() gotSame = %v, want %v", gotSame, tt.wantSame)
			}
		})
	}
}

func BenchmarkSameContent(b *testing.B) {
	for name, path1 := range TestFileMapSet1 {
		path2 := TestFileMapSet2[name]
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = SameContent(path1, path2)
			}
		})
	}
}
