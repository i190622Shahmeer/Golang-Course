package main

import "fmt"

func main() {

	// Maps in Go
	// Maps are unordered collections of key-value pairs
	// Maps are similar to dictionaries in Python or hash tables in other languages
	// Maps are reference types, so when you assign a map to another variable, both variables point to the same map
	// Maps are created using the make function or using a map literal
	// Maps are mutable, meaning you can add, update, or delete key-value pairs
	//Initializing a map
	marks := make(map[string]int)

	// Adding key-value pairs to the map
	marks["Shahmeer"] = 100
	marks["Bob"] = 90
	marks["Alice"] = 95
	marks["John"] = 85

	// Accessing values from the map
	fmt.Println("Marks of Shahmeer:", marks["Shahmeer"])
	fmt.Println("Marks of Bob:", marks["Bob"])
	fmt.Println("Marks of Alice:", marks["Alice"])
	fmt.Println("Marks of John:", marks["John"])

	//updating a map value
	marks["John"] = 95
	fmt.Println("Updated Marks of John:", marks["John"])

	// Deleting a key-value pair
	delete(marks, "Bob")
	fmt.Println("Marks of Bob after deletion:", marks["Bob"])

	// Checking if a key exists in the map
	student, exists := marks["David"]
	fmt.Println("Marks of David:", student, "Exists:", exists)

	student2, exists := marks["Shahmeer"]
	fmt.Println("Marks of Shahmeer:", student2, "Exists:", exists)

	// Iterating over a map
	// maps are unordered so index is key and value is value in maps

	for index, value := range marks {
		fmt.Println("Index:", index, "Value:", value)
	}

	//assigning value to map in initialization
	Age := map[string]int{
		"Shahmeer": 25,
		"Bob":      30,
		"Alice":    28,
		"John":     35,
	}

	// Accessing values from the map
	fmt.Println("Age of Shahmeer:", Age["Shahmeer"])

}
