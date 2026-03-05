package main

import (
	"fmt"
)

// func findCount(s string) int {
// 	symbs := make(map[rune]int)
// 	count := 0
// 	for v := range s {
// 		_, ok := symbs[rune(s[v])]
// 		if ok {
// 			break
// 		}
// 		count++
// 		symbs[rune(s[v])] = 1
// 	}
// 	return count
// }
// func lengthOfLongestSubstring(s string) int {

// 	count := 0
// 	for b := range len(s) {
// 		c := findCount(s[b:])
// 		if c > count {
// 			count = c
// 		}

// 	}

// 	return count
// }

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 {
		if rune(s[0]) == rune(s[1]) {
			return 1
		} else {
			return 2
		}
	}
	l, r := 0, 0
	count := 0
	set := make(map[rune]int)

	for r < len(s) {
		// fmt.Println("l ", l, "r ", r)
		_, ok := set[rune(s[r])]

		if ok {
			count = max(count, len(set))

			templ := set[rune(s[r])] + 1

			if templ-l == 1 {
				l++
			} else {
				l = templ
				for k, v := range set {
					if v < l {
						delete(set, rune(k))
					}
				}
			}
		}

		set[rune(s[r])] = r

		// fmt.Println("set = ", set)
		r++
		if r == len(s) {
			count = max(count, len(set))
		}
	}

	return count

}

func main() {
	s := "abcabcdbb"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = " "
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = ""
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
	s = "a"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = "aba"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
	s = "aab"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
	s = "cdd"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }
