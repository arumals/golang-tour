package main

import "fmt"

type Country struct {
	Name       string
	Population float64
	BirthRate  float64
	DeadRate   float64
	Year       float64
}

func (p *Country) Next() {
	p.Year += 1
	p.Population *= 1 + ((p.BirthRate - p.DeadRate) / 100)
}

func main() {
	c := Country{"Germany", 500000, 1.2, 2.3, 1998}
	fmt.Println(c)
	c.Next()
	fmt.Println(c)
	c.Next()
	fmt.Println(c)
	c.Next()
	fmt.Println(c)
	c.Next()
	fmt.Println(c)
}
