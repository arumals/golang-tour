package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// switch can be used as a cleaner if else
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
