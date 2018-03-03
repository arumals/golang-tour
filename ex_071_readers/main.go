package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("Hello World")
	var b []byte = make([]byte, 1)
	for {
		if n, err := r.Read(b); err == nil {
			fmt.Println(string(b[:n]))
		} else {
			break
		}
	}
}
