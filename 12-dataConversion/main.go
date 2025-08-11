package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Value of num :", num)
	fmt.Printf("Type of num is :  %T\n", num)
	// Type conversion
	// int to float64
	var data float64 = float64(num)
	fmt.Println("Value of data :", data)
	fmt.Printf("Type of data is :  %T\n", data)

	// int to string
	str := strconv.Itoa(num)
	fmt.Println("Value of str :", str)
	fmt.Printf("Type of str is :  %T\n", str)

	// string to int
	number_str := "1234"
	number_int, _ := strconv.Atoi(number_str)
	fmt.Println("Value of number_int :", number_int)
	fmt.Printf("Type of number_int is :  %T\n", number_int)

	// string to float64
	number_string := "3.14"
	number_float, _ := strconv.ParseFloat(number_string, 64)
	fmt.Println("Value of number_float :", number_float)
	fmt.Printf("Type of number_float is :  %T\n", number_float)
}
