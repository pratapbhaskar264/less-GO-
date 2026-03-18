package main 
 
import ("fmt"
         "bufio"
		 "os"
		  "strings"
		"strconv")

func main(){
	fmt.Println("hey")

	reader := bufio.NewReader(os.Stdin)
	input , _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value , _ := strconv.Atoi(input)
	
	//switch case
	switch value {
	case 1 , 3 , 4 :
		fmt.Println("one")
    case 2: 
		fmt.Println("two")
		fmt.Println("ayyehehehehe")
	default :
		fmt.Println("Htt jaa tu")
	}
}