package main

import "fmt"

func main() {

	// map declaration
	country := make(map[string]string)

	// map initialization
	country["canada"] = "North America"
	country["USA"] = "North America"
	country["UK"] = "Europ"
	country["AUS"] = "Australia"
	country["India"] = "Asia"
	// map print/data pull/ data get
	fmt.Println("\n", country)
	fmt.Println("\n", country["canada"])
	fmt.Println("\n", country["USA"])
	fmt.Println("\n", country["UK"])

	// data delete from a map

	delete(country, "India")

}
