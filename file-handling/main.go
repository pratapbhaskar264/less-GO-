package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("heyhey")

	// file, err := os.Create("exe.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("file created")

	// content := "hello world"

	// file.Write([]byte(content))
	// file.Write([]byte(" heyhey"))
	// file.Read([]byte , 10)

	//READ

	file, err := os.Open("exe.txt")

	buffer := make([]byte, 1024)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for {
		f, error := file.Read(buffer)

		if error == io.EOF {
			break
		}

		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println(string(buffer[:f]))

	}

}
