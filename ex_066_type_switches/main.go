package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
