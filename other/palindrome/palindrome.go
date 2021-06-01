package palindrome

func CheckPalindromeString(str string) bool  {
	isPalindrome := true
	for i:=0; i < len(str)/2; i++ {
		if str[i] != str[len(str) -i - 1] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
