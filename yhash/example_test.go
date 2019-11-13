package yhash_test

import (
	"fmt"

	"github.com/1set/gut/yhash"
)

// This example calculates MD5 checksum of "doc.go" file in current folder.
func ExampleFileMD5() {
	hash, err := yhash.FileMD5("doc.go")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("MD5:", hash)
}

// This example calculates SHA-1 checksum of a given string.
func ExampleStringSHA1() {
	hash, err := yhash.StringSHA1("This page is intentionally left blank.")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("SHA1:", hash)
	// Output: SHA1: dfca586074b7cfd24abc19fee213de5558e57422
}

// This example calculates SHA-256 checksum of a slice of bytes.
func ExampleBytesSHA256() {
	hash, err := yhash.BytesSHA256([]byte("This page unintentionally left blank."))
	if err != nil {
		fmt.Println("got error:", err)
		return
	}

	fmt.Println("SHA256:", hash)
	// Output: SHA256: 19e25cb9879fa4510e6ebfacc8836597429da7d5b279ccc073249492de2eae10
}
