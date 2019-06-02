package main

//buffering start

import "fmt"

func foo3(c chan int, someValue int) {
	c <- someValue * 5

}

func main() {

	fooVal := make(chan int)

	for i := 0; i < 10; i++ {
		go foo3(fooVal, i)
	}

	for item := range fooVal {
		fmt.Println(item)
	}
}
