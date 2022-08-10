package main

import "log"

//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for i1, value1 := range nums {
		for i2, value2 := range nums {
			if i1 != i2 && value2+value1 == target {
				result = append(result, i1, i2)
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	log.Println(result)
}
