package main

import (
	"fmt"
	"main/mathutils"
	"strings"
)

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

	var val32 int32 = 128
	var val64 int64 = int64(val32)
	
	var val16 int8 = int8(val32)// overflow, back to minimum value of 

	fmt.Println(val32)
	fmt.Println(val64)
	fmt.Println(val16)

	//array
	var x = name[0]
	var xString = string(x)
	fmt.Println(name) 
	fmt.Println(x) //ascii number from J
	fmt.Println(xString)

	//Type -> alias 
	type NoKTP string
	var ktpJohn NoKTP ="1234567890123"

	fmt.Println(ktpJohn)
	fmt.Println(NoKTP("1234567890123"))

	//boolean
	var name1 = "Antron"
	var name2 = "antron"


	var result = name1 == name2
	fmt.Println(result)

	result = strings.EqualFold(name1, name2)// bandingin equal string, without case sensitive
	result = strings.ToLower(name1) == strings.ToLower(name2) // mengubah huruf kapital ke huruf kecil
	fmt.Println(result)
	fmt.Println("")

	//number
	num := 9
	res := mathutils.Square(num)
	fmt.Printf("Kuadrat dari %d adalah %d\n", num, res)

	w:= 12
	l:= 14
	res = mathutils.Area(w,l)
	fmt.Printf("luas suatu persegi adalah %d, apabila memiliki panjang %d dan lebar %d",res, w, l)

}
