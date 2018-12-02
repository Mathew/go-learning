package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Matt Murdoch"
}

func main() {
	jj := person{
		name: "Jessica Jones",
	}
	fmt.Println(jj)
	changeMe(&jj)
	fmt.Println(jj)
}
