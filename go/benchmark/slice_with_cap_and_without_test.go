package main

// go test -bench=. slice_with_cap_and_without_test.go
import "testing"

func BenchmarkWithReservation(b *testing.B) {
	sourceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 0; i < b.N; i++ {
		targetData := make([]int, 0, len(sourceData))
		for j := 0; j < len(sourceData); j++ {
			targetData = append(targetData, sourceData[j])
		}
	}
}

func BenchmarkWithoutReservation(b *testing.B) {
	sourceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 0; i < b.N; i++ {
		targetData := make([]int, 0)
		for j := 0; j < len(sourceData); j++ {
			targetData = append(targetData, sourceData[j])
		}
	}
}
