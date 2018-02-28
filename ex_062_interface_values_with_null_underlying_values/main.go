package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil")
	} else {
		fmt.Printf("S value is: %v\n", t.S)
	}
}

func describe(i I) {
	fmt.Printf("%v %T\n", i, i)
	i.M()
}

func main() {
	var i I // interface declared variable (nil at this point)
	// describe(i) // will generate an error because we can't call an interface method directly
	var t *T // pointer to empty struct declared variable (nil at this point)
	i = t    // i now contains an structure
	describe(i)
	var t2 I = &T{"Hello World"}
	describe(t2)
}
