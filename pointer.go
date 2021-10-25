package main

// di go default pass by value bukan pass by reference

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // address2 := &address1
	address3 := &Address{"Padang", "Sumatera Barat", "Indonesia"}

	// address1 := "string"
	// address2 := &address1

	// *address2 = "siomay"

	// function new -> pointer kosong nilainya
	address4 := new(Address)
	address4.City = "oke"


	fmt.Println(address1, *address2, address3, address4)
}