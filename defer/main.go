package main

import "fmt"

func add(a, b int) (res int) {
	res = a + b
	return
}
func main() {
	defer fmt.Println("world") // basically it pushes everything to be executed just after every thing without defer in the fn is executed
	//more than one defer are excecuted in stack order .... lifo
	defer fmt.Println(add(2, 3))
	fmt.Println("hello")
}
