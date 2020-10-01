package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	got := binarySearch(4, []int{1, 2, 3, 4, 5})
	if got != 3 {
		t.Error("Failed")
	}
}
