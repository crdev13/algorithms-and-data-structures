class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function detectCycle(head: ListNode | null): ListNode | null {
  let slow = head;
  let fast = head;
  do {
    slow = slow?.next ?? null;
    fast = fast?.next?.next ?? null;
  } while (fast != null && fast != slow);
  if (fast == null) {
    return null;
  }
  slow = head;
  while (slow != fast) {
    slow = slow?.next ?? null;
    fast = fast?.next ?? null;
  }
  return slow;
}
function main() {
  type TestCase = {
    head: number[];
    pos: number;
    expected: number; // index of the cycle start, -1 if no cycle
  };

  const tests: TestCase[] = [
    { head: [3, 2, 0, -4], pos: 1, expected: 1 },
    { head: [1, 2], pos: 0, expected: 0 },
    { head: [1], pos: -1, expected: -1 },
  ];

  // Local helper to build the list and apply the cycle
  function buildList(values: number[], pos: number): ListNode | null {
    if (values.length === 0) return null;

    const nodes = values.map((v) => new ListNode(v));
    for (let i = 0; i < nodes.length - 1; i++) {
      nodes[i].next = nodes[i + 1];
    }
    if (pos >= 0) {
      nodes[nodes.length - 1].next = nodes[pos];
    }
    return nodes[0];
  }

  // Local helper to get index of the cycle start node
  function getCycleStartIndex(
    head: ListNode | null,
    node: ListNode | null,
  ): number {
    if (!node) return -1;
    let idx = 0;
    while (head && head !== node) {
      head = head.next;
      idx++;
    }
    return head === node ? idx : -1;
  }

  tests.forEach((tc, i) => {
    const head = buildList(tc.head, tc.pos);
    const cycleStart = detectCycle(head);
    const actual = getCycleStartIndex(head, cycleStart);
    if (actual === tc.expected) {
      console.log(`✅ Test case ${i + 1} passed`);
    } else {
      console.log(
        `❌ Test case ${i + 1} failed: got ${actual}, expected ${tc.expected}`,
      );
    }
  });
}

main();
