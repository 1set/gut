package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	CaseSameRoot        string
	CaseSameLinkRoot    string
	CaseSameFileMapSet1 map[string]string
	CaseSameFileMapSet2 map[string]string
)

func init() {
	CaseSameRoot = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "same")
	CaseSameLinkRoot = JoinPath(CaseSameRoot, "link")
	CaseSameFileMapSet1 = map[string]string{
		"EmptyDir":       JoinPath(CaseSameRoot, "set1", "empty-folder"),
		"EmptyFile":      JoinPath(CaseSameRoot, "set1", "empty-file.txt"),
		"SmallText":      JoinPath(CaseSameRoot, "set1", "small-text.txt"),
		"LargeText":      JoinPath(CaseSameRoot, "set1", "large-text.txt"),
		"PngImage":       JoinPath(CaseSameRoot, "set1", "image.png"),
		"SvgImage":       JoinPath(CaseSameRoot, "set1", "image.svg"),
		"BrokenSymlink":  JoinPath(CaseSameRoot, "set1", "broken_symlink.txt"),
		"NonePermission": JoinPath(CaseSameRoot, "set1", "none_perm.txt"),
	}
	CaseSameFileMapSet2 = map[string]string{
		"EmptyDir":       JoinPath(CaseSameRoot, "set2", "empty-folder"),
		"EmptyFile":      JoinPath(CaseSameRoot, "set2", "empty-file.txt"),
		"SmallText":      JoinPath(CaseSameRoot, "set2", "small-text.txt"),
		"SmallTextExe":   JoinPath(CaseSameRoot, "set2", "small-text.exe"),
		"SmallTextV2":    JoinPath(CaseSameRoot, "set2", "small-text-v2.txt"),
		"SmallTextV3":    JoinPath(CaseSameRoot, "set2", "small-text-v3.txt"),
		"LargeText":      JoinPath(CaseSameRoot, "set2", "large-text.txt"),
		"LargeTextV2":    JoinPath(CaseSameRoot, "set2", "large-text-v2.txt"),
		"PngImage":       JoinPath(CaseSameRoot, "set2", "image.png"),
		"SvgImage":       JoinPath(CaseSameRoot, "set2", "image.svg"),
		"BrokenSymlink":  JoinPath(CaseSameRoot, "set2", "broken_symlink.txt"),
		"NonePermission": JoinPath(CaseSameRoot, "set2", "none_perm.txt"),
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
		{"Path1 is empty", EmptyString, CaseSameFileMapSet2["SmallText"], false, true},
		{"Path2 is empty", CaseSameFileMapSet1["SmallText"], EmptyString, false, true},
		{"Path1 is not found", "__not_found_file__", CaseSameFileMapSet2["SmallText"], false, true},
		{"Path2 is not found", CaseSameFileMapSet1["SmallText"], "__not_found_file__", false, true},
		{"Path1 got permission denied", CaseSameFileMapSet1["NonePermission"], CaseSameFileMapSet2["SmallText"], false, true},
		{"Path2 got permission denied", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["NonePermission"], false, true},
		{"Path1 is a directory", CaseSameFileMapSet1["EmptyDir"], CaseSameFileMapSet2["SmallText"], false, true},
		{"Path2 is a directory", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["EmptyDir"], false, true},
		{"Path1 is a broken symlink", CaseSameFileMapSet1["BrokenSymlink"], CaseSameFileMapSet2["SmallText"], false, true},
		{"Path2 is a broken symlink", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["BrokenSymlink"], false, true},
		{"Path1 and path2 are exactly the same file", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet1["SmallText"], true, false},
		{"Path1 and path2 are actually the same file", CaseSameFileMapSet1["SmallText"], joinPathNoClean(CaseSameRoot, "set1", "..", "set1", "small-text.txt"), true, false},
		{"Path1 and path2 are files with same content", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["SmallText"], true, false},
		{"Path1 and path2 are files with same content and different permissions", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["SmallTextExe"], true, false},
		{"Path1 and path2 are empty files", CaseSameFileMapSet1["EmptyFile"], CaseSameFileMapSet2["EmptyFile"], true, false},
		{"Path1 and path2 are different files (whitespace)", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["SmallTextV2"], false, false},
		{"Path1 and path2 are different files (newline)", CaseSameFileMapSet1["SmallText"], CaseSameFileMapSet2["SmallTextV3"], false, false},
		{"Path1 and path2 are different files with same size", CaseSameFileMapSet1["LargeText"], CaseSameFileMapSet2["LargeTextV2"], false, false},
		{"Path1 and path2 are symlinks to the same file", JoinPath(CaseSameLinkRoot, "link_content1.txt"), JoinPath(CaseSameLinkRoot, "link2_content1.txt"), true, false},
		{"Path1 and path2 are symlinks to files with same content", JoinPath(CaseSameLinkRoot, "link_content1.txt"), JoinPath(CaseSameLinkRoot, "link_content2.txt"), true, false},
		{"Path1 is a symlink to a directory", JoinPath(CaseSameLinkRoot, "link_folder"), CaseSameFileMapSet2["SmallText"], false, true},
		{"Path1 is a symlink to a file and path2 is the file", JoinPath(CaseSameLinkRoot, "link_content1.txt"), JoinPath(CaseSameLinkRoot, "content1.txt"), true, false},
		{"Path1 is a symlink to a file and path2 is a file with same content", JoinPath(CaseSameLinkRoot, "link_content1.txt"), JoinPath(CaseSameLinkRoot, "content2.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a file", JoinPath(CaseSameLinkRoot, "link_link_content1.txt"), JoinPath(CaseSameLinkRoot, "link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a directory", JoinPath(CaseSameLinkRoot, "link_link_folder"), JoinPath(CaseSameLinkRoot, "link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to path1", JoinPath(CaseSameLinkRoot, "circle_link1"), JoinPath(CaseSameLinkRoot, "circle_link2"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to itself", JoinPath(CaseSameLinkRoot, "link_self_link"), JoinPath(CaseSameLinkRoot, "self_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink which is broken", JoinPath(CaseSameLinkRoot, "link_broken_link"), JoinPath(CaseSameLinkRoot, "broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink which is broken", JoinPath(CaseSameLinkRoot, "link_link_broken_link"), JoinPath(CaseSameLinkRoot, "link_broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a file", JoinPath(CaseSameLinkRoot, "link_link_link_content1.txt"), JoinPath(CaseSameLinkRoot, "link_link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a directory", JoinPath(CaseSameLinkRoot, "link_link_link_folder"), JoinPath(CaseSameLinkRoot, "link_link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to path1", JoinPath(CaseSameLinkRoot, "triple_link1"), JoinPath(CaseSameLinkRoot, "triple_link2"), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "permission") && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}

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
	b.Run("SameFile", func(b *testing.B) {
		path := CaseSameFileMapSet1["LargeText"]
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = SameContent(path, path)
		}
	})

	for name, path1 := range CaseSameFileMapSet1 {
		path2 := CaseSameFileMapSet2[name]
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = SameContent(path1, path2)
			}
		})
	}
}
