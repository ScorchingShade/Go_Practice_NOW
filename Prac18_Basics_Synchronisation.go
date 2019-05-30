package main

import (
	"fmt"
	"sync"
	"time"
)

//we declared sync package to add a waitgroup
//with waitgroup we would be able to wait for a goroutine to finish before programme moves ahead
//when the programme finishes, the waitgroup is notified, which notifies the program in turn

var wg sync.WaitGroup

func saySync(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

	//here we notify the waitgroup about the task when it is done

	//one issue remains...what if this code has some runtime error and panics , the wait will be infinite, to solve that, move to next part in series.
	wg.Done()
}

func main() {

	//we simply add 1 to waitgroup when we need a goroutine to finish
	wg.Add(1)
	go saySync("hey")
	wg.Add(1)
	go saySync("there")
	wg.Wait()

	//with waitgroups we dont waste any time in sleep and loose processing time
	wg.Add(1)
	go saySync("There")
	wg.Wait()
}
