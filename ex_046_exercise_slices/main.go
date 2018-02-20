package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y, _ := range s {
		s[y] = make([]uint8, dx)
		for x, _ := range s[y] {
			s[y][x] = uint8((x + y) / 2)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
