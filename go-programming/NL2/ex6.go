package main

const (
	_  = iota
	y1 = 2017 + iota
	y2 = 2017 + iota
	y3 = 2017 + iota
	y4 = 2017 + iota
)

func main() {
	println(y1, y2, y3, y4)
}
