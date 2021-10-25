package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi dengan nol")
	}else {
		return nilai/pembagi, nil
	}
}

func main() {
	result, err := pembagian(20, 0)

	if err == nil {
		fmt.Print("hasil", result)
	}else {
		fmt.Print("Error", err.Error())
	}
}