package main

import "fmt"

type person2 struct {
	firstName         string
	lastName          string
	favouriteFlavours []string
}

func main() {
	heroes := map[string]person2{
		"Murdoch": {
			firstName:         "Matt",
			lastName:          "Murdoch",
			favouriteFlavours: []string{"Red", "Black"},
		},
		"Jones": {
			firstName:         "Jessica",
			lastName:          "Jones",
			favouriteFlavours: []string{"Blue", "Black"},
		},
	}

	for _, p := range heroes {
		fmt.Printf("%v %v\t", p.firstName, p.lastName)

		for _, icf := range p.favouriteFlavours {
			fmt.Printf("%v,", icf)

		}
		fmt.Println("")
	}

}
