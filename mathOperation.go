package main

import "fmt"

func main() {
	var (
		x int = 1
		y int = 22
	)

	fmt.Println(x+y)

	// augmented assigment
	x += 232
	fmt.Println(x)

	// unary operator
	x++
	fmt.Println(x)
}