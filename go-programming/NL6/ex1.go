package main

import "fmt"

func foo() int {
	return 5
}

func bar() (int, string) {
	return 10, "Hello"
}

func main() {
	foor := foo()
	barrI, barrS := bar()

	fmt.Println(foor)
	fmt.Println(barrI)
	fmt.Println(barrS)
}
