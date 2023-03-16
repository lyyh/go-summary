package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 5.0
	c := float64(a)
	// e := bool(a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
