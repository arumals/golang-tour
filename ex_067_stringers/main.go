package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // here we are implementing the Stringer implicit interface
	return fmt.Sprintf("%v is %v", p.Name, p.Age)
}

func describe(s fmt.Stringer) {
	fmt.Println(s)
}

func main() {
	a := Person{"John Doe", 50}
	describe(a) // it attempts to conver the struct to an string
}
