package main 

import (
	"fmt"
	// "lessGo/utils"
)

func main() {
	fmt.Println("Hello mam")
	// utils.Greet("Good evening 23 times ")
	var name string = "Hey ! how are you"
	fmt.Println(name)

	var float float64 = 264.45
	fmt.Println(float)
	
	var float2  = 264.45 //type inference
	fmt.Println(float2)

	const truth = "rohit shrama is GOAT" //never changing value
	fmt.Println(truth)
	name = "ok"
	fmt.Println(name)


	other_trick := "hehehaha"
	fmt.Println(other_trick) //wow 
}