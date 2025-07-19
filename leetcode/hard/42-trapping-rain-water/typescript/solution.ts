function trap(height: number[]): number {
  let sum = 0;
  let lenHeights = height.length;
  let [maxLeft, i] = [height[0], 0];
  let [maxRight, j] = [height[lenHeights - 1], lenHeights - 1];
  while (i < j) {
    if (maxLeft <= maxRight) {
      sum += maxLeft - height[i];
      i++;
      maxLeft = Math.max(maxLeft, height[i]);
    } else {
      sum += maxRight - height[j];
      j--;
      maxRight = Math.max(maxRight, height[j]);
    }
  }
  return sum;
}

function main() {
  type TestCase = {
    height: number[];
    expected: number;
  };

  const tests: TestCase[] = [
    { height: [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1], expected: 6 },
    { height: [4, 2, 0, 3, 2, 5], expected: 9 },
    { height: [], expected: 0 },
    { height: [1, 2, 3], expected: 0 },
    { height: [3, 2, 1], expected: 0 },
    { height: [2, 0, 2], expected: 2 },
    { height: [4, 2, 3], expected: 1 },
  ];

  tests.forEach((tc, i) => {
    const result = trap(tc.height);
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

