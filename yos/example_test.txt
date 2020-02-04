package yos_test

import (
	"fmt"

	"github.com/1set/gut/yos"
)

// This example lists all files in the current directory and its sub-directories.
func ExampleListFile() {
	entries, err := yos.ListFile(".")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	for _, info := range entries {
		fmt.Println(info.Path, info.Info.Size())
	}
}

// This example lists all .txt and .md files in the current directory and its sub-directories.
func ExampleListMatch() {
	entries, err := yos.ListMatch(".", yos.ListIncludeFile|yos.ListRecursive|yos.ListToLower, "*.txt", "*.md")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	for _, info := range entries {
		fmt.Println(info.Path, info.Info.Size())
	}
}

// This example checks if the annoying thumbnail file exists in the current directory as per the operating system.
func ExampleIsExist() {
	switch {
	case yos.IsOnMacOS():
		fmt.Println("macOS", yos.Exist(".DS_Store"))
	case yos.IsOnWindows():
		fmt.Println("Windows", yos.Exist("Thumbs.db"))
	}
}
