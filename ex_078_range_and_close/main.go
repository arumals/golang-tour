package main

import "fmt"

func main() {
	c := make(chan int, 10)
	// introduce values in the channel
	for i := 0; i < cap(c); i++ {
		c <- i
	}
	close(c)
	// retrieve values from the channel
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
