package main

import "fmt"

func main() {

	c := make(chan int)
	go generateNumbers(c, 100)
	receiveNumbers(c)

}

func generateNumbers(c chan<- int, num int) {
	for i := 0; i < num; i++ {
		c <- i
	}
	close(c)
}

func receiveNumbers(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
