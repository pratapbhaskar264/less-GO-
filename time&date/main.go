package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time&date")

	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Println(currTime.Date())
	fmt.Println(currTime.Day())
	fmt.Println(currTime.Hour())
	fmt.Println(currTime.Local().Hour())
	fmt.Println(currTime.Minute())
	fmt.Println(currTime.Month())
	fmt.Println(currTime.Year())
	fmt.Println(currTime.YearDay())
	fmt.Printf("%d-%d-%d\n", currTime.Day(), currTime.Month(), currTime.Year())
	fmt.Printf("%T\n", currTime.Day())

	formattedTime := currTime.Format("02 January,2006 03:04:05") // format is fixed and we have to use this format only to get the desired output
	fmt.Println(formattedTime)
	fmt.Printf("%T\n", formattedTime)
}
