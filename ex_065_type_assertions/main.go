package main

import "fmt"

func main() {
	var i interface{} = "Hello World" // create an empty interface with underlying value
	var j interface{} = 10
	fmt.Printf("(%v, %T)\n", i, i)
	if v, ok := i.(string); ok {
		fmt.Println("string value:", v)
	}
	if v, ok := j.(int); ok {
		fmt.Println("integer value:", v)
	}
}
