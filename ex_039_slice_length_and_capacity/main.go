package main

import "fmt"

func main() {
	movies := []string{"Forres Gump", "Top Gun", "Karate Kid", "Lord Of the Rings"}
	movies = movies[:0] // skip first two
	fmt.Printf("len=%d,cap=%d %v\n", len(movies), cap(movies), movies)
	movies = movies[:3] // expand to 3 elements
	fmt.Printf("len=%d cap=%d %v\n", len(movies), cap(movies), movies)
	movies = movies[1:] // reduce from the left
	fmt.Printf("len=%d cap=%d %v\n", len(movies), cap(movies), movies)
}
