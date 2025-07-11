package main

import (
	"fmt"
	"reflect"
)

func twoSum(numbers []int, target int) []int {
	leftPointer, rightPointer := 0, len(numbers)-1
	for numbers[leftPointer]+numbers[rightPointer] != target {
		if numbers[leftPointer]+numbers[rightPointer] > target {
			rightPointer--
		} else {
			leftPointer++
		}
	}
	return []int{leftPointer + 1, rightPointer + 1}
}

func main() {
	testCases := []struct {
		numbers []int
		target  int
		expect  []int
	}{
		{numbers: []int{2, 7, 11, 15}, target: 9, expect: []int{1, 2}},
		{numbers: []int{2, 3, 4}, target: 6, expect: []int{1, 3}},
		{numbers: []int{-1, 0}, target: -1, expect: []int{1, 2}},
	}

	for i, tc := range testCases {
		result := twoSum(tc.numbers, tc.target)
		if reflect.DeepEqual(result, tc.expect) {
			fmt.Printf("Test case %d passed ✅: got %v\n", i+1, result)
		} else {
			fmt.Printf("Test case %d failed ❌: got %v, expected %v\n", i+1, result, tc.expect)
		}
	}
}
