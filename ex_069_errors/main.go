package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time // when the error was generated
	What string    // error details
}

func (e *MyError) Error() string { // implicit interface
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error { // returns error interface
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
