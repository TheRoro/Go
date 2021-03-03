package main

import "fmt"

func main() {
	x := 7
	y := 6

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	//else if

	color := "blue"
	
	if color == "red" {
		fmt.Println("the color is red")
	} else if color == "blue" {
		fmt.Println("the color is blue")
	} else {
		fmt.Println("the color is NOT red or blue")
	}

	//switch

	switch color {
	case "red":
		fmt.Println("the color is red")
	case "blue":
		fmt.Println("the color is blue")
	default:
		fmt.Println("the color is NOT red or blue")
	}
}
