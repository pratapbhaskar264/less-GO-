package main

import "fmt" 

func main(){
	name := "alice"
	marks := 67

	// fmt.Println("name :" , name , " , marks: " , marks) // changes line and add spaces req
	// fmt.Printf("name :" , name , " , marks: \n" , marks) // works with format specifiers 
	// fmt.Printf("Name is : %s\n " , name)
	// fmt.Printf("Type of name is %T\n" , name) // type check %T

	fmt.Printf( "Name is : %s\nMarks are : %d " , name , marks ) // multiple in same line 
}