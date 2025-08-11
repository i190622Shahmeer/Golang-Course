package main

import (
	"fmt"
	"time"
)

func printMessage() {
	// for i := 0; ; i++ {
	fmt.Println("Hey World!")
	time.Sleep(2000 * time.Millisecond) // Sleep for 100 milliseconds
	fmt.Println("Hey World again!")
	// }
}

func printMessage2() {

	// for i := 0; ; i++ {
	fmt.Println("Hey World 2!")
	time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
	// }
}

func main() {

	fmt.Println("Learning goroutines in Go!")

	go printMessage()
	go printMessage2()

	time.Sleep(5 * time.Second) // Wait for goroutines to finish

	testwaitgroup()
}
