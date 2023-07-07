package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println("Goroutine 1 acquired the lock")
		n := 2
		if n%2 == 0 {
			return
		}

	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println("Goroutine 2 acquired the lock")

		n := 2
		if n%2 == 0 {
			return
		}

	}()
	wg.Wait()
	fmt.Println("Main goroutine completed")
}
