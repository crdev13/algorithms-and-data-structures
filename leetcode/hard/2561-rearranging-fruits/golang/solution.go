package main

import (
	"fmt"
	"math"
	"sort"
)

func minCost(basket1 []int, basket2 []int) int64 {
	nums := map[int]int{}
	mapBasket1 := map[int]int{}
	mapBasket2 := map[int]int{}
	minNumber := math.MaxInt
	for _, num := range basket1 {
		nums[num] = nums[num] + 1
		mapBasket1[num] = mapBasket1[num] + 1
		minNumber = min(minNumber, num)
	}
	for _, num := range basket2 {
		nums[num] = nums[num] + 1
		mapBasket2[num] = mapBasket2[num] + 1
		minNumber = min(minNumber, num)
	}

	for num, times := range nums {
		if times%2 != 0 {
			return -1
		}
		nums[num] = times / 2
	}

	rest1 := []int{}
	for num, times := range mapBasket1 {
		if times > nums[num] {
			for i := 0; i < times-nums[num]; i++ {
				rest1 = append(rest1, num)
			}
		}
	}
	rest2 := []int{}
	for num, times := range mapBasket2 {
		if times > nums[num] {
			for i := 0; i < times-nums[num]; i++ {
				rest2 = append(rest2, num)
			}
		}
	}

	sort.Ints(rest1)
	sort.Ints(rest2)

	cost := 0

	for index := range rest1 {
		currentMinNumber := min(rest1[index], rest2[len(rest2)-index-1])
		if minNumber*2 < currentMinNumber {
			cost += (2 * minNumber)
		} else {
			cost += currentMinNumber
		}
	}
	return int64(cost)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TestCase struct {
	basket1  []int
	basket2  []int
	expected int64
}

func main() {
	tests := []TestCase{
		{
			basket1:  []int{4, 2, 2, 2},
			basket2:  []int{1, 4, 1, 2},
			expected: 1,
		},
		{
			basket1:  []int{2, 3, 4, 1},
			basket2:  []int{3, 2, 5, 1},
			expected: -1,
		},
		{
			basket1:  []int{1, 1, 2, 3},
			basket2:  []int{1, 2, 2, 3},
			expected: -1,
		},
		{
			basket1:  []int{10},
			basket2:  []int{10},
			expected: 0, // Already equal
		},
		{
			basket1:  []int{5, 5, 5},
			basket2:  []int{5, 5, 5},
			expected: 0, // No swap needed
		},
		{
			basket1:  []int{4, 4, 4, 4, 3},
			basket2:  []int{5, 5, 5, 5, 3},
			expected: 8,
		},
		{
			basket1:  []int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88},
			basket2:  []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8},
			expected: 48,
		},
	}

	for i, tc := range tests {
		result := minCost(tc.basket1, tc.basket2)
		if result == tc.expected {
			fmt.Printf("✅ Test case %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test case %d failed: got %d, expected %d\n",
				i+1, result, tc.expected)
		}
	}
}
