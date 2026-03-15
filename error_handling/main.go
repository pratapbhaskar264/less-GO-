package main

import "fmt" 

func divide( a , b float64)  (float64 , error) { // error handeled .. easy
	
	if b == 0 {
		return 0 , fmt.Errorf( "brother....brother here .... no 0 allowed")
	} 
	return a / b , nil
}

func main() {
  

	fmt.Println("Function check")
	 
	ans , err := divide( 10 , 0 ) 
	if err == nil {
		fmt.Println(ans)
	} else{
		fmt.Println(err)
	}

}