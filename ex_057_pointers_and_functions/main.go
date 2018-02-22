package main

import "fmt"

type Circle struct {
	Radio float64
}

func (c *Circle) IncreaseRadio(inc float64) {
	c.Radio += inc
}

func main() {
	c := Circle{10}
	c.IncreaseRadio(5)
	fmt.Println(c)
	c.IncreaseRadio(2)
	fmt.Println(c)
	c.IncreaseRadio(10)
	fmt.Println(c)
}
