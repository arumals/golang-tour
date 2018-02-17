package main

import "fmt"

func split(a, b string) (x, y string) {
	x, y = b, a
	return
}

func main() {
	fmt.Println(split("Hello", "World"))
}
