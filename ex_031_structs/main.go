package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Person{
		Name: "John Doe",
		Age:  37,
	})
}
