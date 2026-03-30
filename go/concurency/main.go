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

	//runtime_exit1()

	//mutex_for_order()

	/*fb := NewFooBar(10)
	go fb.Foo(func() { fmt.Println("Foo") })
	fb.Bar(func() { fmt.Println("Bar") })
	*/

	//mutex_local()

	sampleONCE()

}
