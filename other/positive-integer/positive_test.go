package positiveinteger

import "testing"

// Test file
func TestFindMinPositive(t *testing.T) {
	arr := []int{2, 3, -7, 6, 8, 1, -10, 15}

	smallestPositive := FindMinPositive(arr)
	if smallestPositive != 4 {
		t.Error("Check the algorithm")
	}
}
