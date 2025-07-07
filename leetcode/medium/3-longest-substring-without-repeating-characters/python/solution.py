class Solution(object):
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen_characters = set()
        longest_substring = 0
        left = 0
        for right in range(len(s)):
            while s[right] in seen_characters:
                seen_characters.remove(s[left])
                left += 1
            seen_characters.add(s[right])
            longest_substring = max(longest_substring, right - left + 1)
        return longest_substring


def run_tests():
    sol = Solution()
    test_cases = [
        ("abcabcbb", 3),
        ("bbbbb", 1),
        ("pwwkew", 3),
        ("", 0),
        ("abcdef", 6),
        ("abba", 2),
    ]

    for s, expected in test_cases:
        result = sol.lengthOfLongestSubstring(s)
        status = "✅" if result == expected else "❌"
        print(f"{status} Input: '{s}' | Expected: {expected}, Got: {result}")


if __name__ == "__main__":
    run_tests()
