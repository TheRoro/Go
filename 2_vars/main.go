package main

import (
	"fmt"
)

func main() {
	//Inferred data type
	var weight = 1.2

	//Shorthand for Inferred
	name := "Brad"
	size := 1.3

	//Constant
	const isCool = true

	//Specified data type
	var age int32 = 35

	//double assignment
	name, email := "Brad", "brad@brad.com"

	fmt.Println(name, age, isCool, size, email, weight)
	fmt.Printf("%T\n", email)
}
