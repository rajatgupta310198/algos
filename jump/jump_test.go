package jumpsearch

import "testing"

func TestJumpSearch(t *testing.T)  {
	arr := [] int { 0, 1, 1, 2, 3, 5, 8, 13, 21,
		34, 55, 89, 144, 233, 377, 610 }
	index := jumpSearch(55, arr)
	if index != 10 {
		t.Error("Failed")
	}
}
