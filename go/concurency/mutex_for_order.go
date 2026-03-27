package main

import (
	"log"
	"sync"
	"time"
)

func mutex_for_order() {
	var mutex sync.Mutex

	mutex.Lock()
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Hi")
		mutex.Unlock()
	}()

	mutex.Lock()
	log.Println("Buy")

}
