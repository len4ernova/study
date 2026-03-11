package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func FillString() string {
	str := ""

	for i := 0; i < 1000; i++ {
		str += strconv.Itoa(i)
	}

	return str
}
func Addsymb() string {
	bs := strings.Builder{}
	for i := 0; i < 100; i++ {
		bs.WriteString(strconv.Itoa(i))
	}
	return bs.String()
}
func BenchmarkFillString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FillString()
	}
}
func BencchmarAddsymb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addsymb()
	}
}
func main() {
	fmt.Println(Addsymb())
}
