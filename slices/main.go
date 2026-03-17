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
    
	names := []string{};
	names = append(names , "kunal" , "bhaskar" ,"rohitsharma" , "aditya")
	fmt.Println(names)

	//habbit 2nd where we use make function to set length and capacity 
    
	list := make( []int , 2 , 5 )
	list = append(list , 1,2,3,5,4,4,4,4,4,4,4)
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(cap(list))

}