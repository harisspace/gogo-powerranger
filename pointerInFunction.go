package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}


func main() {
	jakarta := Address{"Jakarta Selatan", "DKI Jakarta", ""}

	changeCountryToIndonesia(&jakarta)

	fmt.Println(jakarta)
}