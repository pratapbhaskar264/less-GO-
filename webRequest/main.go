package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// fmt.Println("Hello, World!")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting data")
		return
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error getting data")
		return
	}

	fmt.Println(string(data))

}
