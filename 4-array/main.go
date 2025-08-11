package main

import "fmt"

func main() {
	fmt.Println("Arrays in go language")

	var name [5]string
	name[0] = "Shahmeer"
	name[1] = "Pasha"

	fmt.Println("Name is :", name)

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Numbers are :", numbers)
	fmt.Println("Length of numbers is :", len(numbers))
	fmt.Println("Value at index 0 is :", numbers[0])

}
