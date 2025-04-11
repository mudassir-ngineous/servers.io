// goroutine example

// Write a program that spawns 10 goroutines, each printing its number.
// Use a channel to collect results from them.
// Handle potential errors.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (n int) {
			defer wg.Done()
			ch <- n
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println("Goroutine: ", n)
	}
}	
