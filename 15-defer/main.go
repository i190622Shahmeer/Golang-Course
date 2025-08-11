package main

import "fmt"

func add(a int, b int) int {
	// This function adds two integers and returns the result.
	// It is a simple example of a function in Go.
	return a + b
}

func main() {
	// Defer is used to delay the execution of a function until the surrounding function returns.
	// It is often used for cleanup tasks, such as closing files or releasing resources.
	// The deferred function will be executed in LIFO (Last In First Out) order.

	fmt.Println("Starting the main function")
	data := add(5, 10) // Call the add function and store the result in data
	defer fmt.Println("Ending the main function")
	defer fmt.Println("Result of add function:", data)
	fmt.Println("Middle of the main function")

}
