package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, banana, grape"
	// Split the string into a slice of strings
	parts := strings.Split(data, ", ")
	// Print the resulting slice
	fmt.Println("Split Result:", parts)

	str := "one one two two three four"
	count := strings.Count(str, "one")
	fmt.Println("Count of 'one':", count)

	str = "                      Hello,Go!     "
	// Trim leading and trailing spaces
	trimmedStr := strings.TrimSpace(str)
	fmt.Println("Trimmed String:", trimmedStr)

	str1 := "String1"
	str2 := "String2"
	// join two strings
	joinstr := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Joined String:", joinstr)
}
