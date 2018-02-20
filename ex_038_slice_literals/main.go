package main

import "fmt"

func main() {
	countries := []struct {
		Name      string
		Continent string
	}{
		{"Spain", "Europe"},
		{"Mexico", "America"},
		{"Australia", "Oceania"},
	}
	fmt.Println(countries)
}
