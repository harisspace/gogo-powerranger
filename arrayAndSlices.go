package main

import "fmt"

func main() {
	// cara 1
	// var arrayOne [3]int32 = [3]int32{1,2,3}

	// // cara 2 
	// var arrayTwo = [3]int32{1,2,3}

	// // cara 3
	// arrayThree := [3]int{1,2,3}

	// // cara 4
	// arrayFour := [...]int{1,2,3,4,5,6}

	// // slice
	// // cara 1
	// var sliceOne []int32 = []int32{1,2,3}
	// sliceOne = append(sliceOne, 2000)

	// // cara 2
	// var sliceTwo = []int32{4,5,6}

	// // cara 3
	// sliceThree := []int32{7,8,9}
	

	// fmt.Println(arrayOne, arrayTwo, arrayThree, len(arrayOne), sliceOne, sliceTwo, sliceThree[:3], arrayFour)

	// copy slice
	newSlice := make([]string, 7, 10)

	fmt.Println(cap(newSlice))

	// append jika lebih dari capacity maka dibuat array baru
}