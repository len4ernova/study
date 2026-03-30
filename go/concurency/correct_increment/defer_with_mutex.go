package main

import "sync"

var mutex sync.Mutex

func opration() {

}

// обертка для функции - паттерн функционального программирования
func withLock(mutex *sync.Mutex, action func()) {
	if action == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	action()

}
func doSomething1() {
	withLock(&mutex, func() {
		opration()
	})
	opration()
	mutex.Unlock()

}
