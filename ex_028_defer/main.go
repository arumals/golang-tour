package main

import "fmt"

func main() {
	defer fmt.Println("world") // this will be executed at the end of function
	fmt.Println("hello")
}
