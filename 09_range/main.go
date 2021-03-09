package main

import "fmt"

func main() {
	ids := []int{12, 123, 145, 25}

	//loop using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Loop without index put underscore _
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Adds ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	//range with maps
	emails := map[string]string{"Bob": "Bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}
}
