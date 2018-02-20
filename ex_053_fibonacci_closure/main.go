package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y // x=1, y=1 | x=1, y=2 | x=2, 3 ....
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}
