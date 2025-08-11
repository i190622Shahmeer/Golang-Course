package main

import "fmt"

func main() {
	age := 20
	name := "Shahmeer"
	height := 5.5

	//****************fmt.Println() Usage***********************
	//fmt.Println() auto adds space between the values
	fmt.Println("age:", age, "name:", name, "height:", height)
	//Also it auto adds new line at the end of one fmt.Println() function used
	fmt.Println("Hello, World!")

	//****************fmt.Printf() Usage************************
	//Used for formatted printing (Specifying the format of the output)
	//%d is used for integer
	//%s is used for string
	//%.2f is used for float with 2 decimal points
	//\n is used for new line
	//%v is used for any type
	//%T is used for type of the variable
	//%t is used for boolean
	//%c is used for character
	//%p is used for pointer
	//%x is used for hexadecimal
	//%o is used for octal
	//%b is used for binary
	//%e is used for scientific notation
	//%E is used for scientific notation with capital E
	//%q is used for quoted string
	//%U is used for Unicode
	//%u is used for Unicode with capital U
	//%f is used for float
	//%g is used for float with no decimal points
	//%G is used for float with no decimal points and capital G
	fmt.Printf("age: %d name: %s height: %.2f\n", age, name, height)
	fmt.Printf(("Type of age: %T\n"), age)
	fmt.Printf(("Type of name: %T\n"), name)
	fmt.Printf(("Type of height: %T\n"), height)

}
