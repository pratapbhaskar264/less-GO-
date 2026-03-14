package main 

import "fmt"

func wish() { // simple wish program 
	fmt.Println("hey")
}

func add(a int , b int) (int) { //note the return type format
	return a+b
}

func main() {
	fmt.Println("hey Fuction learning")
	wish()
	// var a int = 3
	// var b int = 4
	// fmt.Println(add(a,b))
	fmt.Println(add(2,3))
}