package main

import "fmt"

func main() {
	// 4
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 5, 6}
	//tmp := make([]int, len(nums1)+len(nums2))
	//tmp = mergeArrs(tmp, num1, nums2)
	fmt.Println("res: ", findMedianSortedArrays(nums1, nums2))

	nums1 = []int{0}
	nums2 = []int{1}
	//tmp = make([]int, len(nums1)+len(nums2))
	//tmp = mergeArrs(tmp, nums1, nums2)
	fmt.Println("res: ", findMedianSortedArrays(nums1, nums2))

	nums1 = []int{}
	nums2 = []int{1}
	tmp := make([]int, len(nums1)+len(nums2))
	tmp = mergeArrs(tmp, nums1, nums2)
	fmt.Println("tmp = ", tmp)
	fmt.Println("res: ", findMedianSortedArrays(nums1, nums2))

	// 88
	// num1 := []int{1, 2, 3, 0, 0, 0}
	// merge(num1, 3, []int{2, 5, 6}, 3)
	// fmt.Println(num1)

	// num1 = []int{0}
	// merge(num1, 0, []int{1}, 1)
	// fmt.Println(num1)

	// num1 = []int{2, 0}
	// merge(num1, 1, []int{1}, 1)
	// fmt.Println(num1)

	// 349
	//fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
