package yos

import (
	"os"
	"strings"
	"testing"

	"github.com/1set/gut/ystring"
)

var (
	resourceReadWriteDevice string
	resourceReadOnlyDevice  string
	resourceMoveFileRoot    string
	resourceMoveSymlinkRoot string
	resourceMoveDirRoot     string
)

func init() {
	resourceReadWriteDevice = os.Getenv("RAMDISK_WRITE")
	resourceReadOnlyDevice = os.Getenv("RAMDISK_READONLY")

	testResourceRoot := os.Getenv("TESTRSSDIR")
	resourceMoveFileRoot = JoinPath(testResourceRoot, "yos", "move_file")
	resourceMoveSymlinkRoot = JoinPath(testResourceRoot, "yos", "move_link")
	resourceMoveDirRoot = JoinPath(testResourceRoot, "yos", "move_dir")
}

func TestMoveFile(t *testing.T) {
	var (
		bkRoot         = JoinPath(resourceMoveFileRoot, "backup")
		srcRoot        = JoinPath(resourceMoveFileRoot, "source")
		destRoot       = JoinPath(resourceMoveFileRoot, "destination")
		writeDevice    = JoinPath(resourceReadWriteDevice, "move_file")
		readOnlyDevice = JoinPath(resourceReadOnlyDevice, "move_file")
	)

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		backupPath string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(srcRoot, "missing-text.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to file", JoinPath(srcRoot, "link.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to directory", JoinPath(srcRoot, "link-dir"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a circular symlink", JoinPath(srcRoot, "link-circular"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a broken symlink", JoinPath(srcRoot, "link-broken"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same file", JoinPath(srcRoot, "text.txt"), JoinPath(srcRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same symlink", JoinPath(srcRoot, "link.txt"), JoinPath(srcRoot, "link.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same directory", JoinPath(srcRoot, "empty-dir"), JoinPath(srcRoot, "empty-dir"), emptyStr, emptyStr, true},

		{"Destination is empty", JoinPath(srcRoot, "text.txt"), emptyStr, emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent directory does", JoinPath(srcRoot, "text1.txt"), JoinPath(destRoot, "new1.txt"), JoinPath(bkRoot, "text1.txt"), JoinPath(destRoot, "new1.txt"), false},
		{"Destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Destination is a file", JoinPath(srcRoot, "text2.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text2.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to file", JoinPath(srcRoot, "text3.txt"), JoinPath(destRoot, "link.txt"), JoinPath(bkRoot, "text3.txt"), JoinPath(destRoot, "link.txt"), false},
		{"Destination is a symlink to directory (non-Windows)", JoinPath(srcRoot, "text4.txt"), JoinPath(destRoot, "link-dir"), JoinPath(bkRoot, "text4.txt"), JoinPath(destRoot, "link-dir"), false},
		{"Destination is a circular symlink", JoinPath(srcRoot, "text5.txt"), JoinPath(destRoot, "link-circular"), JoinPath(bkRoot, "text5.txt"), JoinPath(destRoot, "link-circular"), false},
		{"Destination is a broken symlink", JoinPath(srcRoot, "text6.txt"), JoinPath(destRoot, "link-broken"), JoinPath(bkRoot, "text6.txt"), JoinPath(destRoot, "link-broken"), false},
		{"Destination is an empty directory", JoinPath(srcRoot, "text7.txt"), JoinPath(destRoot, "empty-dir"), JoinPath(bkRoot, "text7.txt"), JoinPath(destRoot, "empty-dir", "text7.txt"), false},
		{"Destination is a directory containing other files", JoinPath(srcRoot, "text8.txt"), JoinPath(destRoot, "other-dir"), JoinPath(bkRoot, "text8.txt"), JoinPath(destRoot, "other-dir", "text8.txt"), false},
		{"Destination is a directory containing a file with the same name", JoinPath(srcRoot, "text9.txt"), JoinPath(destRoot, "same-dir"), JoinPath(bkRoot, "text9.txt"), JoinPath(destRoot, "same-dir", "text9.txt"), false},
		{"Destination file got no permissions", JoinPath(srcRoot, "text10.txt"), JoinPath(destRoot, "no_perm_file"), JoinPath(bkRoot, "text10.txt"), JoinPath(destRoot, "no_perm_file"), false},
		{"Destination directory got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Destination is a symlink to source file (non-Windows)", JoinPath(srcRoot, "self_text.txt"), JoinPath(srcRoot, "self_link.txt"), JoinPath(bkRoot, "self_text.txt"), JoinPath(srcRoot, "self_link.txt"), false},

		{"Rename: move empty file", JoinPath(srcRoot, "empty1.txt"), JoinPath(destRoot, "empty1.txt"), JoinPath(bkRoot, "empty1.txt"), JoinPath(destRoot, "empty1.txt"), false},
		{"Rename: move large text file", JoinPath(srcRoot, "large1.txt"), JoinPath(destRoot, "large1.txt"), JoinPath(bkRoot, "large1.txt"), JoinPath(destRoot, "large1.txt"), false},
		{"Rename: move image file", JoinPath(srcRoot, "image1.png"), JoinPath(destRoot, "image1.png"), JoinPath(bkRoot, "image1.png"), JoinPath(destRoot, "image1.png"), false},
		{"Cross-device: move empty file", JoinPath(srcRoot, "empty2.txt"), JoinPath(writeDevice, "empty2.txt"), JoinPath(bkRoot, "empty2.txt"), JoinPath(writeDevice, "empty2.txt"), false},
		{"Cross-device: move large text file", JoinPath(srcRoot, "large2.txt"), JoinPath(writeDevice, "large2.txt"), JoinPath(bkRoot, "large2.txt"), JoinPath(writeDevice, "large2.txt"), false},
		{"Cross-device: move image file", JoinPath(srcRoot, "image2.png"), JoinPath(writeDevice, "image2.png"), JoinPath(bkRoot, "image2.png"), JoinPath(writeDevice, "image2.png"), false},

		{"Cross-device: destination doesn't exist but its parent directory does", JoinPath(srcRoot, "file1.txt"), JoinPath(writeDevice, "new1.txt"), JoinPath(bkRoot, "file1.txt"), JoinPath(writeDevice, "new1.txt"), false},
		{"Cross-device: destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(writeDevice, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination exists and is a file", JoinPath(srcRoot, "file2.txt"), JoinPath(writeDevice, "text.txt"), JoinPath(bkRoot, "file2.txt"), JoinPath(writeDevice, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to file", JoinPath(srcRoot, "file3.txt"), JoinPath(writeDevice, "link.txt"), JoinPath(bkRoot, "file3.txt"), JoinPath(writeDevice, "link.txt"), false},
		{"Cross-device: destination exists and is a symlink to directory", JoinPath(srcRoot, "file4.txt"), JoinPath(writeDevice, "link-dir"), JoinPath(bkRoot, "file4.txt"), JoinPath(writeDevice, "link-dir"), false},
		{"Cross-device: destination exists and is a circular symlink", JoinPath(srcRoot, "file5.txt"), JoinPath(writeDevice, "link-circular"), JoinPath(bkRoot, "file5.txt"), JoinPath(writeDevice, "link-circular"), false},
		{"Cross-device: destination exists and is a broken symlink", JoinPath(srcRoot, "file6.txt"), JoinPath(writeDevice, "link-broken"), JoinPath(bkRoot, "file6.txt"), JoinPath(writeDevice, "link-broken"), false},
		{"Cross-device: destination exists and is a directory", JoinPath(srcRoot, "file7.txt"), JoinPath(writeDevice), JoinPath(bkRoot, "file7.txt"), JoinPath(writeDevice, "file7.txt"), false},
		{"Cross-device: source file got no permissions", JoinPath(srcRoot, "no_perm"), JoinPath(writeDevice, "new_noperm.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination directory got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(writeDevice, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Cross-device: destination file got no permissions", JoinPath(srcRoot, "file8.txt"), JoinPath(writeDevice, "no_perm_file"), JoinPath(bkRoot, "file8.txt"), JoinPath(writeDevice, "no_perm_file"), false},
		{"Cross-device: destination got no spaces for large file", JoinPath(srcRoot, "xlarge-text.txt"), JoinPath(writeDevice, "text.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination is a read-only device", JoinPath(srcRoot, "text.txt"), JoinPath(readOnlyDevice, "new.txt"), emptyStr, emptyStr, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (strings.Contains(tt.name, "permission") || strings.Contains(tt.name, "non-Windows")) && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}
			if strings.Contains(tt.name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice)) {
				t.Skipf("Skipping %q for missing RAM disk", tt.name)
			}

			if err := MoveFile(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("MoveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameFileContent(tt.backupPath, tt.outputPath)
				if err != nil {
					t.Errorf("MoveFile() got error while comparing the files: %v, %v, error: %v", tt.backupPath, tt.outputPath, err)
				} else if !same {
					t.Errorf("MoveFile() the files are not the same: %v, %v", tt.backupPath, tt.outputPath)
					return
				}
			}
		})
	}
}

func TestMoveSymlink(t *testing.T) {
	var (
		bkRoot         = JoinPath(resourceMoveSymlinkRoot, "backup")
		src1Root       = JoinPath(resourceMoveSymlinkRoot, "source1")
		src2Root       = JoinPath(resourceMoveSymlinkRoot, "source2")
		destRoot       = JoinPath(resourceMoveSymlinkRoot, "destination")
		writeDevice    = JoinPath(resourceReadWriteDevice, "move_link")
		readOnlyDevice = JoinPath(resourceReadOnlyDevice, "move_link")
	)

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		backupPath string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, JoinPath(destRoot, "link.txt"), emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(src1Root, "missing-link.txt"), JoinPath(destRoot, "link.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to file", JoinPath(src1Root, "link0.txt"), JoinPath(destRoot, "new-link-file.txt"), JoinPath(bkRoot, "link0.txt"), JoinPath(destRoot, "new-link-file.txt"), false},
		{"Source is a symlink to directory", JoinPath(src1Root, "link-dir"), JoinPath(destRoot, "new-link-dir"), JoinPath(bkRoot, "link-dir"), JoinPath(destRoot, "new-link-dir"), false},
		{"Source is a circular symlink", JoinPath(src1Root, "link-circular"), JoinPath(destRoot, "new-link-circular"), JoinPath(bkRoot, "link-circular"), JoinPath(destRoot, "new-link-circular"), false},
		{"Source is a broken symlink", JoinPath(src1Root, "link-broken"), JoinPath(destRoot, "new-link-broken"), JoinPath(bkRoot, "link-broken"), JoinPath(destRoot, "new-link-broken"), false},
		{"Source and destination is the same file", JoinPath(src1Root, "text.txt"), JoinPath(src1Root, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same symlink", JoinPath(src1Root, "link.txt"), JoinPath(src1Root, "link.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same directory", JoinPath(src1Root, "empty-dir"), JoinPath(src1Root, "empty-dir"), emptyStr, emptyStr, true},

		{"Destination is empty", JoinPath(src1Root, "link.txt"), emptyStr, emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent directory does", JoinPath(src1Root, "link1.txt"), JoinPath(destRoot, "new-link1.txt"), JoinPath(bkRoot, "link1.txt"), JoinPath(destRoot, "new-link1.txt"), false},
		{"Destination and its parent directory don't exist", JoinPath(src1Root, "link.txt"), JoinPath(destRoot, "missing-dir", "new-link.txt"), emptyStr, emptyStr, true},
		{"Destination is a file", JoinPath(src1Root, "link2.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "link2.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to file", JoinPath(src1Root, "link3.txt"), JoinPath(destRoot, "link.txt"), JoinPath(bkRoot, "link3.txt"), JoinPath(destRoot, "link.txt"), false},
		{"Destination is a symlink to directory (non-Windows)", JoinPath(src1Root, "link4.txt"), JoinPath(destRoot, "link-dir"), JoinPath(bkRoot, "link4.txt"), JoinPath(destRoot, "link-dir"), false},
		{"Destination is a circular symlink", JoinPath(src1Root, "link5.txt"), JoinPath(destRoot, "link-circular"), JoinPath(bkRoot, "link5.txt"), JoinPath(destRoot, "link-circular"), false},
		{"Destination is a broken symlink", JoinPath(src1Root, "link6.txt"), JoinPath(destRoot, "link-broken"), JoinPath(bkRoot, "link6.txt"), JoinPath(destRoot, "link-broken"), false},
		{"Destination is an empty directory", JoinPath(src1Root, "link7.txt"), JoinPath(destRoot, "empty-dir"), JoinPath(bkRoot, "link7.txt"), JoinPath(destRoot, "empty-dir", "link7.txt"), false},
		{"Destination is a directory containing other files", JoinPath(src1Root, "link8.txt"), JoinPath(destRoot, "other-dir"), JoinPath(bkRoot, "link8.txt"), JoinPath(destRoot, "other-dir", "link8.txt"), false},
		{"Destination is a directory containing a file with the same name", JoinPath(src1Root, "link9.txt"), JoinPath(destRoot, "same-dir"), JoinPath(bkRoot, "link9.txt"), JoinPath(destRoot, "same-dir", "link9.txt"), false},
		{"Destination file got no permissions", JoinPath(src1Root, "link10.txt"), JoinPath(destRoot, "no_perm_file"), JoinPath(bkRoot, "link10.txt"), JoinPath(destRoot, "no_perm_file"), false},
		{"Destination directory got no permissions", JoinPath(src1Root, "link.txt"), JoinPath(destRoot, "no_perm_dir"), emptyStr, emptyStr, true},

		{"Cross-device: destination doesn't exist but its parent directory does", JoinPath(src2Root, "link1.txt"), JoinPath(writeDevice, "new-link1.txt"), JoinPath(bkRoot, "link1.txt"), JoinPath(writeDevice, "new-link1.txt"), false},
		{"Cross-device: destination and its parent directory don't exist", JoinPath(src2Root, "link.txt"), JoinPath(writeDevice, "missing-dir", "new-link.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination exists and is a file", JoinPath(src2Root, "link2.txt"), JoinPath(writeDevice, "text.txt"), JoinPath(bkRoot, "link2.txt"), JoinPath(writeDevice, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to file", JoinPath(src2Root, "link3.txt"), JoinPath(writeDevice, "link.txt"), JoinPath(bkRoot, "link3.txt"), JoinPath(writeDevice, "link.txt"), false},
		{"Cross-device: destination exists and is a symlink to directory", JoinPath(src2Root, "link4.txt"), JoinPath(writeDevice, "link-dir"), JoinPath(bkRoot, "link4.txt"), JoinPath(writeDevice, "link-dir"), false},
		{"Cross-device: destination exists and is a circular symlink", JoinPath(src2Root, "link5.txt"), JoinPath(writeDevice, "link-circular"), JoinPath(bkRoot, "link5.txt"), JoinPath(writeDevice, "link-circular"), false},
		{"Cross-device: destination exists and is a broken symlink", JoinPath(src2Root, "link6.txt"), JoinPath(writeDevice, "link-broken"), JoinPath(bkRoot, "link6.txt"), JoinPath(writeDevice, "link-broken"), false},
		{"Cross-device: destination exists and is a directory", JoinPath(src2Root, "link7.txt"), JoinPath(writeDevice), JoinPath(bkRoot, "link7.txt"), JoinPath(writeDevice, "link7.txt"), false},
		{"Cross-device: source file got no permissions", JoinPath(src2Root, "no_perm"), JoinPath(writeDevice, "new_noperm.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination directory got no permissions", JoinPath(src2Root, "link.txt"), JoinPath(writeDevice, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Cross-device: destination file got no permissions", JoinPath(src2Root, "link8.txt"), JoinPath(writeDevice, "no_perm_file"), JoinPath(bkRoot, "link8.txt"), JoinPath(writeDevice, "no_perm_file"), false},
		{"Cross-device: destination is a read-only device", JoinPath(src2Root, "link.txt"), JoinPath(readOnlyDevice, "new-link.txt"), emptyStr, emptyStr, true},

		{"Cross-device: source is a symlink to file", JoinPath(src2Root, "link.txt"), JoinPath(writeDevice, "cd-link1.txt"), JoinPath(bkRoot, "link.txt"), JoinPath(writeDevice, "cd-link1.txt"), false},
		{"Cross-device: source is a symlink to directory", JoinPath(src2Root, "link-dir"), JoinPath(writeDevice, "cd-link2.txt"), JoinPath(bkRoot, "link-dir"), JoinPath(writeDevice, "cd-link2.txt"), false},
		{"Cross-device: source is a circular symlink", JoinPath(src2Root, "link-circular"), JoinPath(writeDevice, "cd-link3.txt"), JoinPath(bkRoot, "link-circular"), JoinPath(writeDevice, "cd-link3.txt"), false},
		{"Cross-device: source is a broken symlink", JoinPath(src2Root, "link-broken"), JoinPath(writeDevice, "cd-link4.txt"), JoinPath(bkRoot, "link-broken"), JoinPath(writeDevice, "cd-link4.txt"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (strings.Contains(tt.name, "permission") || strings.Contains(tt.name, "non-Windows")) && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}
			if strings.Contains(tt.name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice)) {
				t.Skipf("Skipping %q for missing RAM disk", tt.name)
			}

			if err := MoveSymlink(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("MoveSymlink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameSymlinkContent(tt.backupPath, tt.outputPath)
				if err != nil {
					t.Errorf("MoveSymlink() got error while comparing the files: %v, %v, error: %v", tt.backupPath, tt.outputPath, err)
				} else if !same {
					t.Errorf("MoveSymlink() the files are not the same: %v, %v", tt.backupPath, tt.outputPath)
					return
				}
			}
		})
	}
}

func TestMoveDir(t *testing.T) {
	var (
		bkRoot         = JoinPath(resourceMoveDirRoot, "backup")
		src1Root       = JoinPath(resourceMoveDirRoot, "source1")
		src2Root       = JoinPath(resourceMoveDirRoot, "source2")
		destRoot       = JoinPath(resourceMoveDirRoot, "destination")
		writeDevice    = JoinPath(resourceReadWriteDevice, "move_file")
		readOnlyDevice = JoinPath(resourceReadOnlyDevice, "move_file")
	)

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		backupPath string
		outputPath string
		wantErr    bool
	}{
		{"Source is empty", emptyStr, JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(src1Root, "missing-text.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to file", JoinPath(src1Root, "link.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to directory", JoinPath(src1Root, "link-dir"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a circular symlink", JoinPath(src1Root, "link-circular"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a broken symlink", JoinPath(src1Root, "link-broken"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source directory contains file got no permissions", JoinPath(src1Root, "link-broken"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same file", JoinPath(src1Root, "text.txt"), JoinPath(src1Root, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same symlink", JoinPath(src1Root, "link.txt"), JoinPath(src1Root, "link.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same directory", JoinPath(src1Root, "empty-dir"), JoinPath(src1Root, "empty-dir"), emptyStr, emptyStr, true},

		{"Destination is empty", JoinPath(src1Root, "text.txt"), emptyStr, emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent directory does", JoinPath(src1Root, "text1.txt"), JoinPath(destRoot, "new1.txt"), JoinPath(bkRoot, "text1.txt"), JoinPath(destRoot, "new1.txt"), false},
		{"Destination and its parent directory don't exist", JoinPath(src1Root, "text.txt"), JoinPath(destRoot, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Destination is a file", JoinPath(src1Root, "text2.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text2.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to file", JoinPath(src1Root, "text3.txt"), JoinPath(destRoot, "link.txt"), JoinPath(bkRoot, "text3.txt"), JoinPath(destRoot, "link.txt"), false},
		{"Destination is a symlink to directory (non-Windows)", JoinPath(src1Root, "text4.txt"), JoinPath(destRoot, "link-dir"), JoinPath(bkRoot, "text4.txt"), JoinPath(destRoot, "link-dir"), false},
		{"Destination is a circular symlink", JoinPath(src1Root, "text5.txt"), JoinPath(destRoot, "link-circular"), JoinPath(bkRoot, "text5.txt"), JoinPath(destRoot, "link-circular"), false},
		{"Destination is a broken symlink", JoinPath(src1Root, "text6.txt"), JoinPath(destRoot, "link-broken"), JoinPath(bkRoot, "text6.txt"), JoinPath(destRoot, "link-broken"), false},
		{"Destination is an empty directory", JoinPath(src1Root, "text7.txt"), JoinPath(destRoot, "empty-dir"), JoinPath(bkRoot, "text7.txt"), JoinPath(destRoot, "empty-dir", "text7.txt"), false},
		{"Destination is a directory containing other files", JoinPath(src1Root, "text8.txt"), JoinPath(destRoot, "other-dir"), JoinPath(bkRoot, "text8.txt"), JoinPath(destRoot, "other-dir", "text8.txt"), false},
		{"Destination is a directory containing a file with the same name", JoinPath(src1Root, "text9.txt"), JoinPath(destRoot, "same-dir"), JoinPath(bkRoot, "text9.txt"), JoinPath(destRoot, "same-dir", "text9.txt"), false},
		{"Destination file got no permissions", JoinPath(src1Root, "text10.txt"), JoinPath(destRoot, "no_perm_file"), JoinPath(bkRoot, "text10.txt"), JoinPath(destRoot, "no_perm_file"), false},
		{"Destination directory got no permissions", JoinPath(src1Root, "text.txt"), JoinPath(destRoot, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Destination is a symlink to source file (non-Windows)", JoinPath(src1Root, "self_text.txt"), JoinPath(src1Root, "self_link.txt"), JoinPath(bkRoot, "self_text.txt"), JoinPath(src1Root, "self_link.txt"), false},

		{"Rename: move an empty directory", JoinPath(src1Root, "empty1.txt"), JoinPath(destRoot, "empty1.txt"), JoinPath(bkRoot, "empty1.txt"), JoinPath(destRoot, "empty1.txt"), false},
		{"Rename: move a directory contains only files", JoinPath(src1Root, "large1.txt"), JoinPath(destRoot, "large1.txt"), JoinPath(bkRoot, "large1.txt"), JoinPath(destRoot, "large1.txt"), false},
		{"Rename: move a directory contains files, symlinks and directories", JoinPath(src1Root, "image1.png"), JoinPath(destRoot, "image1.png"), JoinPath(bkRoot, "image1.png"), JoinPath(destRoot, "image1.png"), false},

		{"Cross-device: destination doesn't exist but its parent directory does", JoinPath(src2Root, "file1.txt"), JoinPath(writeDevice, "new1.txt"), JoinPath(bkRoot, "file1.txt"), JoinPath(writeDevice, "new1.txt"), false},
		{"Cross-device: destination and its parent directory don't exist", JoinPath(src2Root, "text.txt"), JoinPath(writeDevice, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination exists and is a file", JoinPath(src2Root, "file2.txt"), JoinPath(writeDevice, "text.txt"), JoinPath(bkRoot, "file2.txt"), JoinPath(writeDevice, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to file", JoinPath(src2Root, "file3.txt"), JoinPath(writeDevice, "link.txt"), JoinPath(bkRoot, "file3.txt"), JoinPath(writeDevice, "link.txt"), false},
		{"Cross-device: destination exists and is a symlink to directory", JoinPath(src2Root, "file4.txt"), JoinPath(writeDevice, "link-dir"), JoinPath(bkRoot, "file4.txt"), JoinPath(writeDevice, "link-dir"), false},
		{"Cross-device: destination exists and is a circular symlink", JoinPath(src2Root, "file5.txt"), JoinPath(writeDevice, "link-circular"), JoinPath(bkRoot, "file5.txt"), JoinPath(writeDevice, "link-circular"), false},
		{"Cross-device: destination exists and is a broken symlink", JoinPath(src2Root, "file6.txt"), JoinPath(writeDevice, "link-broken"), JoinPath(bkRoot, "file6.txt"), JoinPath(writeDevice, "link-broken"), false},
		{"Cross-device: destination exists and is a directory", JoinPath(src2Root, "file7.txt"), JoinPath(writeDevice), JoinPath(bkRoot, "file7.txt"), JoinPath(writeDevice, "file7.txt"), false},
		{"Cross-device: source contains file got no permissions", JoinPath(src2Root, "no_perm"), JoinPath(writeDevice, "new_noperm.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination directory got no permissions", JoinPath(src2Root, "text.txt"), JoinPath(writeDevice, "no_perm_dir"), emptyStr, emptyStr, true},
		{"Cross-device: destination file got no permissions", JoinPath(src2Root, "file8.txt"), JoinPath(writeDevice, "no_perm_file"), JoinPath(bkRoot, "file8.txt"), JoinPath(writeDevice, "no_perm_file"), false},
		{"Cross-device: destination got no spaces for large file", JoinPath(src2Root, "xlarge-text.txt"), JoinPath(writeDevice, "text.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination is a read-only device", JoinPath(src2Root, "text.txt"), JoinPath(readOnlyDevice, "new.txt"), emptyStr, emptyStr, true},

		{"Cross-device: move an empty directory", JoinPath(src2Root, "empty2.txt"), JoinPath(writeDevice, "empty2.txt"), JoinPath(bkRoot, "empty2.txt"), JoinPath(writeDevice, "empty2.txt"), false},
		{"Cross-device: move a directory contains only files", JoinPath(src2Root, "large2.txt"), JoinPath(writeDevice, "large2.txt"), JoinPath(bkRoot, "large2.txt"), JoinPath(writeDevice, "large2.txt"), false},
		{"Cross-device: move a directory contains files, symlinks and directories", JoinPath(src2Root, "image2.png"), JoinPath(writeDevice, "image2.png"), JoinPath(bkRoot, "image2.png"), JoinPath(writeDevice, "image2.png"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (strings.Contains(tt.name, "permission") || strings.Contains(tt.name, "non-Windows")) && IsOnWindows() {
				t.Skipf("Skipping %q for Windows", tt.name)
			}
			if strings.Contains(tt.name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice)) {
				t.Skipf("Skipping %q for missing RAM disk", tt.name)
			}

			if err := MoveDir(tt.srcPath, tt.destPath); (err != nil) != tt.wantErr {
				t.Errorf("MoveDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				same, err := SameDirEntries(tt.backupPath, tt.outputPath)
				if err != nil {
					t.Errorf("MoveDir() got error while comparing the files: %v, %v, error: %v", tt.backupPath, tt.outputPath, err)
				} else if !same {
					t.Errorf("MoveDir() the files are not the same: %v, %v", tt.backupPath, tt.outputPath)
					return
				}
			}
		})
	}
}
