package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func peformGetRequest() {
	// *************GET Operation************

	// Using the net/http package to make a GET request
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code is not OK:", res.StatusCode)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Response Data:", string(data))
	var todo Todo
	// Decode the JSON response into the Todo struct
	// Use json.NewDecoder to decode the response body directly into the struct
	err = json.NewDecoder(res.Body).Decode(&todo)
	// fmt.Printf("Todo Item: %+v\n", todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo Item Title:", todo)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Shahmeer",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	//converting jsonData to string
	jsonString := string(jsonData)

	//converting string data to reader
	reader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"
	// fmt.Println("JSON String:", jsonString)
	res, err := http.Post(myURL, "application/json", reader)

	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Data:", string(data))
	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error: Status code is not Created:", res.StatusCode)
		return
	}
	fmt.Println("POST request successful. Status code:", res.StatusCode)

}

func performUpdateRequest() {
	// This function will be used to perform an update operation
	// You can implement it similarly to the POST request, but using http.NewRequest with method "PUT"

	todo := Todo{
		UserID:    23344,
		Title:     "Hello Shahmeer",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request
	req, err := http.NewRequest("PUT", myurl, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Data:", string(data))
	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error: Status code is not Created:", res.StatusCode)
		return
	}
	fmt.Println("POST request successful. Status code:", res.StatusCode)
}

func performDeleteRequest() {
	// This function will be used to perform a delete operation
	// You can implement it similarly to the GET request, but using http.NewRequest with method "DELETE"

	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest("DELETE", myurl, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code is not OK:", res.StatusCode)
		return
	}
	fmt.Println("DELETE request successful. Status code:", res.StatusCode)

}

func main() {
	fmt.Println("CRUD Operations in Go")
	// peformGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
