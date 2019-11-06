package yhash_test

import (
	"fmt"

	"github.com/an63/gut/yhash"
)

// This example calculates MD5 checksum of "doc.go" file in current folder.
func ExampleFileMD5() {
	hash, err := yhash.FileMD5("doc.go")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("MD5:", hash)
	// Output: MD5: 3d34538da89e40ff9cb90ca9abe9e45c
}
