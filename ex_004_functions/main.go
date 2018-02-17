package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("The sum of %v and %v is %v\n", 5, 5, add(5, 5))
}
