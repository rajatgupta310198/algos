package jumpsearch

import "math"

func jumpSearch(element int, arr [] int) int  {
	 n := len(arr)
	 step := math.Sqrt(float64(n))
	 prev :=0
	 for ; arr[int(math.Min(step, float64(n))) -1] < element; {
		prev = int(step)
		step += math.Sqrt(float64(n))
		if prev >=n {
			return -1
		}

	 }

	 for ; arr[prev] < element; {
		prev += 1;
		if prev == int(math.Min(step, float64(n))) {
			return -1
		}
	 }

	 if arr[prev] == element {
	 	return  prev
	 }
	 return -1
}
