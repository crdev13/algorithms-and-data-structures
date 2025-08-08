from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            middle = int((left + right) / 2)
            if nums[middle] == target:
                return middle
            if nums[left] <= nums[middle]:
                if nums[left] <= target and target < nums[middle]:
                    right = middle - 1
                else:
                    left = middle + 1
            else:
                if nums[middle] < target and target <= nums[right]:
                    left = middle + 1
                else:
                    right = middle - 1

        return -1


def run_search_tests():
    solution = Solution()
    test_cases = [
        # Example test cases
        ([4, 5, 6, 7, 0, 1, 2], 0, 4),
        ([4, 5, 6, 7, 0, 1, 2], 3, -1),
        ([1], 0, -1),
        # Additional cases
        ([6, 7, 1, 2, 3, 4, 5], 3, 4),  # Rotated array, target in right half
        ([6, 7, 1, 2, 3, 4, 5], 6, 0),  # Target is at beginning
        ([6, 7, 1, 2, 3, 4, 5], 5, 6),  # Target at end
        ([1, 2, 3, 4, 5, 6, 7], 4, 3),  # No rotation
        ([5, 6, 7, 1, 2, 3, 4], 8, -1),  # Target not present
        ([1], 1, 0),  # Single element, target found
        ([3, 1], 1, 1),  # Two elements, rotated
        ([5, 1, 2, 3, 4], 1, 1),  # Target at index 1 after rotation
    ]

    for i, (nums, target, expected) in enumerate(test_cases, 1):
        result = solution.search(nums, target)
        assert result == expected, (
            f"Test {i} failed: nums={nums}, target={target}, expected={expected}, got={result}"
        )
        print(f"Test {i} passed.")


if __name__ == "__main__":
    run_search_tests()
