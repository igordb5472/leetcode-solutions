package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m-1, n-1; i+j >= -1 && j >= 0; {
		if i < 0 || nums2[j] >= nums1[i] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}
	}
}
