package permutaions

func isEven(n int) bool  {
	return n%2 == 0
}

func heapsPermutation(n int, arr string, permutations []string) []string  {
	if n == 1 {
		permutations = append(permutations, arr)
		return permutations
	} else {
		permutations = heapsPermutation(n -1, arr, permutations)

		for i:=0; i<n - 1; i++ {
			if isEven(n) {
				arrRune := []rune(arr)
				arrRune[i], arrRune[n - 1] = arrRune[n - 1], arrRune[i]
				arr = string(arrRune)
			} else {
				arrRune := []rune(arr)
				arrRune[0], arrRune[n - 1] = arrRune[n - 1], arrRune[0]
				arr = string(arrRune)
			}
			permutations = heapsPermutation(n -1, arr, permutations)
		}
		return permutations
	}
}


func HeapsPermutation(n int, arr string)  []string {
	var permutations []string
	permutations = heapsPermutation(n, arr, permutations)
	return  permutations
}

