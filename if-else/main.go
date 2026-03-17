package main

import ("fmt"
		"bufio"
		"os"
		"strconv"
		"strings")

func main(){
	
	 reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
	age ,_  := strconv.Atoi(input)
	if age < 17 {
		fmt.Println("get ready")
	} else if age >= 17 && age < 18 {
		fmt.Println("alomost ready")
	} else{
		fmt.Println("Can vote")
	}

}