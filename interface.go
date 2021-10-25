package main

import "fmt"

// interface hanya berisi abstract method
type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var haris Person

	haris.Name = "Haris"

	sayHello(haris)
}