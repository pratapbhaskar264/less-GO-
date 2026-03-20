package main 

import "fmt"

type Student struct {
	name string
	marks int
}

func main() {
  fmt.Println("Structs in Go")
  
  var student1 Student
  student1.name = "bhaskar"
  student1.marks = 99
  fmt.Println(student1)
}