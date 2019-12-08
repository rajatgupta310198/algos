package main

import (
	"../tree/bst"
	"fmt"
)

func main()  {
	b := new(bst.T)
	//b.Insert(1)
	a := [] int {1, 2, 3, 7, 4, 5, 9, 6, 12, 15}
	for i :=range a {
		fmt.Printf("%d ->", a[i])
		b.Insert(a[i])
		//fmt.Print('\n')
		//b.InOrderPrint()
	}

	b.InOrderPrint()
}

