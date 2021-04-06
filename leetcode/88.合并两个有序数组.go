package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
	} else if n != 0 {
		for i, j := m, 0; i < m+n; i, j = i+1, j+1 {
			nums1[i] = nums2[j]
		}
		sort.Ints(nums1)
	}
}
