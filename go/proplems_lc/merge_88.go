package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		if len(nums1) == n {
			_ = copy(nums1, nums2)
		}
	}
	i, l1, l2 := len(nums1)-1, m-1, n-1
	for i > 0 && l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[i] = nums1[l1]
			nums1[l1] = 0
			i--
			l1--
			if l1 < 0 {
				_ = copy(nums1[:l2+1], nums2[0:l2+1])
				break
			}

			if i == 0 {
				if l1 > 0 {
					nums1[0] = nums1[l1]
				} else {
					nums1[0] = nums2[l2]
				}
			}
		} else {
			nums1[i] = nums2[l2]

			i--
			l2--
		}
	}

}
