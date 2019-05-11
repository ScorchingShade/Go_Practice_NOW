package main

import "fmt"

type person struct {
	Name       string
	Age        int
	Profession string
}

func main() {
	p1 := person{"Ankush", 19, "Student"}

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Profession)
}
