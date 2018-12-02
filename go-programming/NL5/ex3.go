package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	v := vehicle{
		doors: 4,
		color: "black",
	}

	t := truck{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(v, t, s)
	fmt.Println(v.doors, t.fourWheel, s.luxury)
}
