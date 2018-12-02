package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)

	fmt.Println("cores: ", runtime.NumCPU())
	fmt.Println("GOES:", runtime.NumGoroutine())

	go func() {
		fmt.Println("GO!")
		wg.Done()
	}()

	go func() {
		fmt.Println("GOGO!")
		wg.Done()
	}()

	fmt.Println("cores: ", runtime.NumCPU())
	fmt.Println("GOES:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("cores: ", runtime.NumCPU())
	fmt.Println("GOES:", runtime.NumGoroutine())
}
