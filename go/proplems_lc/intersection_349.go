package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}

	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 0
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			if m[v] == 0 {
				res = append(res, v)
				m[v]++
			}
		}

	}
	return res
}
