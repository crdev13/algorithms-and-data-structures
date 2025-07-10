package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	leftPointer := &ListNode{0, head}
	rightPointer := leftPointer
	tmpHead := rightPointer
	for i := 0; i <= n; i++ {
		rightPointer = rightPointer.Next
	}
	for rightPointer != nil {
		rightPointer = rightPointer.Next
		leftPointer = leftPointer.Next
	}
	leftPointer.Next = leftPointer.Next.Next
	return tmpHead.Next
}

func main() {
	type testCase struct {
		input    []int
		n        int
		expected []int
	}

	// Local helper to build linked list
	buildList := func(nums []int) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for _, v := range nums {
			curr.Next = &ListNode{Val: v}
			curr = curr.Next
		}
		return dummy.Next
	}

	// Local helper to convert list to slice
	listToSlice := func(head *ListNode) []int {
		result := make([]int, 0)
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	tests := []testCase{
		{input: []int{1, 2, 3, 4, 5}, n: 2, expected: []int{1, 2, 3, 5}},
		{input: []int{1}, n: 1, expected: []int{}},
		{input: []int{1, 2}, n: 1, expected: []int{1}},
		{input: []int{1, 2}, n: 2, expected: []int{2}},
	}

	for i, tc := range tests {
		head := buildList(tc.input)
		output := listToSlice(removeNthFromEnd(head, tc.n))
		if reflect.DeepEqual(output, tc.expected) {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: got %v, expected %v\n", i+1, output, tc.expected)
		}
	}
}
