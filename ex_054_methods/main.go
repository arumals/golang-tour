package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHi() string {
	return fmt.Sprintf("Hello, my name is %v and I have %v yearls old", p.Name, p.Age)
}

func main() {
	p := Person{"Mary", 23}
	fmt.Println(p.SayHi())
}
