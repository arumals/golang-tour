package main

import (
	"fmt"
)

func compute(a, b float64, fn func(float64, float64) float64) float64 {
	return fn(a, b)
}

func main() {
	sum := func(a, b float64) float64 {
		return a + b
	}
	fmt.Printf("%v + %v = %v\n", 120, 140, compute(120, 140, sum))
	mult := func(a, b float64) float64 {
		return a * b
	}
	fmt.Printf("%v * %v = %v\n", 10, 10, compute(10, 10, mult))
}
