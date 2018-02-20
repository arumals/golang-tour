package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	f := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		f[word] += 1
	}
	return f
}

func main() {
	wc.Test(WordCount)
}
