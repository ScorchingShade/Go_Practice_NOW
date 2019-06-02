package main

//buffering start

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func foo3(c chan int, someValue int) {
	defer wg3.Done()
	c <- someValue * 5

}

func main() {

	fooVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go foo3(fooVal, i)
	}
	wg3.Wait()
	close(fooVal)
	for item := range fooVal {
		fmt.Println(item)
	}

	//time.Sleep(time.Second*2)

}
