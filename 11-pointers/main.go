package main

import "fmt"

func modifyValuebyReference(value *int) {

	*value = *value * 2
	// fmt.Println("Value inside modifyValuebyReference:", *value)
}

func main() {
	// var num int
	// num = 10
	// var ptr *int
	// ptr = &num
	num := 10
	ptr := &num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value of ptr:", *ptr)
	fmt.Println("Address of ptr:", ptr)

	var ptr2 *int
	if ptr2 == nil { //by default pointer is nil
		fmt.Println("ptr2 is nil")
	} else {
		fmt.Println("ptr2 is not nil")
	}

	value := 5
	modifyValuebyReference(&value)
	fmt.Println("Value after modifyValuebyReference:", value)
}
