package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("My name is: ", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p := person{
		name: "Matt Murdoch",
		age:  32,
	}

	saySomething(&p)

}
