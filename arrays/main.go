package main 

import "fmt"

func main() {
	// fmt.Println("function check")
    
	//method1 , 0-based indexing
	var names[5] string

	names[0] = "hey"
	names[2] = "byee"

	fmt.Println( names )
    
	var arr = [5]int{0,1,2,3,4}

	fmt.Println(arr)


}