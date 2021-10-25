package main

import "fmt"

func main() {
	type Married bool

	var isMarried Married = false

	fmt.Print(isMarried)
}