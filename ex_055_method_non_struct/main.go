package main

import (
	"fmt"
	"math"
)

type CustomFloat float64

func (c CustomFloat) Pow(exp float64) float64 {
	return math.Pow(float64(c), exp)
}

func main() {
	cf := CustomFloat(2)
	fmt.Println(cf.Pow(3))
}
