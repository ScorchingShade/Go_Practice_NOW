package main

import "fmt"

const time int = 60

type car struct {
	Name  string
	gas   int
	price int
	speed int
}

//this is a value receiver, it only does calculations in a struct. To do modifications in struct use pointer receiver
func (c car) kmh() int {
	return c.gas * (c.speed / time)
}

func main() {
	p1 := car{"Ferrari", 50, 10, 300}

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.gas)
	fmt.Println(p1.price)

	fmt.Println(p1.kmh())
}
