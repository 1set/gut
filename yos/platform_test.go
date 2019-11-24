package yos

import (
	"os"
	"strconv"
	"testing"

	"github.com/1set/gut/ystring"
)

func TestIsOnPlatform(t *testing.T) {
	type testCaseDef struct {
		osName    string
		isMacOS   bool
		isWindows bool
		isLinux   bool
	}
	tests := []testCaseDef{
		{"MACOS", true, false, false},
		{"WINDOWS", false, true, false},
		{"LINUX", false, false, true},
	}

	currentOsName := os.Getenv("OS_NAME")
	if ystring.IsBlank(currentOsName) {
		t.Skipf("skipping since the envvar 'OS_NAME' is missing: %q", currentOsName)
	}

	var testCase *testCaseDef
	for _, tt := range tests {
		if tt.osName == currentOsName {
			testCase = &tt
			break
		}
	}
	if testCase == nil {
		t.Skipf("skipping since the platform %q is not supported", currentOsName)
	}

	if ok := IsOnMacOS(); testCase.isMacOS != ok {
		t.Errorf("IsOnMacOS() got = %v, want = %v", ok, testCase.isMacOS)
	}
	if ok := IsOnWindows(); testCase.isWindows != ok {
		t.Errorf("IsOnWindows() got = %v, want = %v", ok, testCase.isWindows)
	}
	if ok := IsOnLinux(); testCase.isLinux != ok {
		t.Errorf("IsOnLinux() got = %v, want = %v", ok, testCase.isLinux)
	}
}

func BenchmarkIsOnMacOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOnMacOS()
	}
}

func BenchmarkIsOnWindows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOnWindows()
	}
}

func BenchmarkIsOnLinux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOnLinux()
	}
}

func TestIsOnBitsOfArch(t *testing.T) {
	type testCaseDef struct {
		name    string
		intSize int
		is32b   bool
		is64b   bool
	}
	tests := []testCaseDef{
		{"32-bit", 32, true, false},
		{"64-bit", 64, false, true},
	}

	currIntSize := strconv.IntSize
	var testCase *testCaseDef
	for _, tt := range tests {
		if tt.intSize == currIntSize {
			testCase = &tt
			break
		}
	}
	if testCase == nil {
		t.Skipf("skipping since the %d-bit architecture is not supported", currIntSize)
	}

	if ok := IsOn32bitArch(); testCase.is32b != ok {
		t.Errorf("IsOn32bitArch() for case %q got = %v, want = %v", testCase.name, ok, testCase.is32b)
	}

	if ok := IsOn64bitArch(); testCase.is64b != ok {
		t.Errorf("IsOn64bitArch() for case %q got = %v, want = %v", testCase.name, ok, testCase.is64b)
	}
}

func BenchmarkIsOn32bitArch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOn32bitArch()
	}
}

func BenchmarkIsOn64bitArch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOn64bitArch()
	}
}
