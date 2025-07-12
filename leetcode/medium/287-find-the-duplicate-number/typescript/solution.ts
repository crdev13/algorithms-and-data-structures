function findDuplicate(nums: number[]): number {
  let slow = nums[0];
  let fast = nums[0];
  do {
    slow = nums[slow];
    fast = nums[nums[fast]];
  } while (slow != fast);
  slow = nums[0];
  while (slow != fast) {
    slow = nums[slow];
    fast = nums[fast];
  }
  return slow;
}

function main() {
  type TestCase = {
    nums: number[];
    expected: number;
  };

  const tests: TestCase[] = [
    { nums: [1, 3, 4, 2, 2], expected: 2 },
    { nums: [3, 1, 3, 4, 2], expected: 3 },
    { nums: [3, 3, 3, 3, 3], expected: 3 },
  ];

  tests.forEach((tc, i) => {
    const result = findDuplicate(tc.nums);
    if (result === tc.expected) {
      console.log(`✅ Test case ${i + 1} passed`);
    } else {
      console.log(
        `❌ Test case ${i + 1} failed: got ${result}, expected ${tc.expected}`,
      );
    }
  });
}

main();
