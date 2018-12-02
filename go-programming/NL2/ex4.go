package main

import "fmt"

func main() {
	a2 := 42
	fmt.Printf("%b %d %#x \n", a2, a2, a2)

	b2 := a2 << 1
	fmt.Printf("%b %d %#x", b2, b2, b2)
}
