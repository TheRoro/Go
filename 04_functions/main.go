package main

import (
	"fmt"
)

func greeting(name string) string {
	return "hello " + name
}

func getSum(n1 int, n2 int) int {
	return n1 + n2;
}

func main() {
	fmt.Println(greeting("Rodrigo"))
	fmt.Println(getSum(1, 2))
}