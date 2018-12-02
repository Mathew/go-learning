package main

import "fmt"

func main() {
	people := map[string][]string{
		"murdoch_matt":  {"world", "on", "fire"},
		"cage_luke":     {"almost", "is", "invincible"},
		"jones_jessica": {"super", "strong", "investigator"},
	}

	people["fist_iron"] = []string{"What's", "his", "name?"}

	for name, facts := range people {

		fmt.Printf("%v\t", name)

		for i, fact := range facts {
			fmt.Printf("%v %v\t", i, fact)
		}

		fmt.Println("")
	}

	delete(people, "fist_iron")
	fmt.Println(people)

}
