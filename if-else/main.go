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
	if age < 16 {
		fmt.Println("get ready")
	} else if (age >= 16 && age <= 17) || age < 18  { // go keeps brackets where required 
		// else it removes .... write only things that are required.
		fmt.Println("alomost ready")
	} else{
		fmt.Println("Can vote")
	}

}