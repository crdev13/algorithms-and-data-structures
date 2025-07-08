from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum, current_sum = nums[0], nums[0]
        for i in range(1, len(nums)):
            current_sum = max(nums[i], nums[i] + current_sum)
            max_sum = max(max_sum, current_sum)
        return max_sum


def run_tests():
    sol = Solution()
    test_cases = [
        ([-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
        ([1], 1),
        ([5, 4, -1, 7, 8], 23),
    ]

    for nums, expected in test_cases:
        result = sol.maxSubArray(nums)
        status = "✅" if result == expected else "❌"
        print(f"{status} Input: nums= '{nums}' | Expected: {expected}, Got: {result}")


if __name__ == "__main__":
    run_tests()
