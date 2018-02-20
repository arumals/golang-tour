package main

import "fmt"

type Location struct {
	Latitude  float64
	Longitude float64
}

func main() {
	m := make(map[string]Location)
	fmt.Printf("len=%d %v\n", len(m), m)
	m["Torre Latina"] = Location{19.4339, 99.1406}
	fmt.Printf("len=%d %v\n", len(m), m)
	m["Machu Pichu"] = Location{13.1631, 72.5450}
	fmt.Printf("len=%d %v\n", len(m), m)
}
