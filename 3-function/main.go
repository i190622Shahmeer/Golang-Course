package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello, World!")
	simpleFunction()
	fmt.Println("sum:", add(3, 4))
}
