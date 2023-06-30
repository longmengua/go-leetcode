package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).
*/
func TestQ4(t *testing.T) {
	var nums1 []int
	var nums2 []int
	var result float64

	nums1 = []int{1, 3}
	nums2 = []int{2}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(2.5))

	nums1 = []int{}
	nums2 = []int{3}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(3))

	nums1 = []int{0, 0, 0, 0, 0}
	nums2 = []int{-1, 0, 0, 0, 0, 0, 1}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(0))

	nums1 = []int{}
	nums2 = []int{2, 3}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(2.5))

	nums1 = []int{1, 1, 1}
	nums2 = []int{1, 1, 1}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(1))

	nums1 = []int{3, 4, 5}
	nums2 = []int{1, 2}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(3))

	nums1 = []int{3}
	nums2 = []int{-2, -1}
	result = FindMedianSortedArrays(nums1, nums2)
	Assert(t, result, float64(-1))

	count := 1000
	numberOfDigit := 1000
	min := math.MinInt32
	max := math.MaxInt32
	arr := []time.Duration{}
	for i := 0; i < count; i++ {
		nums1 = generateSortedArray(numberOfDigit, min, max)
		nums2 = generateSortedArray(numberOfDigit, min, max)
		start := time.Now()
		FindMedianSortedArrays(nums1, nums2)
		duration := time.Since(start)
		arr = append(arr, duration)
	}
	fmt.Printf("Average Execution Time: %s\n", calculateAverageDuration(arr))
}

func generateSortedArray(length, min, max int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}

	sort.Ints(arr)
	return arr
}

func calculateAverageDuration(arr []time.Duration) time.Duration {
	sum := time.Duration(0)
	for _, duration := range arr {
		sum += duration
	}

	average := sum / time.Duration(len(arr))
	return average
}
