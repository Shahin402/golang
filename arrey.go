package main

import "fmt"

func main(){
  //normal arrey declaretion
   var car [5]string

   car[0]="Audi"
   car[1]="Mercedece benz"
   car[2]="BMW"
   car[3]="Ford"
   car[4]="Meclaren GT"
       
   fmt.Println(car[0],car[2])
   fmt.Println("whole value of car arrey:",car)


   //arrey in shorthand method 

   team:=[4]string{"fc barcelona","real madrid fc","liverpool fc","fc byarn munich"}

   fmt.Println(team)
   fmt.Println(team[0])
   fmt.Println(team[2])
}