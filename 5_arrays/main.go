package main

import (
	"fmt"
)

func main() {
	// Arrays
	var fruits [2]string
	//Separated declaration
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	//Declare and assign at the same time
	fruits2 := []string{"Banana", "Watermelon", "Grape"}

	fmt.Println(len(fruits2))
	fmt.Println(fruits[1])
	fmt.Println(fruits2[1:2])
}