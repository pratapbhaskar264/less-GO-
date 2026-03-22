package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	data := "  hey,how,are,you,guys,doing,hey  "

	parts := strings.Split(data, ",")

	fmt.Println(parts)
	fmt.Println(strings.Count(data, "hey"))

	trim := strings.Trim(data, " ")

	fmt.Println(trim)

	str1 := "Bhaskar"

	str2 := "Pratap"

	con := strings.Join([]string{str1, str2}, " ")

	fmt.Println(con)

}
