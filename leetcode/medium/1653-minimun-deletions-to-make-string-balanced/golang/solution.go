package main

import (
	"fmt"
	"reflect"
)

func minimumDeletions(s string) int {
	n := len(s)

	aCount := 0

	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' {
			aCount++
		}
	}

	aMinDeletions := aCount
	bCount := 0

	for index := range s {
		aMinDeletions = min(aMinDeletions, bCount+aCount)
		if s[index] == 'b' {
			bCount++
		} else {
			aCount--
		}
	}

	return min(bCount, aMinDeletions)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	type TestCase struct {
		text     string
		expected int
	}

	tests := []TestCase{
		{text: "aabbbbaabababbbbaaaaaabbababaaabaabaabbbabbbbabbabbababaabaababbbbaaaaabbabbabaaaabbbabaaaabbaaabbbaabbaaaaabaa", expected: 52},
		{text: "aababbab", expected: 2},
		{text: "bbaaaaabb", expected: 2},
		{text: "bbbb", expected: 0},
		{text: "aaa", expected: 0},
	}

	for i, tc := range tests {
		result := minimumDeletions(tc.text)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: got %v, expected %v\n", i+1, result, tc.expected)
		}
	}
}
