package main

import (
	"os"
)

func main() {
	/*
		// Create a new file
		// This will create a file named "file.txt" in the current directory
		// If the file already exists, it will be truncated to zero length
		// If the file does not exist, it will be created
		// The os.Create function returns a file pointer and an error
		// If there is an error, it will panic and stop the program
		file, err := os.Create("file.txt") // Create a file named file.txt
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		fmt.Println("File created successfully:", file.Name())

		content := "Hello World"
		_, err1 := io.WriteString(file, content+"\n") // Write content to the file
		if err1 != nil {
			fmt.Println("Error writing to file:", err1)
			return
		}
		fmt.Println("Content written to file successfully.")
	*/
	// Read the file
	file, error := os.Open("file.txt") // Open the file named file.txt
	if error != nil {
		println("Error opening file:", error)
		return
	}
	defer file.Close()           // Ensure the file is closed after reading
	buffer := make([]byte, 1024) // Create a buffer to hold the file content
	_, err2 := file.Read(buffer)
	if err2 != nil {
		println("Error reading file:", err2)
		return
	}
	println("File content:")
	println(string(buffer))

	//Read the file into byte slice
	content, err3 := os.ReadFile("file.txt") // Read the entire file into a byte slice
	if err3 != nil {
		println("Error reading file:", err3)
		return
	}
	println("File content as byte slice:")
	println(string(content))
}
