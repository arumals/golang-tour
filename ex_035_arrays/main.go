package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])
	fmt.Println(a[1])
	people := [5]string{"John", "Peter", "Andrew", "Loren", "Tania"}
	fmt.Println(people)
}
