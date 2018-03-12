package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	// calculate sum of 2 slices in 2 routines simultaneously
	s := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	go sum(s[:2], c)
	go sum(s[2:], c)
	x := <-c
	y := <-c
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
}
