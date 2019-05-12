package main

import "fmt"

const time1 int = 60

type car1 struct {
	Name  string
	gas   int
	price int
	speed int
}

// This is a pointer receiver an modifies value in struct
func (c car1) kmh1() int {
	return c.gas * (c.speed / time1)
}

func (c *car1) new_speed(speed int) {
	c.speed = speed
}
func main() {
	p1 := car1{"Ferrari", 50, 10, 300}

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.gas)
	fmt.Println(p1.price)

	fmt.Println(p1.kmh1())
	p1.new_speed(200)
	fmt.Println(p1.kmh1())

}
