package main

//coming soon

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}

	wg2.Done()
}

func say2(s string) {
	defer cleanup()

	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)

		if i == 2 {
			panic("oh dear, a 2")
		}
	}
}

func main() {

	wg2.Add(1)
	go say2("This is first")
	wg2.Add(1)
	go say2("This is second")
	wg2.Wait()

	wg2.Add(1)
	go say2("This is third")
	wg2.Wait()
}
