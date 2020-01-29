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

	t.Skipf("%s - %s", writeDevice, readOnlyDevice)

	tests := []struct {
		name       string
		srcPath    string
		destPath   string
		backupPath string
		outputPath string
		wantErr    bool
	}{
		// {"Sample", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Source is empty", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source doesn't exist", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source is a symlink to file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source is a symlink to directory", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source is a circular symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source is a broken symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Destination is empty", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination doesn't exist but its parent directory does", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a symlink to directory", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a directory containing other files", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a directory containing nothing", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a directory containing a file with the same name", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a circular symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination is a broken symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination file got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Destination directory got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Source and destination is the same file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source and destination is the same symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Source and destination is the same directory", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Cross-device: destination doesn't exist but its parent directory does", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination and its parent directory don't exist", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a symlink to directory", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a directory", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a circular symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination exists and is a broken symlink", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination file got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination directory got no permissions", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination got no spaces for large file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: destination is a read-only device", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},

		{"Rename: move empty file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Rename: move small text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Rename: move large text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Rename: move image file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: move empty file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: move small text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: move large text file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
		{"Cross-device: move image file", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
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
