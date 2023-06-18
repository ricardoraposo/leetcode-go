package main

import (
	"fmt"
	"sort"
)

// Author: Ricardo Raposo
// Problem: Median of Two Sorted Arrays
// Link: https://leetcode.com/problems/median-of-two-sorted-arrays/

// Question:
// Given two sorted arrays nums1 and nums2 of size m and n respectively.
// Return the median of the two sorted arrays.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	l := len(merged)
	if l%2 == 0 {
		return float64(merged[l/2]+merged[l/2-1]) / 2
	} else {
		return float64(merged[l/2])
	}
}

func main() {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{3, 4, 5, 6}
	result := findMedianSortedArrays(l1, l2)
	fmt.Println(result)
}
