package main

import "fmt"

func fooTwo(numbers ...int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func barTwo(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func main() {
	numbers := []int{4, 5, 6, 7, 8}

	fmt.Println(fooTwo(numbers...))
	fmt.Println(barTwo(numbers))
}
