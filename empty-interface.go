package main

import "fmt"

// interface kosong tidak ada declarasi method satupun, sehingga semua tipe data merupakan implementasinya
func ups() interface{} {
	// return 1
	// return false
	return "upss"
}

func main() {
	value := ups()
	fmt.Println(value)
}