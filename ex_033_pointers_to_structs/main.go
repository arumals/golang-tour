package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"John Doe", 25} // crea a person instance
	fmt.Println(p)              // print instance
	l := &p                     // create a pointer to instance
	(*l).Name = "Scarlet"       // modify name value
	fmt.Println(p)              // print instance
	l.Name = "Julian"           // use simplified assignment
	fmt.Println(p)              // print instance
}
