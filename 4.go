package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, m, n, a, b := -1, -1, len(nums1), len(nums2), 0, 0
	if m == 0 {
		a = nums2[n/2]
		b = a
		if n%2 == 0 {
			a = nums2[n/2-1]
		}
		i, j = m, n
	}
	if n == 0 {
		a = nums1[m/2]
		b = a
		if m%2 == 0 {
			a = nums1[m/2-1]
		}
		i, j = m, n
	}
	for i+j < (m+n)/2 {
		switch {
		case i+1 >= m:
			j++
			a, b = b, nums2[j]
		case j+1 >= n:
			i++
			a, b = b, nums1[i]
		case nums1[i+1] <= nums2[j+1]:
			i++
			a, b = b, nums1[i]
		default:
			j++
			a, b = b, nums2[j]
		}
	}
	if (m+n)%2 == 1 {
		return float64(b)
	}
	return (float64(a) + float64(b)) / 2
}
