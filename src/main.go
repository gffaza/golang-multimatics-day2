package main

import "fmt"

func main() {
	fmt.Println("Hello BSI Go for Future!")

	//variable -> nilai bisa berubah
	var name string = "John Doe"
	country := "Indonesia"

	println(name)
	println(country)

	//constant -> nilai tidak bisa berubah
	const pi = 3.14
	const AppName = "Belajar Golang"

	println(AppName)

}
