package main

import (
	"fmt"
)

type Transport interface {
	Move() string
}

type Airplane string

func (a Airplane) Move() string {
	return string(a) + " is flying"
}

type Car string

func (c Car) Move() string {
	return string(c) + " is roading"
}

type Horse string

func (h Horse) Move() string {
	return string(h) + " is stepping"
}

func main() {
	// var type transport
	var t Transport
	t = Airplane("Boeing")
	fmt.Println(t.Move())
	t = Car("Honda")
	fmt.Println(t.Move())
	t = Horse("Paso Fino")
	fmt.Println(t.Move())
	// slice of interfaces
	ts := []Transport{}
	ts = append(ts, Airplane("Jumbojet"))
	ts = append(ts, Car("Mitsubishi"))
	ts = append(ts, Horse("Alazan"))
	fmt.Println(ts)
}
