package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL Handling in Go")
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	convertedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Parsed URL:", convertedURL)
	fmt.Println("Scheme:", convertedURL.Scheme)
	fmt.Println("Host:", convertedURL.Host)
	fmt.Println("Path:", convertedURL.Path)
	fmt.Println("Raw Query:", convertedURL.RawQuery)

	//modifying URL
	convertedURL.Scheme = "http"
	convertedURL.Host = "jsonplaceholder.typicode.com"
	convertedURL.Path = "/todos/2"
	convertedURL.RawQuery = "userId=1"
	fmt.Println("Modified URL:", convertedURL.String())
	fmt.Println("Parsed URL:", convertedURL)
	fmt.Println("Scheme:", convertedURL.Scheme)
	fmt.Println("Host:", convertedURL.Host)
	fmt.Println("Path:", convertedURL.Path)
	fmt.Println("Raw Query:", convertedURL.RawQuery)

}
