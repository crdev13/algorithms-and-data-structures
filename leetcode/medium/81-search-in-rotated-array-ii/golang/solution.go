package main

import (
	"fmt"
	"reflect"
)

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return true
		}
		if nums[left] <= nums[middle] {
			if nums[left] <= target && target < nums[middle] {
				right = middle - 1
			} else if nums[left] == nums[middle] {
				left++
			} else {
				left = middle + 1
			}
		} else {
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else if nums[middle] == nums[right] {
				right--
			} else {
				right = middle - 1
			}
		}
	}

	return false
}

type TestCase struct {
	nums     []int
	target   int
	expected bool
}

func main() {
	tests := []TestCase{
		// LeetCode examples
		{
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},

		// No rotation
		{
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: false,
		},

		// Single element
		{
			nums:     []int{1},
			target:   1,
			expected: true,
		},
		{
			nums:     []int{1},
			target:   0,
			expected: false,
		},

		// All duplicates
		{
			nums:     []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			nums:     []int{1, 1, 1, 1, 1},
			target:   2,
			expected: false,
		},

		// Duplicates straddling pivot
		{
			nums:     []int{2, 2, 2, 3, 4, 2},
			target:   3,
			expected: true,
		},
		{
			nums:     []int{2, 2, 2, 3, 4, 2},
			target:   5,
			expected: false,
		},

		// Targets at boundaries
		{
			nums:     []int{4, 4, 5, 6, 7, 0, 1, 2, 4},
			target:   4,
			expected: true,
		},
		{
			nums:     []int{6, 7, 0, 1, 2, 4, 4, 4, 5},
			target:   5,
			expected: true,
		},

		// Two elements, rotated
		{
			nums:     []int{3, 1},
			target:   1,
			expected: true,
		},
		{
			nums:     []int{3, 1},
			target:   2,
			expected: false,
		},

		// Many duplicates with small unique set
		{
			nums:     []int{1, 1, 3, 1},
			target:   3,
			expected: true,
		},
		{
			nums:     []int{1, 1, 3, 1},
			target:   2,
			expected: false,
		},
		{
			nums:     []int{1, 0, 1, 1, 1},
			target:   0,
			expected: true,
		},
		{
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			target:   2,
			expected: true,
		},

		// Larger mixed case
		{
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   0,
			expected: true,
		},
		{
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   8,
			expected: false,
		},
	}

	for i, tc := range tests {
		// Keep original for mutation check
		orig := append([]int(nil), tc.nums...)
		got := search(tc.nums, tc.target)

		if got == tc.expected {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: nums=%v target=%d -> got %v, expected %v\n",
				i+1, tc.nums, tc.target, got, tc.expected)
		}

		// Warn if function mutates input
		if !reflect.DeepEqual(tc.nums, orig) {
			fmt.Printf("⚠️  Warning (test %d): nums was mutated during search\n", i+1)
		}
	}
}
