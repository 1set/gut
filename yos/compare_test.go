package yos

import (
	"os"
	"strings"
	"testing"
)

var (
	resourceSameFileRoot     string
	resourceSameFileLinkRoot string
	resourceSameFileMapSet1  map[string]string
	resourceSameFileMapSet2  map[string]string
)

func init() {
	resourceSameFileRoot = JoinPath(os.Getenv("TESTRSSDIR"), "yos", "same_file")
	resourceSameFileLinkRoot = JoinPath(resourceSameFileRoot, "link")
	resourceSameFileMapSet1 = map[string]string{
		"EmptyDir":       JoinPath(resourceSameFileRoot, "set1", "empty-folder"),
		"EmptyFile":      JoinPath(resourceSameFileRoot, "set1", "empty-file.txt"),
		"SmallText":      JoinPath(resourceSameFileRoot, "set1", "small-text.txt"),
		"LargeText":      JoinPath(resourceSameFileRoot, "set1", "large-text.txt"),
		"PngImage":       JoinPath(resourceSameFileRoot, "set1", "image.png"),
		"SvgImage":       JoinPath(resourceSameFileRoot, "set1", "image.svg"),
		"BrokenSymlink":  JoinPath(resourceSameFileRoot, "set1", "broken_symlink.txt"),
		"NonePermission": JoinPath(resourceSameFileRoot, "set1", "none_perm.txt"),
	}
	resourceSameFileMapSet2 = map[string]string{
		"EmptyDir":       JoinPath(resourceSameFileRoot, "set2", "empty-folder"),
		"EmptyFile":      JoinPath(resourceSameFileRoot, "set2", "empty-file.txt"),
		"SmallText":      JoinPath(resourceSameFileRoot, "set2", "small-text.txt"),
		"SmallTextExe":   JoinPath(resourceSameFileRoot, "set2", "small-text.exe"),
		"SmallTextV2":    JoinPath(resourceSameFileRoot, "set2", "small-text-v2.txt"),
		"SmallTextV3":    JoinPath(resourceSameFileRoot, "set2", "small-text-v3.txt"),
		"LargeText":      JoinPath(resourceSameFileRoot, "set2", "large-text.txt"),
		"LargeTextV2":    JoinPath(resourceSameFileRoot, "set2", "large-text-v2.txt"),
		"PngImage":       JoinPath(resourceSameFileRoot, "set2", "image.png"),
		"SvgImage":       JoinPath(resourceSameFileRoot, "set2", "image.svg"),
		"BrokenSymlink":  JoinPath(resourceSameFileRoot, "set2", "broken_symlink.txt"),
		"NonePermission": JoinPath(resourceSameFileRoot, "set2", "none_perm.txt"),
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
		{"Path1 and path2 are actually the same file", resourceSameFileMapSet1["SmallText"], joinPathNoClean(resourceSameFileRoot, "set1", "..", "set1", "small-text.txt"), true, false},
		{"Path1 and path2 are files with same content", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallText"], true, false},
		{"Path1 and path2 are files with same content and different permissions", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextExe"], true, false},
		{"Path1 and path2 are empty files", resourceSameFileMapSet1["EmptyFile"], resourceSameFileMapSet2["EmptyFile"], true, false},
		{"Path1 and path2 are different files (whitespace)", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextV2"], false, false},
		{"Path1 and path2 are different files (newline)", resourceSameFileMapSet1["SmallText"], resourceSameFileMapSet2["SmallTextV3"], false, false},
		{"Path1 and path2 are different files with same size", resourceSameFileMapSet1["LargeText"], resourceSameFileMapSet2["LargeTextV2"], false, false},
		{"Path1 and path2 are symlinks to the same file", JoinPath(resourceSameFileLinkRoot, "link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "link2_content1.txt"), true, false},
		{"Path1 and path2 are symlinks to files with same content", JoinPath(resourceSameFileLinkRoot, "link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "link_content2.txt"), true, false},
		{"Path1 is a symlink to a directory", JoinPath(resourceSameFileLinkRoot, "link_folder"), resourceSameFileMapSet2["SmallText"], false, true},
		{"Path1 is a symlink to a file and path2 is the file", JoinPath(resourceSameFileLinkRoot, "link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "content1.txt"), true, false},
		{"Path1 is a symlink to a file and path2 is a file with same content", JoinPath(resourceSameFileLinkRoot, "link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "content2.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a file", JoinPath(resourceSameFileLinkRoot, "link_link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to a directory", JoinPath(resourceSameFileLinkRoot, "link_link_folder"), JoinPath(resourceSameFileLinkRoot, "link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to path1", JoinPath(resourceSameFileLinkRoot, "circle_link1"), JoinPath(resourceSameFileLinkRoot, "circle_link2"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to itself", JoinPath(resourceSameFileLinkRoot, "link_self_link"), JoinPath(resourceSameFileLinkRoot, "self_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink which is broken", JoinPath(resourceSameFileLinkRoot, "link_broken_link"), JoinPath(resourceSameFileLinkRoot, "broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink which is broken", JoinPath(resourceSameFileLinkRoot, "link_link_broken_link"), JoinPath(resourceSameFileLinkRoot, "link_broken_link"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a file", JoinPath(resourceSameFileLinkRoot, "link_link_link_content1.txt"), JoinPath(resourceSameFileLinkRoot, "link_link_content1.txt"), true, false},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to a directory", JoinPath(resourceSameFileLinkRoot, "link_link_link_folder"), JoinPath(resourceSameFileLinkRoot, "link_link_folder"), false, true},
		{"Path1 is a symlink to a symlink and path2 is the symlink to another symlink to path1", JoinPath(resourceSameFileLinkRoot, "triple_link1"), JoinPath(resourceSameFileLinkRoot, "triple_link2"), false, true},
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
