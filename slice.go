package main

import "fmt"

func main() {
	// normal slice declaration and initialization
	var car []string

	car = append(car, "audi", "BMW", "mercedece benz", "ford")

	fmt.Println(car)
	fmt.Println(car[0])
	// shorthand slice
	web := []string{"html", "css", "javascript", "bootstrap", "recat.js", "node.js"}
	fmt.Println(web)
	fmt.Println(web[4], web[5])

}
