package main 

import "fmt"

func main() {
	// fmt.Println("function check")
    
	//method1 , 0-based indexing
	var names[5] string

	names[0] = "hey"
	names[4] = "byee"

	fmt.Printf( "%q\n" , names )
    
	var arr = [5]int{0,1,2,3,4}

	fmt.Println(arr)
	
	fmt.Println(len(arr)) // length of array 

	//default value 
    var price[5] int
    var val[5] string  // use printf for quoted string
    var boolean[5] bool

	fmt.Printf("empty quoted val = %q\n" , val  )
	fmt.Println(price)
	fmt.Println(val)
	fmt.Println(boolean)

}