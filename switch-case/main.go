package main 
 
import ("fmt"
         "bufio"
		 "os"
		  "strings"
		"strconv")

func main(){
	fmt.Println("hey")
   
	//user input
	reader := bufio.NewReader(os.Stdin)
	input , _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value , _ := strconv.Atoi(input)
	
	//switch case
	switch value {
	case 1:
		fmt.Println("one")
	default :
		fmt.Println("Htt jaa tu")
	}
}