package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read the value from the pointer
	fmt.Println(*b)

	//Change the value with the pointer
	*b = 10
	fmt.Println(a)
}
