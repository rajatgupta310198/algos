package permutaions



func factorialRecursive(n int64)  int64 {
	if n == 1 {
		return 1
	}
	return n * factorialRecursive(n - 1)

}

func FactorialRecursive(n int64)  int64 {
	return factorialRecursive(n)
}
