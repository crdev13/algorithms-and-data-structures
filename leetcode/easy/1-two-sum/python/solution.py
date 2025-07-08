from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dictionary = {}
        for i in range(len(nums)):
            num = nums[i]
            if num in dictionary:
                return [dictionary[num], i]
            else:
                dictionary[target - num] = i
        return []


def run_tests():
    sol = Solution()
    test_cases = [
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
    ]

    for nums, target, expected in test_cases:
        result = sol.twoSum(nums, target)
        status = "✅" if result == expected else "❌"
        print(
            f"{status} Input: nums= '{nums}', target={target} | Expected: {expected}, Got: {result}"
        )


if __name__ == "__main__":
    run_tests()
