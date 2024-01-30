package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	wg := sync.WaitGroup{}

	n := 0

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}(&wg)

	wg.Wait()

	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
