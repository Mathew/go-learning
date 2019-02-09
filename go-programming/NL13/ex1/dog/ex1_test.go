package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

// ExampleYears is a demo of how to put an Example in docs (but also tested).
func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	if Years(10) != 70 {
		t.Error("Years does not equal 70")
	}
}

// go test --bench <- Run the benchmark tests.
// go test --cover <- generates coverage on the command line.
// go test --coverprofile c.out <- Generate coverage file.
// go tool cover -html=c.out <- View covered code in a browser.
