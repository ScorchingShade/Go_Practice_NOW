package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func saySync(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go saySync("hey")
	wg.Add(1)
	go saySync("there")
	wg.Wait()

	wg.Add(1)
	go saySync("There")
	wg.Wait()
}

//hi
