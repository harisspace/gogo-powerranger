package main

import "fmt"

func main() {
	// declaration
	newMap := map[string]string{
		"name":    "haris",
		"address": "jakarta",
	}

	fmt.Println(newMap["name"])
	newMap["age"] = "2"
	fmt.Println(newMap)

	newMap2 := make(map[string]string, 6)
	newMap2["oke"] ="tentu"
	newMap2["ini dihapus"] = "oke hapus saja"

	delete(newMap2, "ini dihapus")

	fmt.Println(newMap2)
}