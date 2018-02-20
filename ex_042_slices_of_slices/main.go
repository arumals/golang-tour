package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{3, 4, 5},
		[]int{2, 3},
		[]int{9, 1, 3, 2},
	}
	for i := 0; i <= len(matrix)-1; i++ {
		for j := 0; j <= len(matrix[i])-1; j++ {
			fmt.Println(matrix[i][j])
		}
		fmt.Println("-----")
	}
}
