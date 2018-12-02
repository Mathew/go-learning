package main

import "fmt"

var x2 = 42
var y2 = "James Bond"
var z2 = true

func main() {
	s := fmt.Sprintf("%v\t %v\t %v\t", x2, y2, z2)
	println(s)
}
