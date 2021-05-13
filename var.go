package main

import "fmt"

func main(){
 
	var name string             //normal variable diclaration
	var city string
	name = "shahin alam"
    city= "chittagong"

	fmt.Println(name,"\n",city)

	age, id:= 24, 8   //shorthand variable diclaration
	fmt.Println("age:",age,)
	fmt.Println("ID:",id,)
   
}