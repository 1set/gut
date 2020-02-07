package yos

import (
	"sort"
	"testing"
)

func TestSortListByName(t *testing.T) {
	fileList, err := ListFile(resourceListRoot)
	if err != nil {
		t.Errorf("SortListByName got error to list %v: %v", resourceListRoot, err)
	}

	sort.Stable(SortListByName(fileList))
	verifyTestResult(t, "SortListByName", expectedResultMap["SortByName"], fileList, nil)
}

func TestSortListBySize(t *testing.T) {
	fileList, err := ListFile(resourceListRoot)
	if err != nil {
		t.Errorf("SortListBySize got error to list %v: %v", resourceListRoot, err)
	}

	sort.Stable(SortListBySize(fileList))
	verifyTestResult(t, "SortListBySize", expectedResultMap["SortBySize"], fileList, nil)
}

func TestSortListByModTime(t *testing.T) {
	fileList, err := ListFile(resourceListRoot)
	if err != nil {
		t.Errorf("SortListByModTime got error to list %v: %v", resourceListRoot, err)
	}

	sort.Stable(SortListByModTime(fileList))
	verifyTestResult(t, "SortListByModTime", expectedResultMap["SortByModTime"], fileList, nil)
}
