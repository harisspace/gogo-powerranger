package main

import "fmt"

func main() {
	// loop in go must initial with for
	// x := 1

	// while loop in go
	// for x < 5 {
	// 	println("this is number :", x)
	// 	x++
	// }

	// for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("this is number: ", i)
	// }

	// looping an array or slices
	 slicesOne := []string{"luiigi","mario","luffy","naruto"}

	for i := 0; i < len(slicesOne); i++ {
		fmt.Printf("this is number %v \n", slicesOne[i])
	}

	// for in di go
	for index, value := range slicesOne {
		fmt.Printf("hello this is index %v and the value is: %v \n", index, value)
	}

	// break and continue di go
	for i := 0; i < 10; i++ {
		if (i == 2) {
			fmt.Println("this is number two so continue")
			continue
		}else if i == 8{
			fmt.Println("this is number eight so break")
			break
		}
		fmt.Println("this is value number:", i)
	}
}