from typing import List


class Solution:
    def maxSum(self, nums: List[int]) -> int:
        maxElement = -101
        maxSum = maxElement

        setOfNumbers = set(nums)
        for number in setOfNumbers:
            maxElement = max(maxElement, number)

        setOfNumbers.remove(maxElement)
        maxSum = maxElement

        for number in setOfNumbers:
            if maxSum + number > maxSum:
                maxSum += number

        return maxSum


def main():
    test_cases = [
        {"nums": [1, 2, 3, 4, 5], "expected": 15},
        {"nums": [1, 1, 0, 1, 1], "expected": 1},
        {"nums": [1, 2, -1, -2, 1, 0, -1], "expected": 3},
        {"nums": [5, 5, 5, 5], "expected": 5},
        {"nums": [-1, -2, -3, -4], "expected": -1},
        {"nums": [0, 0, 0], "expected": 0},
        {
            "nums": [4, 2, 1, 2, 3, 4],
            "expected": 10,  # [1,2,3,4]
        },
    ]

    solution = Solution()
    for i, tc in enumerate(test_cases, 1):
        result = solution.maxSum(tc["nums"])
        if result == tc["expected"]:
            print(f"✅ Test case {i} passed")
        else:
            print(f"❌ Test case {i} failed: got {result}, expected {tc['expected']}")


if __name__ == "__main__":
    main()
