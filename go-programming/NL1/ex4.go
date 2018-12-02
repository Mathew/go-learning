package main

import "fmt"

type mat int

var x3 mat

func main() {

	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42

	fmt.Println(x3)

	y := int(x3)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
