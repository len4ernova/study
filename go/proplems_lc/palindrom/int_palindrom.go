package main

func int_palindrom(x int) bool {
	var xx []int
	for {
		xx = append(xx, x%10)
		x = x / 10

		if x < 10 {
			xx = append(xx, x)
			break
		}
	}
	for i, j := 0, len(xx)-1; i < len(xx)/2; i, j = i+1, j-1 {
		if xx[i] != xx[j] {
			return false
		}
	}
	return true
}
