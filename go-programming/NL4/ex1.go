package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for i, v := range a {
		println(i, v)
	}
	fmt.Printf("%T", a)
}