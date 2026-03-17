package main

import "fmt" 

func main() {
	// fmt.Println("fn check")

   	numbers := []int{1,2,3,4,5}
    
	// habbit One
	fmt.Println("numbers is :",numbers)
	fmt.Printf( "Type of numbers is : %T\n" , numbers )
    fmt.Println(len(numbers))
	numbers = append(numbers ,3,4,4) //append works this way
	fmt.Println(numbers)

}