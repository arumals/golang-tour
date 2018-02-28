package main

import (
	"fmt"
)

type Dimensionable interface {
	Count() int
}

type Numbers []int

func (n Numbers) Count() int {
	return len(n)
}

type Name struct {
	First string
	Last  string
}

func (n *Name) Count() int {
	return len(n.First) + len(n.Last)
}

func main() {
	var nums Dimensionable = Numbers{5, 10, 15, 20, 25, 30}
	fmt.Println(nums.Count())         // print count of numbers
	fmt.Printf("%v %T\n", nums, nums) // print tuple values [value, type]
	var name Dimensionable = &Name{"John", "Doe"}
	fmt.Println(name.Count())         // print full name size
	fmt.Printf("%v %T\n", name, name) // print tuple values [value, type]
}
