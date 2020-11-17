package ystring

import (
	"os"
	"testing"
)

func TestNewLine(t *testing.T) {
	var (
		winNewLine  = "\r\n"
		unixNewLine = "\n"
	)

	switch currentOsName := os.Getenv("OS_NAME"); {
	case currentOsName == "WINDOWS" && NewLine != winNewLine:
		t.Errorf("NewLine on Windows got = %q, want = %q", NewLine, winNewLine)
	case currentOsName == "MACOS" && NewLine != unixNewLine:
		t.Errorf("NewLine on macOS got = %q, want = %q", NewLine, unixNewLine)
	case currentOsName == "LINUX" && NewLine != unixNewLine:
		t.Errorf("NewLine on Linux got = %q, want = %q", NewLine, unixNewLine)
	}
}
