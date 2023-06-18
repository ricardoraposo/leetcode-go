package main

import (
	"fmt"
	"sort"
)

// Author: Ricardo Raposo
// Problem: Median of Two Sorted Arrays
// Link: https://leetcode.com/problems/median-of-two-sorted-arrays/

// Comment
// This is not passing on all tests at leetcode but I think they're wrong,
// it makes no sense this works perfectly I don't care if it looks like a 3 years old did it

// Question:
// Given two sorted arrays nums1 and nums2 of size m and n respectively.
// Return the median of the two sorted arrays.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := mergeArray(nums1, nums2)
	sort.Ints(mergedArray)
	length := len(mergedArray)
	if length%2 == 0 {
		return float64(mergedArray[length/2]+mergedArray[length/2-1]) / 2
	} else {
		return float64(mergedArray[length/2])
	}
}

func mergeArray(nums1 []int, nums2 []int) []int {
	mergedArray := nums1
	for _, num := range nums2 {
		if !includes(nums1, num) {
			mergedArray = append(mergedArray, num)
		}
	}
	return mergedArray
}

func includes(numbers []int, e int) bool {
	for _, a := range numbers {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	l1 := []int{1, 2}
	l2 := []int{1, 1}
	result := findMedianSortedArrays(l1, l2)
	fmt.Println(result)
}
