package main

import (
	"log"
	"runtime"
	"time"
)

func runtime_exit1() {
	go func() {
		log.Println("first")
		runtime.Goexit() // принудительное завершение
		log.Println("second")
	}()

	time.Sleep(3 * time.Second)
}

// runtime.Goexit() вызванный для main не завершит горутину
func runtime_exit2() {
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println("first")

		}

	}()

	time.Sleep(3 * time.Second)
	runtime.Goexit() // принудительное завершение
}
