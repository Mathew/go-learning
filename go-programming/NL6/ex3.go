package main

import "fmt"

func whoAmI() {
	fmt.Println("Batman")
}

func main() {
	defer whoAmI()
	fmt.Print("I am: ")

}
