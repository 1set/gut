package yos

import (
	"os"
	"path/filepath"
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

func expectedErrorCheck(t *testing.T, err error) {
	if err == nil {
		return
	}

	errType, errIn := unwrapErrorStruct(err)
	if errType != "path" && errIn != filepath.ErrBadPattern {
		t.Errorf("unexpected %s error: %v", errType, errIn)
	}

	errType, errNest := unwrapErrorStruct(errIn)
	if errType != "normal" {
		t.Errorf("nested %s error: %v", errType, errNest)
	}
}

func unwrapErrorStruct(err error) (string, error) {
	if err == nil {
		return "null", nil
	}
	switch err.(type) {
	case *os.PathError:
		errP, _ := err.(*os.PathError)
		return "path", errP.Err
	case *os.LinkError:
		errL, _ := err.(*os.LinkError)
		return "path", errL.Err
	case *os.SyscallError:
		errS, _ := err.(*os.SyscallError)
		return "syscall", errS.Err
	default:
		return "normal", err
	}
}
