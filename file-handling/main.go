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

		// con, err := ioutil.ReadFile("exe.txt")

		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(con))

		//this will not be used with bigger files as it will read the whole
		//  file into memory and can cause out of memory error

		// there fore we will create a buffer and read the file in chunks

	}

}
