package main

import "fmt"



func add(a , b *int) int {
	return (*a)+(*b);
}

func modify(b *int)  {
	*b = *b + 1 // value changed at address
}

func main() {
	fmt.Println("Pointers in Go")

	var num1 int = 36

	var num2 *int = &num1

	fmt.Println(num1)
	fmt.Println("value at address of nums1 :" , *num2)
	fmt.Println("address of num2 :",&num2)
	fmt.Println("address of num1 :",num2)

	val := 45 
	ptr := &val

	fmt.Println(val)
	fmt.Println(*ptr)
    
	fmt.Println(add(&val , ptr ))

	a := 6

	b := 3

    fmt.Println(add(&a , &b))

	hehe := 68 

	modify(&hehe)

	fmt.Println(hehe)
}