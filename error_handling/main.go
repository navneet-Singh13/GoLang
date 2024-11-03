package main 


import "fmt"
func divide(a float64 , b float64 ) (float64 , string ){ 
	 if ( b == 0 ) { 
		return 0  , "denominator must not be 0 " 
	 }
	 return a / b  , "nil"  ; 
}
func main ( ) { 
	// fmt.Println("Error Handling in go! ") ; 
	// ans , _ := divide( 10 , 0 ) ;
	// fmt.Println("Division of Two Numbers is : " , ans  )
	ans , _ := divide(10 , 0) ; 
	fmt.Println("Divsion : " , ans  ) ; 

}
