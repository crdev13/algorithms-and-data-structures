class Solution(object):
    def coinChange(self, coins: list[int], amount: int) -> int:
        coins.sort()
        ways_to_make_change = [float("inf")] * (amount + 1)
        ways_to_make_change[0] = 0
        for temp_amount in range(1, amount + 1):
            for coin in coins:
                difference = temp_amount - coin
                if difference < 0:
                    break

                ways_to_make_change[temp_amount] = min(
                    ways_to_make_change[temp_amount],
                    1 + ways_to_make_change[difference],
                )

        result = ways_to_make_change[amount]
        return int(result) if result < float("inf") else -1


def run_tests():
    sol = Solution()
    test_cases = [
        ([1, 2, 5], 11, 3),
        ([1, 2, 4], 5, 2),
        ([2], 3, -1),
        ([1], 0, 0),
    ]

    for coins, amount, expected in test_cases:
        result = sol.coinChange(coins, amount)
        status = "✅" if result == expected else "❌"
        print(f"{status} Input: '{coins}' | Expected: {expected}, Got: {result}")


if __name__ == "__main__":
    run_tests()
