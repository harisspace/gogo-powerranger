package main

import "fmt"

// assertasi type data dari interface kosong
func random() interface{} {
	return "oke"
}

func main() {
	random := random()

	randomString := random.(string)
	// randomInt := random.(int) // panic

	fmt.Print(randomString)

	// using switch statement
	switch random := random.(type) {
	case string:
		fmt.Print(random, "string")
	case int:
		fmt.Print(random, "int")
	default:
		fmt.Print("Unknown")
	}

}