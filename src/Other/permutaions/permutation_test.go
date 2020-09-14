package permutaions

import (
	"testing"
)

func TestHeapsPermutation(t *testing.T) {
	arr := "abcdefghij"
	permutations := HeapsPermutation(10, arr)
	if len(permutations) != 3628800 {
		t.Error("Not all permutations")
	}

}

func TestFactorialRecursive(t *testing.T) {
	n := int64(5)
	factorial := FactorialRecursive(n)

	if factorial != 120 {
		t.Error("Wrong factorial")
	}
}
