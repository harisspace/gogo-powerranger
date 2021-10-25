package main

import "fmt"

func main() {
	var number8 int8 = 100
	var number16 int16 = int16(number8)
	var number64 int64 = 127

	var convertion int8 = int8(number64)

	fmt.Println(number8, number16, convertion)
}