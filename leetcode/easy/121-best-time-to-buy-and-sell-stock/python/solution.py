from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        current_profit = 0

        for i in range(1, len(prices)):
            current_change = prices[i] - prices[i-1]
            current_profit = max(0, current_profit+current_change)
            max_profit = max(max_profit, current_profit)

        return max_profit


def main():
    solution = Solution()

    test_cases = [
        {
            "prices": [7, 1, 5, 3, 6, 4],
            "expected": 5  # Buy at 1, sell at 6
        },
        {
            "prices": [7, 6, 4, 3, 1],
            "expected": 0  # No profit possible
        },
        {
            "prices": [1, 2, 3, 4, 5],
            "expected": 4  # Buy at 1, sell at 5
        },
        {
            "prices": [2, 4, 1],
            "expected": 2  # Buy at 2, sell at 4
        },
        {
            "prices": [3, 3, 3, 3],
            "expected": 0  # No change in price
        }
    ]

    for i, tc in enumerate(test_cases, 1):
        result = solution.maxProfit(tc["prices"])
        if result == tc["expected"]:
            print(f"✅ Test case {i} passed")
        else:
            print(f"❌ Test case {i} failed: got {result}, expected {tc['expected']}")


if __name__ == "__main__":
    main()
