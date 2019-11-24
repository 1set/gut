package yos

import (
	"runtime"
)

// IsOnMacOS indicates whether the application is running on macOS.
func IsOnMacOS() bool {
	return runtime.GOOS == `darwin`
}

// IsOnWindows indicates whether the application is running on Windows.
func IsOnWindows() bool {
	return runtime.GOOS == `windows`
}

// IsOnLinux indicates whether the application is running on Linux.
func IsOnLinux() bool {
	return runtime.GOOS == `linux`
}
