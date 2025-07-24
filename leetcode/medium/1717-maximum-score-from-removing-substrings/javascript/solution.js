var maximumGain = function (s, x, y) {
  let [maxValuableChar, minValuableChar] = ["a", "b"];
  let [maxScore, minScore] = [x, y];
  let [maxValueableCharCounter, minValuableCharCounter] = [0, 0];
  let totalScore = 0;
  if (x < y) {
    maxValuableChar = "b";
    minValuableChar = "a";
    maxScore = y;
    minScore = x;
  }
  for (let i = 0; i < s.length; i++) {
    if (maxValuableChar == s[i]) {
      maxValueableCharCounter++;
    } else if (minValuableChar == s[i]) {
      if (maxValueableCharCounter > 0) {
        maxValueableCharCounter--;
        totalScore += maxScore;
      } else {
        minValuableCharCounter++;
      }
    } else {
      totalScore +=
        Math.min(maxValueableCharCounter, minValuableCharCounter) * minScore;
      maxValueableCharCounter = 0;
      minValuableCharCounter = 0;
    }
  }
  totalScore +=
    Math.min(maxValueableCharCounter, minValuableCharCounter) * minScore;
  return totalScore;
};

function main() {
  const tests = [
    {
      s: "cdbcbbaaabab",
      x: 4,
      y: 5,
      expected: 19,
    },
    {
      s: "aabbaaxybbaabb",
      x: 5,
      y: 4,
      expected: 20,
    },
    {
      s: "ababbababa",
      x: 2,
      y: 3,
      expected: 14,
    },
    {
      s: "abcabcabc",
      x: 10,
      y: 1,
      expected: 30,
    },
  ];

  tests.forEach((tc, i) => {
    const result = maximumGain(tc.s, tc.x, tc.y);
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
