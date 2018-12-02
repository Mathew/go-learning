package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteFlavours []string
}

func main() {
	mm := person{
		firstName:         "Matt",
		lastName:          "Murdoch",
		favouriteFlavours: []string{"Red", "Black"},
	}

	jj := person{
		firstName:         "Jessica",
		lastName:          "Jones",
		favouriteFlavours: []string{"Blue", "Black"},
	}

	for _, p := range []person{mm, jj} {
		fmt.Printf("%v %v\t", p.firstName, p.lastName)

		for _, icf := range p.favouriteFlavours {
			fmt.Printf("%v,", icf)

		}
		fmt.Println("")
	}

}
