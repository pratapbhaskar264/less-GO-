package main 

import "fmt"

type student struct {
	name string
	marks int
}

func main() {
  fmt.Println("Structs in Go")
  
  var student1 student
  student1.name = "bhaskar"
  student1.marks = 99
  fmt.Println(student1)

  student2 := student {
	  name : "kunal",
	  marks : 99,
	}

  fmt.Println(student2)

 // 3rd method
//   var person2 = new(Student)
//   person2. .....

// we can create an struct of structs as well
}