package yos

import (
	"os"
	"strings"
	"testing"

	"github.com/1set/gut/ystring"
)

var (
	emptyStr                string
	resourceReadWriteDevice string
	resourceReadOnlyDevice  string
	testResourceRoot        string
)

func init() {
	resourceReadWriteDevice = os.Getenv("RAMDISK_WRITE")
	resourceReadOnlyDevice = os.Getenv("RAMDISK_READONLY")
	testResourceRoot = os.Getenv("TESTRSSDIR")
}

func preconditionCheck(t *testing.T, name string) {
	if (strings.Contains(name, "permission") || strings.Contains(name, "non-Windows")) && IsOnWindows() {
		t.Skipf("Skipping %q for Windows", name)
	}
	if strings.Contains(name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice)) {
		t.Skipf("Skipping %q for missing RAM disk", name)
	}
}
