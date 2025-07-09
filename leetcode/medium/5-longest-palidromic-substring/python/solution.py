class Solution:
    def longestPalindrome(self, s: str) -> str:
        longest = s[0]
        for index in range(0, len(s)):
            even_longest_substring = self.computePalindromeLen(s, index, index + 1)
            odd_longest_substring = self.computePalindromeLen(s, index, index + 2)
            print(even_longest_substring + "----")
            print(odd_longest_substring + "----")
            temp_max = (
                even_longest_substring
                if len(even_longest_substring) >= len(odd_longest_substring)
                else odd_longest_substring
            )
            longest = longest if len(longest) >= len(temp_max) else temp_max
        return longest

    def computePalindromeLen(self, text, leftIndex, rightIndex):
        if rightIndex >= len(text):
            return ""
        while (
            leftIndex >= 0
            and rightIndex < len(text)
            and text[leftIndex] == text[rightIndex]
        ):
            leftIndex -= 1
            rightIndex += 1
        return text[leftIndex + 1 : rightIndex]


def run_tests():
    sol = Solution()
    test_cases = [
        ("babad", "bab"),
        ("cbbd", "bb"),
    ]

    for text, expected in test_cases:
        result = sol.longestPalindrome(text)
        status = "✅" if result == expected else "❌"
        print(f"{status} Input: text= '{text}' | Expected: {expected}, Got: {result}")


if __name__ == "__main__":
    run_tests()
