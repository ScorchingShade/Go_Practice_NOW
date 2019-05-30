package main

// the defer statement will be evaluated when it gets hit, but wont complete till the surrounding function either completes or panics

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

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

func saySync1(s string) {

	//previously we were in a situation where the for loop could come in an error and we would had gone in infinite waiting
	//now the wg1.Done will only work if the for loop works and hence we wont go in infinite wait
	defer wg1.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	foo()

	foo2()

	wg1.Add(1)
	go saySync1("hey")
	wg1.Add(1)
	go saySync1("there")
	wg1.Wait()

	wg1.Add(1)
	go saySync1("There")
	wg1.Wait()
}
