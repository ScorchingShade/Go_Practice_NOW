package main

// the defer statement will be evaluated when it gets hit, but wont complete till the surrounding function either completes or panics

import (
	"fmt"
)

func foo() {

	//the defer statement is clearly 1st in priority since its written first, but it actually will only be evaluated after the entire function runs.
	//this means, first the "Doing Something" statement will be evaluated.
	//if there are multiple defers, it follows first in last out order

	//that means the last defer statement in order will be executed the first

	defer fmt.Println("Seriously defer")
	defer fmt.Println("Defering something")
	fmt.Println("Doing something")
}

func foo2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		//basically counts down
	}
}

func main() {
	foo()

	foo2()
}
