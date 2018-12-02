package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Run with following to show race condition: 'go run --race ex4.go'
func main() {

	const numOfGoroutines = 100
	incrementor := 0

	wg := sync.WaitGroup{}
	wg.Add(numOfGoroutines)

	mutex := sync.Mutex{}

	for x := 0; x < numOfGoroutines; x++ {
		go func() {
			mutex.Lock()
			goVal := incrementor
			runtime.Gosched()
			goVal++
			incrementor = goVal
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementor)
}
