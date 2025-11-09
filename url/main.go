package main

import (
	"fmt"
	"net/url"

)

func main(){
	fmt.Println("kar diya")
	myurl := "https://www.example.com/path?query=golang#section"
	fmt.Println("URL is:",myurl)

	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("Parsed URL: %+v\n", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
}