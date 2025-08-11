package main

import (
	"fmt"
	Name "mygo/Service1"
)

func main() {
	fmt.Println("Hello, World!")
	Name.Myname("Hello from Myname function!")
	// Name.Myname2()
	fmt.Println(Name.Myname2())

	// Variable Declaration Way 1
	// var person string = "Shahmeer"
	// fmt.Println("Hello, " + person + "!")
	// var person2 = "Shahmeer2"
	// fmt.Println("Hello, " + person2 + "!")

	// var decide = false
	// var decide2 bool = true
	// println("decide is ", decide)
	// fmt.Println("decide2 is ", decide2)

	// var number int = 10
	// fmt.Println("number is ", number)
	// var number2 = 20
	// fmt.Println("number2 is ", number2)

	//Variable Declaration Way 2
	person3 := "Shahmeer3"
	fmt.Println("Hello, " + person3 + "!")

	Age := 20
	fmt.Println("Age is ", Age)

	//************************
	//variable declaration  with small letter can be only used in this file
	//variable declaration with capital letter can be used in other files and packages
	//******************SIMILARLY*******************
	//Function Declaration with small letter can be only used in this file
	//Function Declaration with capital letter can be used in other files and packages
	//************************

	var Public = "Data is Public and can be used in other files and packages"
	fmt.Println("Public is ", Public)
	var Private = "Data is Private and can be used in this file only"
	fmt.Println("Private is ", Private)
}
