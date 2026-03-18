package main

import "fmt"

func main() {
	fmt.Println("function check")

	// var i int = 0

	// for  ; i < 10 ; i++  { // intitalization 
	// 	fmt.Println(i)
	// }

	number := []int{1,2,3,4}
    fmt.Println(number)
	for index , value := range number{
		fmt.Printf(" %d %d \n" , index , value)
	}

	var name string = "sat"
	for index , value := range name{
		fmt.Printf(" %d %c \n" , index , value)
	}
}