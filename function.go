package main

import (
	"fmt"
	"strconv"
)

// declaration
func iniFungsi(name string) {
	fmt.Println("hello %v", name)
}

// return value
func iniReturnValue(name string, age int) string {
	return "Hello " + name + "dengan umur: " + strconv.FormatInt(int64(age), age)
}

// return multiple value
func returnMultiple(name string) (string, string){
	return "hello" , name
}

// named return value
func namedReturnValue() (firstname, lastname string) {
	firstname = "Haris"
	lastname = "Akbar"
	
	// return firstname, lastname
	return
}

// variadic function -> hampir sama rest parameter di js
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {

		total += value
	}
	return total
}

// function value-> function datanya bisa disimpan di variable (first class citizen)
func funcValue() string {
	return "hello"
}

// function as parameter
type Filter func(string) string

func sayHelloToFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println(nameFilter)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}
	return name
}

// anonymous function -> hampir sama dengan js

// recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return value
	}else {
		return value * factorialRecursive(value - 1)
	}
}

func main() {
	iniFungsi("haris")
	data := iniReturnValue("haris", 20)
	hello, akbar := returnMultiple("akbar")
	firstname, lastname := namedReturnValue()
	slices := []int {1,2,3,4,5}
	total := sumAll(slices...)
	iniFunc := funcValue

	// anonymous func
	anonim := func(name string) string {
		return name
	}

	sayHelloToFilter("Budi", spamFilter)
	sayHelloToFilter("anjing", spamFilter)

	factorial := factorialRecursive(5)
	fmt.Println(data, hello, akbar, firstname, lastname, total, iniFunc(), funcValue(), anonim("haris akbar oke"), factorial)
}