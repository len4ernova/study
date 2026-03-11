package main

import "fmt"

func main() {
	val := 121
	fmt.Println(int_palindrom(val))
	fmt.Println(int_palindrom(-121))
	fmt.Println(int_palindrom(10))
	fmt.Println(int_palindrom(1))

	fmt.Println(reverseString("abc"))
	fmt.Println(reverseString("АБВ"))

	fmt.Println(palindrom("bob"))
	fmt.Println(palindrom("bobr"))
	fmt.Println(palindrom("boc"))
	fmt.Println(palindrom("boob"))
}
