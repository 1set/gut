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
)

func init() {
	resourceReadWriteDevice = os.Getenv("RAMDISK_WRITE")
	resourceReadOnlyDevice = os.Getenv("RAMDISK_READONLY")

	testResourceRoot := os.Getenv("TESTRSSDIR")
	resourceMoveFileRoot = JoinPath(testResourceRoot, "yos", "move_file")
}

func TestMoveFile(t *testing.T) {
	var (
		bkRoot         = JoinPath(resourceMoveFileRoot, "backup")
		srcRoot        = JoinPath(resourceMoveFileRoot, "source")
		destRoot       = JoinPath(resourceMoveFileRoot, "destination")
		writeDevice    = JoinPath(resourceReadWriteDevice, "move_file")
		readOnlyDevice = JoinPath(resourceReadOnlyDevice, "move_file")
	)

	t.Logf("%s - %s", writeDevice, readOnlyDevice)

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		backupPath string
		outputPath string
		wantErr    bool
	}{
		// {"Sample", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Source is empty", emptyStr, JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source doesn't exist", JoinPath(srcRoot, "missing-text.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to file", JoinPath(srcRoot, "link.txt"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a symlink to directory", JoinPath(srcRoot, "link-dir"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a circular symlink", JoinPath(srcRoot, "link-circular"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source is a broken symlink", JoinPath(srcRoot, "link-broken"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},
		//{"Source got no permissions", JoinPath(srcRoot, "no_perm"), JoinPath(destRoot, "text.txt"), emptyStr, emptyStr, true},	// just rename/relink, it works without perm

		{"Source and destination is the same file", JoinPath(srcRoot, "text.txt"), JoinPath(srcRoot, "text.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same symlink", JoinPath(srcRoot, "link.txt"), JoinPath(srcRoot, "link.txt"), emptyStr, emptyStr, true},
		{"Source and destination is the same directory", JoinPath(srcRoot, "empty-dir"), JoinPath(srcRoot, "empty-dir"), emptyStr, emptyStr, true},

		{"Destination is empty", JoinPath(srcRoot, "text.txt"), emptyStr, emptyStr, emptyStr, true},
		{"Destination doesn't exist but its parent directory does", JoinPath(srcRoot, "text1.txt"), JoinPath(destRoot, "new1.txt"), JoinPath(bkRoot, "text1.txt"), JoinPath(destRoot, "new1.txt"), false},
		{"Destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Destination is a file", JoinPath(srcRoot, "text2.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text2.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to file", JoinPath(srcRoot, "text3.txt"), JoinPath(destRoot, "link.txt"), JoinPath(bkRoot, "text3.txt"), JoinPath(destRoot, "link.txt"), false},
		{"Destination is a symlink to directory", JoinPath(srcRoot, "text4.txt"), JoinPath(destRoot, "link-dir"), JoinPath(bkRoot, "text4.txt"), JoinPath(destRoot, "link-dir"), false},
		{"Destination is a circular symlink", JoinPath(srcRoot, "text5.txt"), JoinPath(destRoot, "link-circular"), JoinPath(bkRoot, "text5.txt"), JoinPath(destRoot, "link-circular"), false},
		{"Destination is a broken symlink", JoinPath(srcRoot, "text6.txt"), JoinPath(destRoot, "link-broken"), JoinPath(bkRoot, "text6.txt"), JoinPath(destRoot, "link-broken"), false},
		{"Destination is an empty directory", JoinPath(srcRoot, "text7.txt"), JoinPath(destRoot, "empty-dir"), JoinPath(bkRoot, "text7.txt"), JoinPath(destRoot, "empty-dir", "text7.txt"), false},
		{"Destination is a directory containing other files", JoinPath(srcRoot, "text8.txt"), JoinPath(destRoot, "other-dir"), JoinPath(bkRoot, "text8.txt"), JoinPath(destRoot, "other-dir", "text8.txt"), false},
		{"Destination is a directory containing a file with the same name", JoinPath(srcRoot, "text9.txt"), JoinPath(destRoot, "same-dir"), JoinPath(bkRoot, "text9.txt"), JoinPath(destRoot, "same-dir", "text9.txt"), false},
		// {"Destination file got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "no_perm_file"), emptyStr, emptyStr, true},	// just rename/relink, it works without perm
		{"Destination directory got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "no_perm_dir"), emptyStr, emptyStr, true},

		{"Cross-device: destination doesn't exist but its parent directory does", JoinPath(srcRoot, "file1.txt"), JoinPath(writeDevice, "new1.txt"), JoinPath(bkRoot, "file1.txt"), JoinPath(writeDevice, "new1.txt"), false},
		{"Cross-device: destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(writeDevice, "missing-dir", "new-file.txt"), emptyStr, emptyStr, true},
		{"Cross-device: destination exists and is a file", JoinPath(srcRoot, "file2.txt"), JoinPath(writeDevice, "text.txt"), JoinPath(bkRoot, "file2.txt"), JoinPath(writeDevice, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to file", JoinPath(srcRoot, "file3.txt"), JoinPath(writeDevice, "link.txt"), JoinPath(bkRoot, "file3.txt"), JoinPath(writeDevice, "link.txt"), false},
		{"Cross-device: destination exists and is a symlink to directory", JoinPath(srcRoot, "file4.txt"), JoinPath(writeDevice, "link-dir"), JoinPath(bkRoot, "file4.txt"), JoinPath(writeDevice, "link-dir"), false},
		{"Cross-device: destination exists and is a circular symlink", JoinPath(srcRoot, "file5.txt"), JoinPath(writeDevice, "link-circular"), JoinPath(bkRoot, "file5.txt"), JoinPath(writeDevice, "link-circular"), false},
		{"Cross-device: destination exists and is a broken symlink", JoinPath(srcRoot, "file6.txt"), JoinPath(writeDevice, "link-broken"), JoinPath(bkRoot, "file6.txt"), JoinPath(writeDevice, "link-broken"), false},
		//{"Cross-device: destination exists and is a directory", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//{"Cross-device: source file got no permissions", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//{"Cross-device: destination file got no permissions", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//{"Cross-device: destination directory got no permissions", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//{"Cross-device: destination got no spaces for large file", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//{"Cross-device: destination is a read-only device", JoinPath(srcRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), JoinPath(bkRoot, "file.txt"), JoinPath(writeDevice, "new.txt"), false},
		//
		//{"Rename: move empty file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Rename: move small text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Rename: move large text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Rename: move image file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Cross-device: move empty file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Cross-device: move small text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Cross-device: move large text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		//{"Cross-device: move image file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "permission") && IsOnWindows() {
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
