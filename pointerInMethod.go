package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr" + man.Name
}

func main() {
	haris := Man{"Haris"}
	haris.Married()

	fmt.Println(haris.Name)
}