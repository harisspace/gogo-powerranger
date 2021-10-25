package main

import "fmt"

// golang saat declarasi variable bukan berisi null/nil tapi default value misal boolean maka false
// nil hanya dapat di implementasi di function, slice, map, interface, pointer, channel

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := newMap("haris")

	if person == nil {
		fmt.Print("Nama kosong")
		panic("panic bro nama kosong")
	}else {
		fmt.Print(person)
	}
}