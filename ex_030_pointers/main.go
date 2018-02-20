package main

import "fmt"

func main() {
	i := 42
	fmt.Println(i)  // print i value
	fmt.Println(&i) // print i address
	p := &i         // now p is a pointer to i
	fmt.Println(p)  // print p which is a pointer
	fmt.Println(*p) // print pointer value
	*p = 20         // change the pointer value
	fmt.Println(*p) // print p pointer value
	fmt.Println(i)  // print new i value
}
