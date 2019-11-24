package yos

import (
	"github.com/1set/gut/ystring"
	"os"
	"testing"
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
