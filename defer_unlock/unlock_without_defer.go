package main

import (
	"fmt"
	"sync"
)

func Test() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mutex.Lock() // will wait here indefinitely if Goroutine 2 acquires lock first

		fmt.Println("Goroutine 1 acquired the lock")
		n := 2
		if n%2 == 0 {
			return
		}
		mutex.Unlock() // mutex is never unlocked

	}()
	go func() {
		defer wg.Done()
		mutex.Lock() // will wait here indefinitely if Goroutine 1 acquires lock first

		fmt.Println("Goroutine 2 acquired the lock")

		n := 2
		if n%2 == 0 {
			return
		}
		mutex.Unlock() // mutex is never unlocked

	}()
	wg.Wait()
	fmt.Println("Main goroutine completed")
}
