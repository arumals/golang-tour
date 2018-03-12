package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "David"
	ch <- "Peter"
	ch <- "George"
	// ch <- "Steve" // this will provoke a fatal error
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
