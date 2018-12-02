package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with following to show race condition: 'go run --race ex5.go'
func main() {

	const numOfGoroutines = 100
	var incrementor int64

	wg := sync.WaitGroup{}
	wg.Add(numOfGoroutines)

	for x := 0; x < numOfGoroutines; x++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementor)
}
