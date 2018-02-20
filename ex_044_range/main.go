package main

import "fmt"

func main() {
	cars := []string{"Camaro", "Mustang", "Bettle", "Corvette", "Monza", "Tracker"}
	for i, name := range cars {
		fmt.Printf("Car %d: %v\n", i, name)
	}
}
