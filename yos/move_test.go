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
		{"Sample", JoinPath(srcRoot, "text.txt"), JoinPath(destRoot, "text.txt"), JoinPath(bkRoot, "text.txt"), JoinPath(destRoot, "text.txt"), false},
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
