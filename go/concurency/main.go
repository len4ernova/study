package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Кол-во горутин", runtime.NumGoroutine())
	//server_with_panic()

	//mainNeverExit()

	//goroutins_index()

	//gmpModel()
	//asyncPreemtible()
	runtime_exit1()
}
