package main

import "fmt"

func foo1(c chan int, someValue int) {
	c <- someValue * 5

}

func main() {
	fooVal := make(chan int)

	go foo1(fooVal, 5)

	go foo1(fooVal, 3)

	//v1 := <- fooVal
	//v2 := <- fooVal

	v1, v2 := <-fooVal, <-fooVal

	fmt.Println(v1, v2)
}
