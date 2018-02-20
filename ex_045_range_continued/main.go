package main

import "fmt"

func main() {
	nums := make([]int, 10)
	fmt.Printf("len: %d cap: %d %v\n", len(nums), cap(nums), nums)
	for i := range nums {
		nums[i] = 1 << uint(i)
	}
	fmt.Printf("len: %d cap: %d %v\n", len(nums), cap(nums), nums)
}
