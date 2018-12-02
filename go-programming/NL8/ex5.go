package main

import (
	"fmt"
	"sort"
)

type userThree struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []userThree

type byLast []userThree

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAge) Less(i int, j int) bool {
	return b[i].Age < b[j].Age
}

func (b byLast) Len() int {
	return len(b)
}

func (b byLast) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLast) Less(i int, j int) bool {
	return b[i].Last < b[j].Last
}

func main() {
	u1 := userThree{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := userThree{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := userThree{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []userThree{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(byAge(users))
	fmt.Println(users)
	sort.Sort(byLast(users))
	fmt.Println(users)
}
