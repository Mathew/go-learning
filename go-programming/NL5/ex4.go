package main

import "fmt"

func main() {

	hero := struct {
		names         []string
		identity      string
		costumeColors map[string]string
	}{
		names:    []string{"Matt", "Murdoch"},
		identity: "DareDevil",
		costumeColors: map[string]string{
			"primary":   "Red",
			"secondary": "Black",
		},
	}

	fmt.Println(hero)

}
