package main

import "strconv"

//https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	value := strconv.Itoa(x)
	value2 := ""
	for dolzina := len(value) - 1; 0 <= dolzina; dolzina-- {
		value2 += string(value[dolzina])
	}
	return value == value2
}
