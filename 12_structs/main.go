package main

import (
	"fmt"
	"strconv"
)

//Person ... I hate go-lint
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

//Value receiver
func (p Person) greet() string {
	return "Hello, my name is: " + p.firstName + " " +
		p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

//Pointer receiver
func (p *Person) hasBday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	p.lastName = spouseLastName
}

func main() {

	person1 := Person{firstName: "Rodrigo", lastName: "Ramirez", city: "Lima", gender: "Male", age: 20}
	person2 := Person{"Rodrigo", "Ramirezzz", "Lima", "Male", 20}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.hasBday()
	fmt.Println(person1.greet())
	person2.getMarried("Forsay")
	fmt.Println(person2.greet())
}
