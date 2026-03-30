package main

import (
	"fmt"
	"sync"
)

var mutexGran sync.Mutex
var cache map[string]string

func doSomething() {
	mutexGran.Lock()
	item := cache["key"]
	mutex.Unlock()
	fmt.Println(item)

}
