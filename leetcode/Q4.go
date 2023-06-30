package leetcode

/*
Time Complexity: O(log(min(m, n)))
Space Complexity: O(1)
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	imin, imax, halflen := 0, m, (m+n+1)/2

	for imin <= imax {
		i := (imin + imax) / 2
		j := halflen - i

		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxleft, minright int

			if i == 0 {
				maxleft = nums2[j-1]
			} else if j == 0 {
				maxleft = nums1[i-1]
			} else {
				maxleft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxleft)
			}

			if i == m {
				minright = nums2[j]
			} else if j == n {
				minright = nums1[i]
			} else {
				minright = min(nums1[i], nums2[j])
			}

			return float64(maxleft+minright) / 2.0
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
