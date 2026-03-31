package main

import (
	"fmt"
	"time"
)

func sayhello() {
	fmt.Println("hello1")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("hello2")
}
func sayhii() {
	fmt.Println("hii")
}

func main() {
	fmt.Println("Hello, World!")
	sayhello()
	sayhii()
}
