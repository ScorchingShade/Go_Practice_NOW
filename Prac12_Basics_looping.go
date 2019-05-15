package main

//heads up, go does not have a while loop

import "fmt"

//basic for loop
// for loop without any parameter runs like while (true) or while (1) , that is forever
func main() {

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
}
