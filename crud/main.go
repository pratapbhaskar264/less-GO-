package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	data, err := http.Get("https://dummyjson.com/todos/45")

	if err != nil {
		fmt.Println("error finding todosss")
		return
	}

	defer data.Body.Close()
	orData, err := io.ReadAll(data.Body)

	if err != nil {
		fmt.Println("error finding todosss")
		return
	}

	fmt.Println(string(orData))

}
