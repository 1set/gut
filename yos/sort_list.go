package yos

// SortListByName implements sort.Interface based on the Info.Name() field of FilePathInfo.
type SortListByName []*FilePathInfo

func (n SortListByName) Len() int {
	return len(n)
}

func (n SortListByName) Less(i, j int) bool {
	return n[i].Info.Name() < n[j].Info.Name()
}

func (n SortListByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// SortListBySize implements sort.Interface based on the Info.Size() field of FilePathInfo.
type SortListBySize []*FilePathInfo

func (s SortListBySize) Len() int {
	return len(s)
}

func (s SortListBySize) Less(i, j int) bool {
	return s[i].Info.Size() < s[j].Info.Size()
}

func (s SortListBySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortListByModTime implements sort.Interface based on the Info.ModTime() field of FilePathInfo.
type SortListByModTime []*FilePathInfo

func (t SortListByModTime) Len() int {
	return len(t)
}

func (t SortListByModTime) Less(i, j int) bool {
	return t[i].Info.ModTime().Before(t[j].Info.ModTime())
}

func (t SortListByModTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
