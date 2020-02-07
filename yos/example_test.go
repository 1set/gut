package yos_test

import (
	"fmt"
	"sort"

	"github.com/1set/gut/yos"
)

// This example moves all image files from sub-folders of the source directory to the same destination directory.
func Example() {
	srcRoot := "source"
	destRoot := "collect"

	srcFiles, err := yos.ListMatch(srcRoot, yos.ListIncludeFile|yos.ListRecursive|yos.ListToLower, "*.jpg", "*.jpeg", "*.png", "*.gif")
	if err != nil {
		fmt.Printf("fail to list %q: %v\n", srcRoot, err)
		return
	}
	sort.Stable(yos.SortListByModTime(srcFiles))

	cntMove, cntSkip := 0, 0
	for _, src := range srcFiles {
		destPath := yos.JoinPath(destRoot, src.Info.Name())
		// skip if the same file already exists
		if yos.ExistFile(destPath) {
			if same, err := yos.SameFileContent(src.Path, destPath); same && err == nil {
				cntSkip++
				continue
			}
		}

		if err := yos.MoveFile(src.Path, destPath); err == nil {
			cntMove++
		}
	}

	fmt.Printf("total: %d\nmove : %d\nskip : %d\nfail : %d\n", len(srcFiles), cntMove, cntSkip, len(srcFiles)-cntMove-cntSkip)
}

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
func ExampleExist() {
	switch {
	case yos.IsOnMacOS():
		fmt.Println("macOS", yos.Exist(".DS_Store"))
	case yos.IsOnWindows():
		fmt.Println("Windows", yos.Exist("Thumbs.db"))
	}
}

// This example lists all Go source code files in the current directory and sorts alphabetically.
func ExampleSortListByName() {
	if srcFiles, err := yos.ListMatch(".", yos.ListIncludeFile|yos.ListRecursive, "*.go"); err == nil {
		sort.Stable(yos.SortListByName(srcFiles))
	}
}

// This example lists all files in the current directory and sorts by size in descending order.
func ExampleSortListBySize() {
	if srcFiles, err := yos.ListMatch(".", yos.ListIncludeFile|yos.ListRecursive, "*"); err == nil {
		sort.Stable(sort.Reverse(yos.SortListBySize(srcFiles)))
	}
}

// This example lists all files in the current directory and sorts by last modified time in ascending order.
func ExampleSortListByModTime() {
	if srcFiles, err := yos.ListMatch(".", yos.ListIncludeFile|yos.ListRecursive, "*"); err == nil {
		sort.Stable(yos.SortListByModTime(srcFiles))
	}
}
