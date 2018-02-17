package main

import "fmt"

func main() {
	fmt.Println("counting")
	// defers are put onto a stack, last in first out
	for i := 0; i <= 5; i++ {
		defer fmt.Println("deferer", i)
	}
	fmt.Println("done")
}
