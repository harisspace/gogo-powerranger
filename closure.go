package main

import "fmt"

func main() {
	// dapat mengakses variable diluar scope
	counter := 0

	increment := func() {
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}