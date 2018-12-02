package main

import "fmt"

type user struct {
	first string
	last  string
	age   int
}

func (u user) speak() {
	fmt.Printf("My name is %v %v and I am %v years old", u.first, u.last, u.age)
}

func main() {
	lc := user{
		first: "Luke",
		last:  "Cage",
		age:   35,
	}

	lc.speak()
}
