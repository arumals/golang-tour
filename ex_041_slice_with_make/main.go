package main

import "fmt"

func main() {
	colors := make([]string, 5) // create a dynamic array of 5
	fmt.Println("len=%d cap=%d %v\n", len(colors), cap(colors), colors)
	colors = colors[:cap(colors)]
	fmt.Println("len=%d cap=%d %v\n", len(colors), cap(colors), colors)
	colors = colors[1:] // reduce one from left
	fmt.Println("len=%d cap=%d %v\n", len(colors), cap(colors), colors)
}
