package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2 * z)
	for z*z > x {
		// with the original calculation leaks at 5
		z = float64(int(z - (z*z-x)/(2*z)))
	}
	return z
}

func main() {
	for i := 0; i <= 1000; i++ {
		fmt.Printf("The closest int sqrt value for %v is %v\n", i, Sqrt(float64(i)))
	}
}
