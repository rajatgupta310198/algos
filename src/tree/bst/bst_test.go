package bst

import "testing"

func TestBstInsert(t *testing.T) {
	elements := []int{4, 3, 1, 2, 10, -1, 6}
	bst := new(T)
	for i := range elements {
		if !bst.Insert(i) {
			t.Error("Should have inserted element")
		}
	}

	// if bst.Insert(4) {
	// 	t.Error("Should not insert redundant element")
	// }

}
