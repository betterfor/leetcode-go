package code

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = m - 1, n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	if j >= 0 {
		for k := 0; k <= j; k++ {
			nums1[k] = nums2[k]
		}
	}
}
