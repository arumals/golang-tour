package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	l := len(b)                 // obtain byte length
	for i := 0; i <= l-1; i++ { // walk across the byte slice
		b[i] = byte('A') // append and 'A' converted to byte to it
	}
	return l, nil // return the slice dimension
}

func main() {
	reader.Validate(MyReader{})
}
