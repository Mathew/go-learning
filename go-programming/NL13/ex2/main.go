package main

import (
	"fmt"
	"github.com/mathew/go-learning/go-programming/NL13/ex2/quote"
	"github.com/mathew/go-learning/go-programming/NL13/ex2/word"
)

func main() {
	fmt.Println(word.UseCount(quote.Sediment))
	fmt.Println(word.Count(quote.Sediment))
}
