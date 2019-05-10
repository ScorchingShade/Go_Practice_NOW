package main

import (
	"fmt"
)

func pointer_test() {
	// to specify a pointer simply use & equivalent to a var
	x := 5
	a := &x
	//print var
	fmt.Println(x)
	// print value to pointer
	fmt.Println(*a)
	// print address
	fmt.Println(&a)
}

func main() {

	pointer_test()
}
