from typing import List, Callable


def run_merge_tests(merge_function: Callable[[List[int], List[int]], List[int]]):
    """
    Runs a series of tests for the merge function.
    merge_function should return a merged array.
    Arguments:
        merge_function: function(A, B) -> List[int]
    """
    test_cases = [
        {
            "A": [1, 3, 5, 0, 0, 0],
            "B": [2, 4, 6],
            "expected": [1, 2, 3, 4, 5, 6],
            "description": "Interleaved merge",
        },
        {
            "A": [1, 2, 3, 0, 0, 0],
            "B": [4, 5, 6],
            "expected": [1, 2, 3, 4, 5, 6],
            "description": "B entirely greater than A",
        },
        {
            "A": [4, 5, 6, 0, 0, 0],
            "B": [1, 2, 3],
            "expected": [1, 2, 3, 4, 5, 6],
            "description": "B entirely smaller than A",
        },
        {
            "A": [2, 4, 6, 0, 0, 0],
            "B": [1, 3, 5],
            "expected": [1, 2, 3, 4, 5, 6],
            "description": "Perfect alternating merge",
        },
        {
            "A": [1, 2, 3, 0, 0, 0],
            "B": [],
            "expected": [1, 2, 3, 0, 0, 0],
            "description": "Empty B array",
        },
        {
            "A": [0, 0, 0],
            "B": [1, 2, 3],
            "expected": [1, 2, 3],
            "description": "Empty A with only buffer",
        },
        {
            "A": [1, 2, 4, 0, 0],
            "B": [3, 5],
            "expected": [1, 2, 3, 4, 5],
            "description": "Partial merge",
        },
    ]

    for i, case in enumerate(test_cases, 1):
        A = case["A"][:]  # copy to avoid mutation between tests
        result = merge_function(A, case["B"])
        assert result == case["expected"], (
            f"Test {i} failed: {case['description']}\nExpected {case['expected']}, got {result}"
        )
        print(f"Test {i} passed: {case['description']}")


def sorted_merged_1(A: List[int], B: List[int]) -> List[int]:
    if len(B) == 0:
        return A

    elements_to_move = len(B)
    a_pointer, b_pointer, last_element = (
        len(A) - elements_to_move - 1,
        len(B) - 1,
        len(A) - 1,
    )
    while elements_to_move > 0:
        if a_pointer >= 0 and b_pointer >= 0 and A[a_pointer] >= B[b_pointer]:
            A[last_element] = A[a_pointer]
            A[a_pointer] = 0
            a_pointer -= 1
        elif a_pointer >= 0 and b_pointer >= 0 and A[a_pointer] < B[b_pointer]:
            A[last_element] = B[b_pointer]
            B[b_pointer] = 0
            b_pointer -= 1
            elements_to_move -= 1
        elif a_pointer < 1:
            A[last_element] = B[b_pointer]
            B[b_pointer] = 0
            b_pointer -= 1
            elements_to_move -= 1
        else:
            A[last_element] = A[a_pointer]
            A[a_pointer] = 0
            a_pointer -= 1
        last_element -= 1
    return A


def main():
    run_merge_tests(sorted_merged_1)


if __name__ == "__main__":
    main()
