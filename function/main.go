package main 

import "fmt"


func SimpleFunction() { 
	fmt.Println("simple Function ")
}
func Add( a int , b int ) ( int ) { 
	 return a + b ; 
}

func main () { 
	 fmt.Println("We're learning functions in golang") ;
	 SimpleFunction( ) ; 
	 ans := Add( 1 ,2 ) ;
	 fmt.Println("addition of 1 & 2 : " , ans ) ;
	 
}
