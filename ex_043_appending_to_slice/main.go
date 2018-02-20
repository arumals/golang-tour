package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 100)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 120)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 130, 140, 150)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
