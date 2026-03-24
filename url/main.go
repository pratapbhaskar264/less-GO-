package main

import (
	"fmt"
	"net/url"
	// "io"
	// "net/http"
)

func main() {
	fmt.Println("URL FILE : ")

	myURL := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("cannot be parsed")
	}

	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)

	parsedURL.Host = "lalalala" // modifying

	newURl := parsedURL.String() // creating url

	parsedURL2, err := url.Parse(newURl)

	if err != nil {
		fmt.Println("cannot be parsed")
	}

	fmt.Println(parsedURL2.Host)
	fmt.Println(parsedURL2.Scheme)
	fmt.Println(parsedURL2.Path)
	fmt.Println(parsedURL2.RawQuery)

}
