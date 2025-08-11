package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World!", i)
	}
	//infinte loop
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		if counter > 3 {
			break
		}
		counter++
	}

	//range loop (provides index and value)
	for index, value := range []int{10, 22, 33, 44, 55} {
		fmt.Println("Index:", index, "Value:", value)
	}

	data := "Hello, World!"
	for index, value := range data {
		fmt.Println("Index:", index, "Value:", string(value))
		// fmt.Printf("Index is %d and Value is %c\n", index, value)
	}
}
