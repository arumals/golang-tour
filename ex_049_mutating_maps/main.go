package main

import "fmt"

func main() {
	capitals := make(map[string]string)
	capitals["US"] = "DC"
	capitals["Mexico"] = "Mexico City"
	capitals["Peru"] = "Lima"
	capitals["Spain"] = "Madrid"
	capitals["France"] = "Paris"
	fmt.Println(capitals)
	if capital, exists := capitals["Peru"]; exists { // verify existence and print it
		fmt.Printf("The capital of %s is %s\n", "Peru", capital)
	}
	delete(capitals, "Peru") // remove element from map
	if _, exists := capitals["Peru"]; !exists {
		fmt.Println("There isn't such country")
	}
}
