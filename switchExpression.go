package main

import "fmt"

func main() {
	// name := "haris"

	// switch name {
	// case "haris":
	// 	fmt.Println("hello haris")
	// case "budi":
	// 	fmt.Println("hello budi")
	// default:
	// 	fmt.Println("kenalan donk")
	// }

	// switch case shorthand
		switch name := "budi"; name {
		case "haris":
				fmt.Println("hello haris")
			case "budi":
				fmt.Println("hello budi")
			default:
				fmt.Println("kenalan donk")
		}

	// switch case without condition
	name1 := "budi"
	switch {
	case name1 == "haris":
		fmt.Println("hello haris")
	case name1 == "budi":
		fmt.Println("hello budi")
	default:
		fmt.Println("kenalan donk")
	}
}