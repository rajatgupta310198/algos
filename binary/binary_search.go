package binarysearch

func binarySearch(element int, arr []int) int {

	low := 0
	high := len(arr)
	index := -1
	mid := (low + high) / 2

	for low <= high {
		i := mid
		if element == arr[i] {
			index = i
			break
		} else if element < arr[i] {
			high = mid
			mid = (low + high) / 2
		} else {
			low = mid
			mid = (low + high) / 2
		}
	}

	return index

}
