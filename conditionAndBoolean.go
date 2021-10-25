package main

import "fmt"

func main() {
	var (
		x int = 2
		y int = 9
	)

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)

	// if statement
	name := "haris"

	if name == "budi" {
		fmt.Println("hello budi")
	}else if name == "haris" {
		fmt.Println("hello haris")
	}else {
		fmt.Println("kenalan donk")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("nama terlalu panjang")
	}

	// if shorthand
	if length2 := len(name); length2 > 5 {
		fmt.Println("nama terlalu panjang")
	}
}