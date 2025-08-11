package main

import "fmt"

func main() {
	x := 6
	y := 6
	if x > 5 && y > 5 {
		fmt.Println("x and y is greater than 5")
	} else if x < 5 && y < 5 {
		fmt.Println("x and y is less than 5")
	} else {
		fmt.Println("BYE")
	}

}
