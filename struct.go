package main

import "fmt"

type Component struct {
	name, address string
	age           int
}
// struct method
func (component Component) sayHello() {
	fmt.Printf("hello %v", component.name)
}

func main() {
	// var component Component
	// component.age = 12

	// struct literal
	// joko := Component{
	// 	name: "joko",
	// 	address: "jakarta",
	// 	age: 20,
	// }

	// joko := Component{"joko", "jakarta", 3}

	// struct method
	joko := Component{
		name: "joko",
	}

	joko.sayHello()
}