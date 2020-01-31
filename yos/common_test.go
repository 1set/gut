package yos

import (
	"strings"
	"testing"

	"github.com/1set/gut/ystring"
)

func preconditionCheck(t *testing.T, name string) {
	if (strings.Contains(name, "permission") || strings.Contains(name, "non-Windows")) && IsOnWindows() {
		t.Skipf("Skipping %q for Windows", name)
	}
	if strings.Contains(name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice)) {
		t.Skipf("Skipping %q for missing RAM disk", name)
	}
}
