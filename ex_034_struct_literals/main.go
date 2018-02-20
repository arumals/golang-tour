package main

import "fmt"

type City struct {
	Name       string
	Population int
}

var (
	madrid    = City{"Madrid", 150000}
	barcelona = City{Name: "Barcelona", Population: 100000}
	cadiz     = &City{"Cadiz", 120000}
)

func main() {
	fmt.Println(madrid, barcelona, cadiz)
}
