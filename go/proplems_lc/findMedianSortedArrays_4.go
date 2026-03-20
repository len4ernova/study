package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tmp []int
	if len(nums1) == 0 {
		tmp = make([]int, len(nums2))
		_ = copy(tmp, nums2)
	} else if len(nums2) == 0 {
		tmp = make([]int, len(nums1))
		_ = copy(tmp, nums1)
	} else {
		tmp = make([]int, len(nums1)+len(nums2))

		tmp = mergeArrs(tmp, nums1, nums2)
	}

	if len(tmp)&1 == 1 {
		return float64(tmp[len(tmp)/2])
	} else {
		res := float64(tmp[len(tmp)/2-1]+tmp[len(tmp)/2]) / 2
		return res
	}
}

func mergeArrs(tmp []int, nums1 []int, nums2 []int) []int {
	i, l1, l2 := len(tmp)-1, len(nums1)-1, len(nums2)-1
	for i >= 0 && l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			tmp[i] = nums1[l1]
			i--
			l1--

			if l1 < 0 {
				_ = copy(tmp[:l2+1], nums2[:l2+1])
				break
			}
		} else {
			tmp[i] = nums2[l2]
			i--
			l2--

			if l2 < 0 {
				_ = copy(tmp[:l1+1], nums1[:l1+1])
				break
			}
		}
	}
	return tmp
}
