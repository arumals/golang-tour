package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't resolve sqrt on %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	if v, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sqrt 2 = %f\n", v)
	}
	if v, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sqrt -2 = %f\n", v)
	}
}
