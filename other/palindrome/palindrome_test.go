package palindrome

import "testing"

func TestCheckPalindromeString(t *testing.T) {
	testString := "geeksskeeg"
	if !CheckPalindromeString(testString) {
		t.Error("Not passed")
	}
}
