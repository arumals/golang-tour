package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// single line
	fmt.Println(Person{"Code Guy", 30})
	// multiline
	fmt.Println(Person{
		Name: "John Doe",
		Age:  37,
	})
}
