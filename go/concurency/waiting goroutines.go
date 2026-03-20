package main

import (
	"fmt"
	"sync"
)

func waitingGoroutines() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 6; i++ {
		defer wg.Done()
		fmt.Println("test")
	}
	wg.Wait()
}
