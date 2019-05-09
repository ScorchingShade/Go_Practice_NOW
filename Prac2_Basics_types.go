package main

import (
	"fmt"
)

//You have to provide return type after declaring params in go function. Also use keyword func
func add(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func mult(x, y float64) float64 {
	return x * y
}

func div(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println("Golang calculator\n")

	var a, b, choice float64

	fmt.Println("Provide first number:\n")

	fmt.Scan(&a)
	fmt.Println("Provide second number:\n")

	fmt.Scan(&b)
	fmt.Println("Chose options\n1) Addition\n2) Subtraction\n3) Multiplication\n4) Division")

	//if we don't want to explicitly define the type and variables, go can figure that
	//out on its own during compilation we just need to use := symbol and no var or type.
	//limitation of this is that once compiled, the variable's type can never change during program
	//also default data types will be assigned. e.g if funtion add takes 32 bit args, the default type assigned to vars will be float64 and hence throw error.

	//e.g choice := 0
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println(add(a, b))
	case 2:
		fmt.Println(sub(a, b))
	case 3:
		fmt.Println(mult(a, b))
	case 4:
		fmt.Println(div(a, b))
	default:
		fmt.Println("Bad Input")

	}

}
