package main

import "fmt"

type Location struct {
	Latitude  float64
	Longitude float64
}

func main() {
	m := map[string]Location{
		"Torre Latina": Location{19.4339, 99.1406},
		"Machu Pichu":  Location{13.1631, 72.5450},
	}
	fmt.Println(m)
}
