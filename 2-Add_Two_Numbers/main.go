package main

import (
	"fmt"
	"strconv"
)

// Author: Ricardo Raposo
// Problem: Add two numbers
// Link: https://leetcode.com/problems/add-two-numbers/

// Question:
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

func addTwoNumbers(l1, l2 []int) int {
	y := ""
	z := ""

	for i := len(l1) - 1; i >= 0; i-- {
		y = y + fmt.Sprint(l1[i])
	}

	for i := len(l2) - 1; i >= 0; i-- {
		z = z + fmt.Sprint(l2[i])
	}

	a, _ := strconv.Atoi(y)
	b, _ := strconv.Atoi(z)

	return a + b
}

func turnToArray(result int, sequence []int) []int {
	if result != 0 {
		x := result % 10
		sequence = append([]int{x}, sequence...)
		return turnToArray(result/10, sequence)
	}
	return sequence
}

func revertArray(sequence []int) []int {
	revertedSequence := []int{}
	for i := len(sequence) - 1; i >= 0; i-- {
		revertedSequence = append(revertedSequence, sequence[i])
	}
	return revertedSequence
}

func main() {
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}
