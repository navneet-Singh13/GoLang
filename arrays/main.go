package main 

import "fmt"

func main ( ) { 

	 var name[5]string 
	 name[0] = "price"
	 name[2] = "akash"
     var numbers = [5]int{ 1 ,2 ,3 , 4 , 5} ; 

	 fmt.Println("Names : " , name ) ; 
	 fmt.Println("Numbers : " , numbers ) ; 
	 fmt.Println(len(numbers)) ; 
	 fmt.Println("value of name at 2nd idx  :" , name[2]) ;
	  
}