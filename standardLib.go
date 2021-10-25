package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there this is greeting"

	fmt.Println(strings.Contains(greeting,"is"))
	fmt.Println(strings.Split(greeting," "))
	fmt.Println(strings.Index(greeting, "ll"))

	numberOne := []int{1,2,3,4534,34234,444}
	sort.Ints(numberOne)
	fmt.Println(numberOne)
	index := sort.SearchInts(numberOne, 444444)
	fmt.Println(index)
}