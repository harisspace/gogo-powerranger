package main

import "fmt"

func endApp() {
	// recover bisa tetap menjalankan code walaupun sudah panic dan menangkap errornya
	message := recover()

	fmt.Println("end Application", message)
}

func runApp(error bool) {
	// defer dapat menjalankan function walaupun terjadi error
	defer endApp()

	if error {
		// panic menyetop program
		panic("Application error")
	}
	fmt.Println("application running")
}

func main() {
	runApp(true)
	fmt.Println("haris")
}