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
