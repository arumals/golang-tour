package main

import (
	"fmt"
)

type Figure interface {
	Area() float64
}

func Area(f Figure) float64 {
	return f.Area()
}

type Circle float64

func (c Circle) Area() float64 {
	return 3.1416 * (float64(c) * float64(c))
}

type Square float64

func (s Square) Area() float64 {
	return float64(s) * float64(s)
}

type Rectangle struct {
	SideA float64
	SideB float64
}

func main() {
	c := Circle(20)
	s := Square(20)
	r := Rectangle{30, 10}
	fmt.Println(Area(c))
	fmt.Println(Area(s))
	fmt.Println(r)
	// fmt.Println(Area(r))
}
