package main

import (
	"fmt"
	"runtime"
)

func gmpModel() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("CPU (логические ядра):", runtime.NumCPU())
	runtime.GOMAXPROCS(16)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("CPU (логические ядра):", runtime.NumCPU())
}
