package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 int = 36

	fmt.Println(num1)

	var num2 float64 = float64(num1)

	fmt.Println(num2)
	fmt.Printf("Type of %f is %T ", num2, num2)

	str := "100"

	val2, err := strconv.Atoi(str)

	// fmt.Println(str, "ok")
	fmt.Println(err)
	fmt.Println(val2)

	str2 := "45.45"

	floatData, _ := strconv.ParseFloat(str2, 64)

	fmt.Println(floatData)

}
