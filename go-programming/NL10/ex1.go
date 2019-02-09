package main

import "fmt"

func main() {

	funcLiteralSolve()
	bufferedChannelSolve()
}

func funcLiteralSolve() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferedChannelSolve() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
