package main

import "fmt"

func main() {
	people := [5]string{
		"Peter",
		"Mary",
		"Ellen",
		"Tom",
		"Dory",
	}
	fmt.Println(people)
	b1 := people[:2] // takes 0,1 (final is not included)
	b2 := people[2:] // takes 2, 3, 4
	fmt.Println(b1)
	fmt.Println(b2)
}
