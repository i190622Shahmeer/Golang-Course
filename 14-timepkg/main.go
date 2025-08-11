package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Current Time:", currentTime)
	// For time formating it follows golang launch time format (When go was launched)
	fmt.Println("Current Time:", currentTime.Format("2006-01-02 15:04:05"))          // for 24 hour format
	fmt.Println("Current Time:", currentTime.Format("02-01-2006, Monday"))           // for date format with day name
	fmt.Println("Current Time:", currentTime.Format("02-01-2006, Monday,15:04:05"))  // for 24 hour format with day name
	fmt.Println("Current Time:", currentTime.Format("02-01-2006, Monday, 3:04 PM"))  // for am pm 12 hour format with day name
	fmt.Println("Current Time:", currentTime.Format("02-01-2006, Monday, 15:04 PM")) // for am pm 24 hour format with day name
	//type of data
	// fmt.Printf("Type of currentTime: %T\n", currentTime)

	// Parsing a date string to time.Time
	// The layout must match the format of the date string
	// For example, if the date string is "2025-05-28", the layout should be "2006-01-02"
	// The layout string is a reference date in the format "Mon Jan 2 15:04:05 MST 2006"
	//What is parsing?
	//  Parsing is the process of converting a string representation of a date and time into a time.Time object in Go.
	layout_str := "2006-01-02"
	datestr := "2025-05-28"
	formattedTime, _ := time.Parse(layout_str, datestr)
	fmt.Println("Formatted Time:", formattedTime)

	// add 1 more day to the current time
	new_date := currentTime.AddDate(0, 0, 1) // AddDate(years, months, days)
	fmt.Println("New Date after adding 1 day:", new_date.Format("2006-01-02 15:04:05"))
}
