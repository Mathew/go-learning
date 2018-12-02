package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length

}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	crc := circle{
		radius: 10.23,
	}

	sq := square{
		length: 10,
	}

	info(crc)
	info(sq)

}
