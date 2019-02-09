package word

import (
	"fmt"
	"github.com/mathew/go-learning/go-programming/NL13/ex2/quote"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "expected", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.Sediment)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "Three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}

	}
}
