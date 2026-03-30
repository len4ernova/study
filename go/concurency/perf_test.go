package main

import (
	"sync"
	"testing"
)

// go test -bench=. perf_test.go
func BenchMarkMutexWithDefer(b *testing.B) {
	var counter int64
	var mutex sync.Mutex
	for j := 0; j < b.N; j++ {
		func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
		}()
	}
}

func BenchMarkMutexWithoutDefer(b *testing.B) {
	var counter int64
	var mutex sync.Mutex
	for j := 0; j < b.N; j++ {
		func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
}
