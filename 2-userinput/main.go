package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name = ""
	// This is a simple program that takes user input and prints it to the console.
	// It demonstrates how to read input from the user in Go.
	fmt.Println("Hey, What's your name?")

	//This fmt.Scan() only takes one word as input and skips after white space
	// fmt.Scan(&name)
	// fmt.Println("Hey, Mr." + name)

	//For full line input, we can use bufio package
	// bufio.NewReader() is used to create a new reader that reads from standard input (os.Stdin).
	reader := bufio.NewReader(os.Stdin)
	//\n is used to skip the new line character
	name, _ = reader.ReadString(('\n'))
	fmt.Println("Hey, Mr." + name)
}
