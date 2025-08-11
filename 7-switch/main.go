package main

import "fmt"

func main() {
	month := 6
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	default:
		fmt.Println("Invalid month")

	}

	day := "Thursday"
	// Using switch with multiple cases
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday", "Thursday":
		fmt.Println("Today is Wednesday or Thursday")
	default:
		fmt.Println("Today is not Monday or Tuesday")
	}
}
