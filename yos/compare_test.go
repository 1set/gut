package yos

import (
	"os"
	"testing"
)

var (
	TestCaseRootSame string
	TestFileMapSet1  map[string]string
	TestFileMapSet2  map[string]string
)

func init() {
	TestCaseRootSame = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "same")
	TestFileMapSet1 = map[string]string{
		"EmptyDir":  JoinPath(TestCaseRootSame, "set1", "empty-folder"),
		"EmptyFile": JoinPath(TestCaseRootSame, "set1", "empty-file.txt"),
		"SmallText": JoinPath(TestCaseRootSame, "set1", "small-text.txt"),
		"LargeText": JoinPath(TestCaseRootSame, "set1", "large-text.txt"),
		"PngImage":  JoinPath(TestCaseRootSame, "set1", "image.png"),
		"SvgImage":  JoinPath(TestCaseRootSame, "set1", "image.svg"),
	}
	TestFileMapSet2 = map[string]string{
		"EmptyDir":    JoinPath(TestCaseRootSame, "set2", "empty-folder"),
		"EmptyFile":   JoinPath(TestCaseRootSame, "set2", "empty-file.txt"),
		"SmallText":   JoinPath(TestCaseRootSame, "set2", "small-text.txt"),
		"SmallTextV2": JoinPath(TestCaseRootSame, "set2", "small-text-v2.txt"),
		"SmallTextV3": JoinPath(TestCaseRootSame, "set2", "small-text-v3.txt"),
		"LargeText":   JoinPath(TestCaseRootSame, "set2", "large-text.txt"),
		"LargeTextV2": JoinPath(TestCaseRootSame, "set2", "large-text-v2.txt"),
		"PngImage":    JoinPath(TestCaseRootSame, "set2", "image.png"),
		"SvgImage":    JoinPath(TestCaseRootSame, "set2", "image.svg"),
	}
}

func TestSameContent(t *testing.T) {
	tests := []struct {
		name     string
		path1    string
		path2    string
		wantSame bool
		wantErr  bool
	}{
		{"Path1 is empty", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path2 is empty", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 is not found", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path2 is not found", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 is a directory", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path2 is a directory", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 is a broken symlink", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path2 is a broken symlink", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are exactly the same file", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are actually the same file", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are files with same content", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are files with same content and different permissions", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are empty files", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are different files", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
		{"Path1 and path2 are different files with same size", TestFileMapSet1[""], TestFileMapSet2[""], false, true},
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
