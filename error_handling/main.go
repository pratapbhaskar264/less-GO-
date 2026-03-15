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
	 
	// ans , err := divide( 10 , 0 ) 
	// if err == nil {
	// 	fmt.Println(ans)
	// } else{
	// 	fmt.Println(err)
	// }

	//avoid all this and use _ evry where you need to ignore the usecase -> 
	// of something being returned
    // only reason to ignore and use _ is because if declared var another ->
	// -> than _ will not be used then it will show error  (golang rizz)
    
	ans , _ := divide( 10 , 4 ) 
	fmt.Println(ans)


}