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

	// defer will be used say file opend , defer and close it in very next line , it'll will be closed at the end of function (reduces chance of failure if we forget to close it later)

}
