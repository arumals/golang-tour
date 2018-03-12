package main

import (
	"fmt"
	"time"
)

func say(m string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(m)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// both functions will run in parallel
	go say("world")
	say("hello")
}
