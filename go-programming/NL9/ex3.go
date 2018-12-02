package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Run with following to show race condition: 'go run --race ex3.go'
func main() {

	const numOfGoroutines = 100
	incrementor := 0

	wg := sync.WaitGroup{}
	wg.Add(numOfGoroutines)

	for x := 0; x < numOfGoroutines; x++ {
		go func() {
			goVal := incrementor
			runtime.Gosched()
			goVal++
			incrementor = goVal
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementor)
}
