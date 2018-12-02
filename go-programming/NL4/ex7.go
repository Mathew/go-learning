package main

import "fmt"

func main() {
	people := [][]string{
		{"First", "Last", "Phrase"},
		{"1st", "2nd", "Boo"},
	}

	fmt.Println(people)

	for _, person := range people {
		for _, attribute := range person {
			println(attribute)
		}
	}
}
