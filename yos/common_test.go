package yos

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/1set/gut/ystring"
)

var (
	emptyStr                string
	resourceReadWriteDevice string
	resourceReadOnlyDevice  string
	resourceProtectedDevice string
	testResourceRoot        string
)

func init() {
	resourceReadWriteDevice = os.Getenv("RAMDISK_WRITE")
	resourceReadOnlyDevice = os.Getenv("RAMDISK_READONLY")
	resourceProtectedDevice = os.Getenv("RAMDISK_PROTECT")
	testResourceRoot = os.Getenv("TESTRSSDIR")
}

func preconditionCheck(t *testing.T, name string) {
	if (strings.Contains(name, "permission") || strings.Contains(name, "non-Windows")) && IsOnWindows() {
		t.Skipf("Skipping %q for Windows", name)
	}
	if strings.Contains(name, "Cross-device") && (ystring.IsBlank(resourceReadWriteDevice) || ystring.IsBlank(resourceReadOnlyDevice) || ystring.IsBlank(resourceProtectedDevice)) {
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
	switch e := err.(type) {
	case *os.PathError:
		return "path", e.Err
	case *os.LinkError:
		return "path", e.Err
	case *os.SyscallError:
		return "syscall", e.Err
	default:
		return "normal", err
	}
}

func Test_opError(t *testing.T) {
	type args struct {
		op   string
		path string
		err  error
	}
	tests := []struct {
		name string
		args args
		want *os.PathError
	}{
		{"All arguments are default", args{emptyStr, emptyStr, nil}, &os.PathError{Op: emptyStr, Path: emptyStr}},
		{"Nil error", args{"op1", "p1", nil}, &os.PathError{Op: "op1", Path: "p1"}},
		{"Plain error", args{"o2", "p2", errors.New("flat")}, &os.PathError{Op: "o2", Path: "p2", Err: errors.New("flat")}},
		{"Unwrap LinkError", args{"o3", "p3", &os.LinkError{Op: "l", Old: "oo", New: "nn", Err: errors.New("my link")}}, &os.PathError{Op: "o3", Path: "p3", Err: errors.New("my link")}},
		{"Unwrap PathError", args{"o4", "p4", &os.PathError{Op: "a", Path: "b", Err: errors.New("my path")}}, &os.PathError{Op: "o4", Path: "p4", Err: errors.New("my path")}},
		{"Unwrap SyscallError", args{"o5", "p5", &os.SyscallError{Syscall: "ccc", Err: errors.New("my sys")}}, &os.PathError{Op: "o5", Path: "p5", Err: errors.New("my sys")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opError(tt.args.op, tt.args.path, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opError() = %v, want %v", got, tt.want)
			}
		})
	}
}
