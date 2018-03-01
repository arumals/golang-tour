package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	fmt.Printf("%v %T", i, i) // nil nil
	// i.M() // runtime error (can't call an interface method directly)
}
