package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling in go language")
	//"_" is used to ignore the error or unused value
	// ans, _ := divide(10, 0)
	// fmt.Println("Division of two numbers is:", ans)

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Division of two numbers is:", ans)
}
