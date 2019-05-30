package main

import (
	"fmt"
	"time"
)

//concurrency is dealing with things as they come. You don't necessarily doing two tasks parallelly you are essentially
//just doing whatever tasks come to you as you go simultaneously

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	// a go routine in go is like a lightweight thread and is triggered by go keywords
	go say("hey")
	go say("there")

	//if both go routines are run before programme finishes, then they will run concurrently but nothing tells a programme what should finish first!
	//we need to make the programme wait for them to finish to display anything

	time.Sleep(time.Second)
}
