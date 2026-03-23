package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("heyhey")

	file, err := os.Create("exe.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("file created")

	content := "hello world"

	file.Write([]byte(content))
	file.Write([]byte(" heyhey"))
	// file.Read([]byte , 10)
}
