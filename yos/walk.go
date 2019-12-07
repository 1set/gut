package yos

import (
	"os"
)

type PathInfo struct {
	Path string
	FullPath string
	Info *os.FileInfo
}

func ListDir(root string) (items []*PathInfo, err error) {
	panic("Todo: and check if full is needed, and case like /home/bob/../alice")
}

func ListDirMatch(root, pattern string) (items []*PathInfo, err error) {
	panic("Todo: only check filename part, and check if full is needed, and case like /home/bob/../alice")
}
