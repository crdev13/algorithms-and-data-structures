from typing import List


class Solution:
    def maxArea(
        self, h: int, w: int, horizontalCuts: List[int], verticalCuts: List[int]
    ) -> int:
        horizontalCuts.append(h)
        horizontalCuts.sort()
        verticalCuts.append(w)
        verticalCuts.sort()
        maxHorizontalDistance = 0
        horizontalReference = 0

        for cut in horizontalCuts:
            maxHorizontalDistance = max(
                maxHorizontalDistance, cut - horizontalReference
            )
            horizontalReference = cut

        maxVerticalDistance = 0
        verticalReference = 0
        for cut in verticalCuts:
            maxVerticalDistance = max(maxVerticalDistance, cut - verticalReference)
            verticalReference = cut

        return (maxHorizontalDistance * maxVerticalDistance) % (10**9 + 7)


def main():
    MOD = 10**9 + 7

    test_cases = [
        {
            "h": 5,
            "w": 4,
            "horizontalCuts": [1, 2, 4],
            "verticalCuts": [1, 3],
            "expected": 4,
        },
        {"h": 5, "w": 4, "horizontalCuts": [3, 1], "verticalCuts": [1], "expected": 6},
        {"h": 5, "w": 4, "horizontalCuts": [3], "verticalCuts": [3], "expected": 9},
        {
            "h": 100000,
            "w": 100000,
            "horizontalCuts": [50000],
            "verticalCuts": [50000],
            "expected": (50000 * 50000) % MOD,
        },
    ]

    for i, tc in enumerate(test_cases, 1):
        result = Solution().maxArea(
            tc["h"], tc["w"], tc["horizontalCuts"], tc["verticalCuts"]
        )
        if result == tc["expected"]:
            print(f"✅ Test case {i} passed")
        else:
            print(f"❌ Test case {i} failed: got {result}, expected {tc['expected']}")


if __name__ == "__main__":
    main()
