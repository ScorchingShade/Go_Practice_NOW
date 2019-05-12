package main

import "fmt"

const time1 int = 60

type car1 struct {
	Name  string
	gas   int
	price int
	speed int
}

func (c car1) kmh1() int {
	return c.gas * (c.speed / time1)
}

// This is a pointer receiver an modifies value in struct
func (c *car1) new_speed(speed int) {
	c.speed = speed
}

// This is function, returning a struct doing exact same shit as pointer receiver but is less efficient
func carstruct(c car1, speed int) car1 {
	c.speed = speed
	return c
}
func main() {
	p1 := car1{"Ferrari", 50, 10, 300}

	//basic prints before any pointer or value or functions
	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.gas)
	fmt.Println(p1.price)

	// a value receiver, can't modify only calculate
	fmt.Println(p1.kmh1())
	//a pointer receiver
	p1.new_speed(200)
	fmt.Println(p1.kmh1())

	//a function doing the exact same thing as pointer receiver but is less efficient.
	p1 = carstruct(p1, 100)
	fmt.Println(p1.kmh1())

}

//essentially pointer receiver can do the work of both value and pointer receiver. It can be used in both contexts
//The value receiver makes a copy of the struct and is suitable for small structs
//The pointer receiver is suitable for large structs that directly modify it.
