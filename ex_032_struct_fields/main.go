package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"John Doe", 25}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
