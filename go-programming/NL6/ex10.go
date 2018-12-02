package main

import "fmt"

func counter() func(number int) int {
	count := 0

	return func(number int) int {
		count += number

		return count
	}
}

func main() {

	countr := counter()
	countr(5)
	countr(5)
	countr(5)
	countr(5)
	result := countr(5)
	fmt.Println(result)
	countr(5)
	countr(5)
	result = countr(5)

	fmt.Println(result)
}
