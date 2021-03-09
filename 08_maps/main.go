package main

import "fmt"

func main() {
	//map -> [key]value
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)

	//Declaration with key and values
	emails2 := map[string]string{"Bob2": "Bob2@gmail.com", "Sharon2": "sharon2@gmail.com"}
	emails2["Mike2"] = "mike2@gmail.com"
	fmt.Println(emails2)
}
