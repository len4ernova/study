package main

import "sync"

var value int
var mutex sync.Mutex

func inc() {
	mutex.Lock()
	value++
	mutex.Unlock()
}

func mutex_local() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			inc()
		}()
	}
	wg.Wait()
}
