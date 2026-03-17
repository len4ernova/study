package main

import (
	"fmt"
	"time"
)

func goroutins_index() {
	for i := 0; i < 5; i++ {
		i := i
		//fmt.Println(unsafe.Pointer(&i))

		go func() {
			fmt.Println(i)
		}()

	}
	time.Sleep(2 * time.Second)
}
