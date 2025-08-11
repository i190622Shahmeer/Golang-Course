package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web Req in Go using fake api") //search fake api on google
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close() //defer is used to close the response body after the function returns
	//response type
	fmt.Printf("Response Type: %T\n", res)
	//response status code
	fmt.Println("Response Status Code:", res.StatusCode)
	//read respose (response in json)
	// fmt.Println("Response:", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// Print the response body
	fmt.Println("Response Body:", string(data))
}
