package main


// import formatted I/O package
import (
	"algos/src/Algo/other/permutaions"
	"fmt"
)


// We define a function add which takes 2 int variables as function parameters
// Returns addition of the parameters with type int
func add(a int, b int) int {
	return a + b
}

// We define another function compareA which returns bool
func compareA(a float64, b float64) bool {
	return a > b
}


// Any program in go start with main function inside "main' package
// You can put main inside any other folders like src/main/main.go
func main()  {
	fmt.Printf("Hello World!\n")

	// Notice type of variable declarations
	// Here we declare 'a' with type int
	var a int
	// After declaring 2 is assigned
	a = 2
	// We declare and assigned variable 'b' with value 4
	var b  = 4

	// This is another way of defining variable where we  put := for assigning data to variable in LHS
	// c is not declared we just assigned it with 4.2 which explicitly makes it a variable of float type with value 4.2
	c:=4.2

	d:=2.42

	// C style I/o printing with %d representing placeholder for in variables specifically of type int
	fmt.Printf("Addition of %d and %d is = %d\n", a, b, add(a, b))

	// %f is placeholder for float type and we have one additional thing %.1 which determines precision limit
	fmt.Printf("Value of c with floating 1 digit precision %.1f\n", c)

	fmt.Printf("Value of d with floating 2 digit precision %.2f\n", d)

	// %v is placeholder for bool
	fmt.Printf("Is %.2f is grater thant %.2f ? \nAnswer is %v\n", c, d, compareA(c, d))



	// importing permutations package created under src/Algo/other
	fmt.Println(permutaions.FactorialRecursive(10))

	stringVar := "abcd"
	permutationStrings := permutaions.HeapsPermutation(4, stringVar)
	fmt.Println(permutationStrings)
}
