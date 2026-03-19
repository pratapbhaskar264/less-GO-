package main 

import "fmt"

// key value pair of unordered data 

func main(){
	// fmt.Println("function check")
	student := make(map[string]int)

	student["bhaskar"] = 99
	student["chomu"] = 95
	student["kunal"] = 95

	fmt.Println(student)
	delete(student , "chomu")
	value , existance := student["chomu"]
	if !existance  {
		fmt.Println("chomu not found")
	} else{
		fmt.Println(value)
	}
	fmt.Println(student["kunal"])
	fmt.Println(student)
}