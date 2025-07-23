package main

import (
	"fmt"
)

func isInterleaveDP(s1 string, s1Pointer int, s2 string, s2Pointer int, s3 string, s3Pointer int, dictionary map[string]bool) bool {
	if len(s1) == s1Pointer && len(s2) == s2Pointer && len(s3) == s3Pointer {
		return true
	}
	result := false
	if s1Pointer+s2Pointer != s3Pointer {
		return false
	}
	if s1Pointer < len(s1) && s1[s1Pointer] == s3[s3Pointer] {
		key := fmt.Sprintf("%v-%v-%v", s1Pointer, s2Pointer, s3Pointer)
		if val, ok := dictionary[key]; ok {
			result = result || val
		} else {
			tmpResult := isInterleaveDP(s1, s1Pointer+1, s2, s2Pointer, s3, s3Pointer+1, dictionary)
			result = result || tmpResult
			dictionary[key] = tmpResult
		}
	}
	if s2Pointer < len(s2) && s2[s2Pointer] == s3[s3Pointer] {
		key := fmt.Sprintf("%v-%v-%v", s1Pointer, s2Pointer, s3Pointer)
		if val, ok := dictionary[key]; ok && result {
			result = result || val
		} else {
			tmpResult := isInterleaveDP(s1, s1Pointer, s2, s2Pointer+1, s3, s3Pointer+1, dictionary)
			result = result || tmpResult
			dictionary[key] = tmpResult
		}
	}

	return result
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if s1 == "" && s2 == "" && s3 == "" {
		return true
	}

	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return isInterleaveDP(s1, 0, s2, 0, s3, 0, make(map[string]bool))
}

func main() {
	type TestCase struct {
		s1, s2, s3 string
		expected   bool
	}

	tests := []TestCase{
		{s1: "a", s2: "b", s3: "a", expected: false},
		{s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac", expected: true},
		{s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc", expected: false},
		{s1: "", s2: "", s3: "", expected: true},
		{s1: "abc", s2: "", s3: "abc", expected: true},
		{s1: "", s2: "abc", s3: "abc", expected: true},
		{s1: "abc", s2: "def", s3: "adbcef", expected: true},
		{s1: "abc", s2: "def", s3: "abdecf", expected: true},
		{s1: "abababababababababababababababababababababababababababababababababababababababababababababababababb", s2: "babababababababababababababababababababababababababababababababababababababababababababababababaaaba", s3: "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb", expected: false},
	}

	for i, tc := range tests {
		result := isInterleave(tc.s1, tc.s2, tc.s3)
		if result == tc.expected {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: got %v, expected %v\n", i+1, result, tc.expected)
		}
	}
}
