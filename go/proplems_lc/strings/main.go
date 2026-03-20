package main

import "fmt"

func main() {
	s := []byte("Hello")
	reverseString(s)
	fmt.Println(string(s))
}
