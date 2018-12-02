package main

import "fmt"

func main() {
	f := func(name string) string {
		return fmt.Sprint("Hello ", name)
	}

	fmt.Print(f("Luke"))
}
