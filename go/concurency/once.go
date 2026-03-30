package main

import (
	"fmt"
	"sync"
)

func sampleONCE() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(x int) {
			defer wg.Done()
			once.Do(onceBody)
			fmt.Println(x)
		}(i)

	}
	wg.Wait()
}
