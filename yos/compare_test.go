package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	resourceSameRoot        string
	resourceSameLinkRoot    string
	resourceSameFileMapSet1 map[string]string
	resourceSameFileMapSet2 map[string]string
)

func init() {
	resourceSameRoot = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "same")
	resourceSameLinkRoot = JoinPath(resourceSameRoot, "link")
	resourceSameFileMapSet1 = map[string]string{
		"EmptyDir":       JoinPath(resourceSameRoot, "set1", "empty-folder"),
		"EmptyFile":      JoinPath(resourceSameRoot, "set1", "empty-file.txt"),
		"SmallText":      JoinPath(resourceSameRoot, "set1", "small-text.txt"),
		"LargeText":      JoinPath(resourceSameRoot, "set1", "large-text.txt"),
		"PngImage":       JoinPath(resourceSameRoot, "set1", "image.png"),
		"SvgImage":       JoinPath(resourceSameRoot, "set1", "image.svg"),
		"BrokenSymlink":  JoinPath(resourceSameRoot, "set1", "broken_symlink.txt"),
		"NonePermission": JoinPath(resourceSameRoot, "set1", "none_perm.txt"),
	}
	resourceSameFileMapSet2 = map[string]string{
		"EmptyDir":       JoinPath(resourceSameRoot, "set2", "empty-folder"),
		"EmptyFile":      JoinPath(resourceSameRoot, "set2", "empty-file.txt"),
		"SmallText":      JoinPath(resourceSameRoot, "set2", "small-text.txt"),
		"SmallTextExe":   JoinPath(resourceSameRoot, "set2", "small-text.exe"),
		"SmallTextV2":    JoinPath(resourceSameRoot, "set2", "small-text-v2.txt"),
		"SmallTextV3":    JoinPath(resourceSameRoot, "set2", "small-text-v3.txt"),
		"LargeText":      JoinPath(resourceSameRoot, "set2", "large-text.txt"),
		"LargeTextV2":    JoinPath(resourceSameRoot, "set2", "large-text-v2.txt"),
		"PngImage":       JoinPath(resourceSameRoot, "set2", "image.png"),
		"SvgImage":       JoinPath(resourceSameRoot, "set2", "image.svg"),
		"BrokenSymlink":  JoinPath(resourceSameRoot, "set2", "broken_symlink.txt"),
		"NonePermission": JoinPath(resourceSameRoot, "set2", "none_perm.txt"),
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
		{"Path1 is empty", emptyStr, resourceSameFileMapSet2["SmallText"], false, true},
		{"Path2 is empty", resourceSameFileMapSet1["SmallText"], emptyStr, false, true},
		{"Path1 is not found", "__not_found_file__", resourceSameFileMapSet2["SmallText"], false, true},
		{"Path2 is not found", resourceSameFileMapSet1["SmallText"], "__not_found_file__", false, true},
		{"Path1 got permission denied", resourceSameFileMapSet1["NonePermission"], resourceSameFileMapSet2["SmallText"], false, true},
		{"Path2 got permission denied", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["NonePermission"], false, true},
		{"Path1 is a directory", resourceSameFileMapSet1["EmptyDir"], resourceSameFileMapSet2["SmallText"], false, true},
		{"Path2 is a directory", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["EmptyDir"], false, true},
		{"Path1 is a broken symlink", resourceSameFileMapSet1["BrokenSymlink"], resourceSameFileMapSet2["SmallText"], false, true},
		{"Path2 is a broken symlink", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["BrokenSymlink"], false, true},
		{"Path1 and path2 are exactly the same file", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet1["SmallText"], true, false},
		{"Path1 and path2 are actually the same file", resourceSameFileMapSet1["SmallText"], joinPathNoClean(resourceSameRoot, "set1", "..", "set1", "small-text.txt"), true, false},
		{"Path1 and path2 are files with same content", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallText"], true, false},
		{"Path1 and path2 are files with same content and different permissions", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextExe"], true, false},
		{"Path1 and path2 are empty files", resourceSameFileMapSet1["EmptyFile"], resourceSameFileMapSet2["EmptyFile"], true, false},
		{"Path1 and path2 are different files (whitespace)", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextV2"], false, false},
		{"Path1 and path2 are different files (newline)", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextV3"], false, false},
		{"Path1 and path2 are different files with same size", resourceSameFileMapSet1["LargeText"], resourceSameFileMapSet2["LargeTextV2"], false, false},
		{"Path1 and path2 are symlinks to the same file", JoinPath(resourceSameLinkRoot, "link_content1.txt"), JoinPath(resourceSameLinkRoot, "link2_content1.txt"), true, false},
		{"Path1 and path2 are symlinks to files with same content", JoinPath(resourceSameLinkRoot, "link_content1.txt"), JoinPath(resourceSameLinkRoot, "link_content2.txt"), true, false},
		{"Path1 is a symlink to a directory", JoinPath(resourceSameLinkRoot, "link_folder"), resourceSameFileMapSet2["SmallText"], false, true},
		{"Path1 is a symlink to a file and path2 is the file", JoinPath(resourceSameLinkRoot, "link_content1.txt"), JoinPath(resourceSameLinkRoot, "content1.txt"), true, false},
		{"Path1 is a symlink to a file and path2 is a file with same content", JoinPath(resourceSameLinkRoot, "link_content1.txt"), JoinPath(resourceSameLinkRoot, "content2.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a file", JoinPath(resourceSameLinkRoot, "link_link_content1.txt"), JoinPath(resourceSameLinkRoot, "link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a directory", JoinPath(resourceSameLinkRoot, "link_link_folder"), JoinPath(resourceSameLinkRoot, "link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to path1", JoinPath(resourceSameLinkRoot, "circle_link1"), JoinPath(resourceSameLinkRoot, "circle_link2"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to itself", JoinPath(resourceSameLinkRoot, "link_self_link"), JoinPath(resourceSameLinkRoot, "self_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink which is broken", JoinPath(resourceSameLinkRoot, "link_broken_link"), JoinPath(resourceSameLinkRoot, "broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink which is broken", JoinPath(resourceSameLinkRoot, "link_link_broken_link"), JoinPath(resourceSameLinkRoot, "link_broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a file", JoinPath(resourceSameLinkRoot, "link_link_link_content1.txt"), JoinPath(resourceSameLinkRoot, "link_link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a directory", JoinPath(resourceSameLinkRoot, "link_link_link_folder"), JoinPath(resourceSameLinkRoot, "link_link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to path1", JoinPath(resourceSameLinkRoot, "triple_link1"), JoinPath(resourceSameLinkRoot, "triple_link2"), false, true},
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
		path := resourceSameFileMapSet1["LargeText"]
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = SameContent(path, path)
		}
	})

	for name, path1 := range resourceSameFileMapSet1 {
		path2 := resourceSameFileMapSet2[name]
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = SameContent(path1, path2)
			}
		})
	}
}
