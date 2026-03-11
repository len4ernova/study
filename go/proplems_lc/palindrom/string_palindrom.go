package main

import "strings"

func reverseString(s string) string {
	rs := []rune(s)

	var newS strings.Builder
	for k := len(rs) - 1; k >= 0; k-- {
		newS.WriteString(string(rs[k]))
	}

	return newS.String()
}

func palindrom(s string) bool {
	if len(s) == 1 || s == "" {
		return true
	}

	temps := []rune(s)

	for i, j := 0, len(s)-1; i < len(temps)/2; i, j = i+1, j-1 {
		if temps[i] != temps[j] {
			return false
		}
	}
	return true

}
