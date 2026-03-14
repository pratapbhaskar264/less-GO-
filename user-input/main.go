package main

import ("fmt" 
 		"bufio" 
 		"os" ) 

func main() {
	// fmt.Println("hey")
	// var name string

	// fmt.Scan(&name) // <- input taken but only reads upto white space
	// fmt.Println(name)

	// to prevent whitespace role here we go with buffer reader 
       
	reader := bufio.NewReader(os.Stdin)
	name , _ := reader.ReadString('\n')

	fmt.Println(name)



}