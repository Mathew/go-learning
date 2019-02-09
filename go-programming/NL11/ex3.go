package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
	// example of type-casting / asserting: e.(customErr).info
	fmt.Println(e, e.(customErr).info)
}

func main() {
	foo(customErr{info: "There was a problem."})
}
