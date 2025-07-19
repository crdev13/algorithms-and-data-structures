package main

import (
	"fmt"
	"reflect"
)

func nextGreaterElements(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen*2)
	tempNums := make([]int, numsLen*2)
	for index := range nums {
		result[index] = -1
		result[index+numsLen] = -1
		tempNums[index] = nums[index]
		tempNums[index+numsLen] = nums[index]
	}
	nums = tempNums

	stack := []int{0}
	for index := 1; index < len(nums); index++ {
		lastIndex := stack[len(stack)-1]
		for len(stack) > 0 && nums[lastIndex] < nums[index] {
			result[lastIndex] = nums[index]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				lastIndex = stack[len(stack)-1]
			}
		}
		stack = append(stack, index)
	}

	return result[:numsLen]
}

func main() {
	type TestCase struct {
		nums     []int
		expected []int
	}

	tests := []TestCase{
		{nums: []int{1, 2, 1}, expected: []int{2, -1, 2}},
		{nums: []int{1, 2, 3, 4, 3}, expected: []int{2, 3, 4, -1, 4}},
		{nums: []int{5, 4, 3, 2, 1}, expected: []int{-1, 5, 5, 5, 5}},
		{nums: []int{1, 1, 1, 1}, expected: []int{-1, -1, -1, -1}},
		{nums: []int{3, 8, 4, 1, 2}, expected: []int{8, -1, 8, 2, 3}},
	}

	for i, tc := range tests {
		result := nextGreaterElements(append([]int{}, tc.nums...)) // prevent mutation
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: got %v, expected %v\n", i+1, result, tc.expected)
		}
	}
}
