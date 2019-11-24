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

// IsOn32bitArch indicates whether the application is running on 32-bit architecture.
func IsOn32bitArch() bool {
	return (^uint(0) >> 31) == 1
}

// IsOn64bitArch indicates whether the application is running on 64-bit architecture.
func IsOn64bitArch() bool {
	return (^uint(0) >> 63) == 1
}
