package positive_integer

/*
Source: https://www.geeksforgeeks.org/find-the-smallest-positive-number-missing-from-an-unsorted-array-set-3/
*/
func FindMinPositive(arr []int) int {
	defaultSmallest :=1
	mySet := map[int]bool{} // for set implementation in go
	n := len(arr)
	for i:= 0; i < n; i++ {
		if defaultSmallest < arr[i] {
			mySet[arr[i]] = true
		} else if defaultSmallest == arr[i] {
			defaultSmallest = defaultSmallest + 1
			for range mySet {

				_, ok := mySet[defaultSmallest]
				if ok {
					delete(mySet, defaultSmallest)
					defaultSmallest = defaultSmallest + 1
				}
			}
		}
	}
	return  defaultSmallest
}

