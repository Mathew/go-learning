package main

import "fmt"

func main() {

	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for _, v := range s {
		println(v)
	}

	fmt.Printf("%T", s)

}
