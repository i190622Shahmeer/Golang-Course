package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are :", numbers)
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Numbers after appending 6, 7, 8 are :", numbers)
	fmt.Println("Length of numbers is :", len(numbers))
	fmt.Println("Capacity of numbers is :", cap(numbers))
	fmt.Printf("Datatype of numbers is %T\n", numbers)

	//declaring a null slice
	name := []string{}
	fmt.Println("Name is :", name)

	//initializing a slice with make function
	// make([]data_type, length, capacity)
	numbers2 := make([]int, 3, 5) //5 length and 10 capacity
	numbers2 = append(numbers2, 4)
	fmt.Println("Numbers2 are :", numbers2)
	fmt.Println("Length of numbers2 is :", len(numbers2))
	fmt.Println("Capacity of numbers2 is :", cap(numbers2))

	numbers2 = append(numbers2, 5)
	fmt.Println("Numbers2 after appending 5 are :", numbers2)
	fmt.Println("Length of numbers2 after appending 5 is :", len(numbers2))
	fmt.Println("Capacity of numbers2 after appending 5 is :", cap(numbers2))

	numbers2 = append(numbers2, 6)
	fmt.Println("Numbers2 after appending 6 are :", numbers2)
	fmt.Println("Length of numbers2 after appending 6 is :", len(numbers2))
	fmt.Println("Capacity of numbers2 after appending 6 is :", cap(numbers2))

	//empty slice
	stock := make([]int, 0)
	fmt.Println("Stock is :", stock)
	fmt.Println("Length of stock is :", len(stock))
	fmt.Println("Capacity of stock is :", cap(stock))

}
