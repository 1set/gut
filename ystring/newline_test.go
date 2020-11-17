package ystring

import (
	"os"
	"testing"
)

func TestNewLine(t *testing.T) {
	var (
		winStyle  = "\r\n"
		unixStyle = "\n"
	)

	currentOsName := os.Getenv("OS_NAME")

	if currentOsName == "WINDOWS" && NewLine != winStyle {
		t.Errorf("NewLine on Windows got = %q, want = %q", NewLine, winStyle)
	}
	if currentOsName == "MACOS" && NewLine != unixStyle {
		t.Errorf("NewLine on macOS got = %q, want = %q", NewLine, unixStyle)
	}
	if currentOsName == "LINUX" && NewLine != unixStyle {
		t.Errorf("NewLine on Linux got = %q, want = %q", NewLine, unixStyle)
	}
}
