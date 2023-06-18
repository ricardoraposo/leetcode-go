package main

import (
	"fmt"
	"strconv"
)

// Author: Ricardo Raposo
// Problem: Palindrome Number
// Link: https://leetcode.com/problems/palindrome-number/

// Question:
// Given an integer x, return true if x is a palindrome , and false otherwise.

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	if strX == revertString(strX) {
		return true
	}
  return false
}

func revertString(s string) string {
	revStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		revStr = revStr + string(s[i])
	}
	return revStr
}

func main() {
  result1 := isPalindrome(121)
  result2 := isPalindrome(123)
  fmt.Println(result1)
  fmt.Println(result2)
}
