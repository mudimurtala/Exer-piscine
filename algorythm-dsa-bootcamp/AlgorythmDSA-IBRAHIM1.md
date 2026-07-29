<!-- week01_session01_big_o_notation_and_complexity_analysis.md -->

# Week 1, Session 1: Big O Notation and Complexity Analysis

**Topic:** Introduction to Big O notation, time and space complexity, the major complexity classes, and a first full worked interview problem (Contains Duplicate)
**Week:** 1 of 12
**Instructor:** Not stated by name in the transcript (referred to by students only through general greetings)

---

## Why This Session Matters

Before you touch a single algorithm, you need a ruler to measure it with. That ruler is Big O. This session is the foundation for literally everything else in this bootcamp, so do not rush it. If you get comfortable here, patterns like binary search, two pointers, and sliding window will make a lot more sense later, because you will already understand *why* they are fast.

Think of this document as your personal reference. Read it slowly the first time, run the code yourself, and come back to it whenever a later session references complexity, because it will, constantly.

---

## Concept Breakdown

### What Big O Actually Is

Big O notation is a mathematical way of describing how the runtime or memory use of an algorithm grows as the size of its input grows. It is not a stopwatch. It does not measure seconds. It measures a *rate of growth*.

Here is the everyday analogy: imagine two delivery drivers. One gets faster or slower depending on how much traffic there is, the weather, and how tired they are that day. The other one always takes exactly ten minutes no matter what is going on around them, because they use the same fixed route every time. Big O is interested in the *shape* of how delivery time changes as the number of packages grows, not in whether one driver happens to have a faster car today. A faster machine (say a MacBook versus an older laptop) is like a faster car. Big O ignores that. It only cares about the shape of the growth curve as input size increases.

The formal definition, in the instructor's words: Big O describes the upper bound, or worst case scenario, of an algorithm's runtime in terms of input size. As your input size n grows toward infinity, Big O tells you how the number of operations grows alongside it.

### Why Engineers Actually Use This

This is not academic trivia. As a real product gets more users, the number of requests per second and queries per second grows too. If your algorithm's growth rate is bad, your application will feel fine with ten users and completely fall over with ten million. Big O is how you predict that ahead of time instead of discovering it in production.

Concretely, Big O helps you:

* Compare two different approaches to the same problem before you write either one
* Predict how an approach will behave as data grows
* Decide where optimization actually matters, since not every optimization is worth the engineering effort
* Plan resource usage, including memory and processing budgets
* Become a sharper problem solver in general, because you start reading constraints and limits instinctively instead of guessing

### The Two Formal Rules of Big O

There are two rules you need to internalize immediately, because almost everything else builds on them.

**Rule one: ignore constants.** If an algorithm does 2n operations, we call it O(n), not O(2n). If it does n squared work but with some constant multiplier in front, we still just write O(n squared). The constant does not change the shape of the growth curve as n gets huge, so we drop it.

**Rule two: discard lower order terms.** If your total work comes out to n squared plus n, the n squared term will always dominate once n gets large enough, so we simplify the whole thing down to O(n squared). We only keep the term that grows fastest.

A helpful way to picture this: if n is 100, n squared is 10,000, but the extra plain n term only adds 100. That 100 becomes a rounding error compared to 10,000 as n keeps growing, so we drop it.

### Time Complexity Versus Space Complexity

**Time complexity** describes how the number of operations grows as input size grows. It does not literally count seconds, because two different machines run at two different speeds. Instead it is a mathematical relationship between the size of the input and the number of steps needed to process it.

**Space complexity** describes how much memory an algorithm needs as input grows. This is calculated as:

```
space complexity = input space + extra auxiliary space
```

Input space is whatever you were given to work with, your variables, your parameters. Extra auxiliary space is anything additional you allocate yourself: a second array, a hash map, a segment tree you build from scratch, or even the hidden memory used by recursive function calls stacking up.

An important mentoring note here: a solution can be fast in time and terrible in space, or the reverse. Neither one automatically wins. Part of becoming a strong engineer is learning to make the right trade off for the situation you are actually in, not chasing the fastest possible time complexity no matter the memory cost.

### Walking Through the Complexity Classes

This is the part worth memorizing cold, because you will use it in every single problem from here forward.

| Complexity | Name | Everyday feel | Typical trigger |
|---|---|---|---|
| O(1) | Constant | Grabbing a book by its exact shelf position | Fixed number of operations regardless of input size |
| O(log n) | Logarithmic | Finding a word by repeatedly halving the dictionary | Input shrinks by a constant factor each step |
| O(n) | Linear | Reading every page of a book once | One pass through all n elements |
| O(n log n) | Linearithmic | Sorting a deck of cards efficiently | A linear pass combined with a divide and halve step, common in efficient sorting |
| O(n squared) | Quadratic | Comparing every student in a class to every other student | Nested loop where both loops run roughly n times |
| O(2 to the n) | Exponential | Listing every possible combination of toppings on a pizza | Generating all subsets of a set |
| O(n factorial) | Factorial | Trying every possible seating arrangement at a table | Generating all orderings, permutations |

Let's go through each with the reasoning the instructor gave, plus extra grounding.

#### O(1): Constant Time

No matter how large the input gets, the algorithm always takes the same number of steps. Accessing an array by index is the textbook example: whether the array has 10 elements or 10,000, `array[index]` takes one step.

```cpp
// Returns the value stored at the given index.
// This is O(1) because array access by index never depends on array size.
int getFirstElement(vector<int>& arr) {
    return arr[0];
}
```

A subtle point the instructor raised that trips people up: if you count through a fixed set of 26 letters of the alphabet, that is still O(1), even though it feels like a loop is "doing work." The number 26 never changes no matter how big the rest of your problem gets, so it is a constant, not a variable tied to input size. The same logic applies to O(26 times 26), O(1000), or any other fixed number. They all collapse to O(1) because they never scale with n.

```cpp
// Counts vowels among the 26 letters of the alphabet.
// This loop always runs exactly 26 times, regardless of any other input,
// so its time complexity is O(1), not O(n).
int countVowelsInAlphabet() {
    string alphabet = "abcdefghijklmnopqrstuvwxyz";
    string vowels = "aeiou";
    int count = 0;
    for (char c : alphabet) {
        if (vowels.find(c) != string::npos) {
            count++;
        }
    }
    return count;
}
```

#### O(n): Linear Time

The running time grows directly in proportion to the size of the input. The instructor's analogy: this is like the function y equals mx plus c. As x changes, y changes in direct proportion. Print every element in an array once, and that is O(n), because n prints happen for n elements.

```cpp
// Prints every element in the array exactly once.
// Runtime grows in direct proportion to array size, so this is O(n).
void printAllElements(vector<int>& arr) {
    for (int i = 0; i < arr.size(); i++) {
        cout << arr[i] << endl;
    }
}
```

One important trap: looping through only half the array, meaning n divided by two iterations, is still O(n). Remember rule one, we ignore constants, and n divided by two is the same thing as one half times n. The constant one half gets dropped, leaving plain O(n).

#### O(log n): Logarithmic Time

This is one of the most important complexity classes to truly understand, because so many efficient algorithms are built around it, especially anything based on divide and conquer.

The everyday analogy the instructor used, and it is a genuinely good one: imagine searching for a word in a paper dictionary. You could read every single entry from the start, which is linear and slow. Or you could open the dictionary to the middle. If your word comes before that page, you know it must be in the left half, so you throw away the right half entirely and repeat the process on what remains. Each step cuts the remaining search space in half. That halving is the signature of logarithmic time.

**Binary search** is the classic example of this pattern, and we will formalize it fully in the pattern breakdown section below, because pattern recognition matters more than memorizing this one example.

```cpp
// Classic iterative binary search on a sorted array.
// Each comparison eliminates half of the remaining search space,
// giving O(log n) time complexity.
int binarySearch(vector<int>& arr, int target) {
    int low = 0;
    int high = arr.size() - 1;

    while (low <= high) {
        int mid = low + (high - low) / 2;
        if (arr[mid] == target) {
            return mid;
        } else if (arr[mid] < target) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return -1; // not found
}
```

A binary search tree (BST) applies this exact same halving idea, but on a tree shape instead of a flat array. In a balanced BST, every value smaller than the current node lives in the left subtree, and every value larger lives in the right subtree. So if you are looking for a value and it is bigger than the current node, you immediately know you can ignore the entire left subtree and only search right. This is why the instructor said the recursive search on a balanced BST also runs in O(log n): each comparison eliminates roughly half of the remaining nodes.

```cpp
// Recursive search on a binary search tree.
// Comparable to binary search, but on a tree shape instead of a flat array.
// On a balanced tree this runs in O(log n).
struct TreeNode {
    int value;
    TreeNode* left;
    TreeNode* right;
};

TreeNode* searchBST(TreeNode* node, int target) {
    if (node == nullptr || node->value == target) {
        return node;
    }
    if (target < node->value) {
        return searchBST(node->left, target);
    }
    return searchBST(node->right, target);
}
```

Why the logarithm specifically? A balanced binary tree fills each level completely before starting the next one, and each level holds twice as many nodes as the level above it. Level zero has one node, level one has two, level two has four, level three has eight, and so on, following two raised to the power of the level number. If n is the total number of nodes and h is the height of the tree, then n is roughly two to the power of h. Solving for h means taking the logarithm, and importantly, this logarithm is always base two, not base ten, because each step only branches two ways. So height, and therefore the number of comparisons needed, is O(log n).

**A crucial detail on best, average, and worst case:** a perfectly balanced tree gives you the ideal O(log n) case. A tree built from randomly ordered inserts still tends to stay roughly balanced on average, so average case also lands near O(log n). But the worst case happens when you insert already sorted data into a plain BST with no rebalancing. Every new node just becomes a child on the same side, over and over, and the tree degenerates into what is effectively a straight line, called a skewed tree. At that point, searching it is no better than a plain linear scan, meaning it becomes O(n), not O(log n). This is exactly why self balancing trees exist, though that is a topic for a future session.

In interview contexts specifically, when people say "Big O of N" for a general algorithm, they usually mean the worst case, since that is the safest number to reason about. Historically, average case was sometimes written using theta and best case using omega, but in practice, worst case dominates the conversation.

#### O(n log n): Linearithmic Time

This shows up constantly in efficient sorting algorithms. Standard library sort functions, like the one used in C++, typically achieve O(n log n) using algorithms conceptually similar to quick sort. You are doing roughly a linear amount of work, n, but that work is organized using a divide and halve strategy that contributes the log n factor.

```cpp
// Sorting is a very common source of n log n behavior.
// std::sort in C++ typically runs in approximately O(n log n).
sort(nums.begin(), nums.end());
```

#### O(n squared): Quadratic Time

This happens when you have two nested loops, and the performance is governed by the highest power present, which here is two. If the outer loop runs n times and the inner loop also runs n times for every single pass of the outer loop, the total number of operations is n times n, or n squared.

```cpp
// Two nested loops, both running roughly n times.
// Total operations scale as n times n, giving O(n squared).
void printAllPairs(vector<int>& arr) {
    int n = arr.size();
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cout << arr[i] << ", " << arr[j] << endl;
        }
    }
}
```

**A trap worth flagging clearly, because the instructor emphasized this heavily:** two nested loops do not automatically mean O(n squared). You always have to actually reason through what each loop is doing. There are at least three important cases where nested loops do *not* give you n squared:

1. **Fixed inner loop.** If the inner loop only ever runs a fixed number of times, say five, regardless of how big n is, the total work is five times n, which simplifies to O(n), not O(n squared), because five is a constant that gets dropped.

2. **Two different input sizes.** If the outer loop runs across an array of size n and the inner loop runs across a *different* array of size m, the total work is n times m. This is only equal to n squared if n and m happen to be the same value. If the two array sizes genuinely differ, you must keep it written as O(n times m), because collapsing it to n squared would be mathematically wrong and could cause you to misjudge whether your solution fits within the problem's constraints.

3. **The inner pointer never resets.** This is the classic two pointer setup. It looks nested because there is a loop inside a loop, but if the inner pointer only ever moves forward across the *entire* run of the algorithm, and never resets back to the start for each outer iteration, then across the whole algorithm it moves at most n times total, not n times for every single outer step. The total work stays O(n), not O(n squared). We formalize this fully as the two pointer pattern below, because recognizing it correctly is genuinely one of the more subtle skills in this session.

#### O(2 to the n): Exponential Time

This is your first step into what is called nonpolynomial time, as opposed to the polynomial classes above it (constant, linear, quadratic, and so on all involve n raised to some fixed power; exponential instead has a fixed base raised to the power of n, which grows dramatically faster).

The classic trigger is generating every possible subset of a set. If your set has 4 elements, the total number of subsets is two to the fourth power, which is 16. Each additional element in the input doubles the number of subsets, because for every existing subset, you now have the choice to either include the new element or not.

```cpp
// Generates all subsets of a set using the classic
// "take it or leave it" recursive template.
// With n elements there are 2^n possible subsets, so this is O(2^n).
void generateSubsets(vector<int>& nums, int index, vector<int>& current, vector<vector<int>>& result) {
    if (index == nums.size()) {
        result.push_back(current);
        return;
    }
    // Choice one: leave this element out
    generateSubsets(nums, index + 1, current, result);

    // Choice two: take this element
    current.push_back(nums[index]);
    generateSubsets(nums, index + 1, current, result);
    current.pop_back();
}
```

An important mentoring point here: whether you write subset generation recursively or iteratively (for example using bit masking, where each bit represents whether an element is included), you will always land on the same O(2 to the n) complexity. The implementation style does not change the underlying growth rate.

#### O(n factorial): Factorial Time

This is described as the highest form of brute force, and it comes up when you are trying every possible ordering, meaning every permutation, of a set of elements. With 3 elements, there are 3 factorial, meaning 3 times 2 times 1, equal to 6 possible orderings.

```cpp
// Generates all permutations (orderings) of the input elements.
// With n elements, there are n! possible orderings, so this is O(n!).
void generatePermutations(vector<int>& nums, int start, vector<vector<int>>& result) {
    if (start == nums.size()) {
        result.push_back(nums);
        return;
    }
    for (int i = start; i < nums.size(); i++) {
        swap(nums[start], nums[i]);
        generatePermutations(nums, start + 1, result);
        swap(nums[start], nums[i]); // backtrack
    }
}
```

The practical takeaway: n has to stay very small for factorial time to be usable. Once n grows even modestly, the number of operations explodes far beyond what any computer can process in a reasonable time. Three factorial is tiny and completely fine. One hundred factorial is a number so large it is essentially meaningless to try to compute directly.

---

## Pattern Breakdown

This is, by the instructor's own framing, the single most valuable skill in this whole domain, and it is also usually the hardest part to build, harder than understanding any one specific example once someone points the pattern out to you. So let's separate the patterns from the examples that happened to illustrate them.

### Pattern One: Binary Search (Divide and Conquer on Sorted or Monotonic Data)

**Signal that tells you this pattern applies:** the input is sorted, or more generally, there is some monotonic property, meaning you can always tell which half of the remaining space your answer must be in just by checking the middle element. This also applies to searching a balanced binary search tree, and to a whole family of variations like searching for an upper bound, a lower bound, or applying binary search directly on the answer itself rather than on an array (a technique you will meet again later in the bootcamp).

**General template, separate from any one specific example:**

```
function binarySearch(sorted_input, target):
    low = start of input
    high = end of input

    while low <= high:
        mid = low + (high - low) / 2

        if input[mid] equals target:
            return mid
        else if input[mid] < target:
            low = mid + 1        # answer must be in the right half
        else:
            high = mid - 1       # answer must be in the left half

    return "not found"
```

**How the transcript's example is just one instance of this pattern:** the array search example and the balanced binary search tree traversal example both follow this exact shape. In the array version, "mid" is an index calculated from low and high. In the BST version, "mid" is simply whatever value the current node happens to hold, since the tree structure itself already encodes the sorted ordering. In both cases, one comparison lets you discard half of the remaining possibilities, and that is what produces the O(log n) complexity in each case.

### Pattern Two: The Two Pointer Pattern (Inner Pointer That Never Resets)

**Signal that tells you this pattern applies:** you have two loops that look nested at first glance, but on closer inspection, one of the pointers only ever moves forward and is never reset back to the start for each outer iteration. This often shows up in problems about contiguous subarrays or substrings, where you are trying to maintain some kind of sliding window of elements.

**General template:**

```
function twoPointer(array):
    left = 0

    for right in range(0, length of array):
        # expand the window by including array[right]

        while (some condition means the window is invalid):
            # shrink the window from the left
            left = left + 1

        # window from left to right is currently valid,
        # do whatever bookkeeping you need here

    return result
```

**Why this is O(n) and not O(n squared):** even though there is a loop inside a loop, the left pointer only ever moves forward, and across the *entire* run of the algorithm it moves at most n times in total, not n times for every single step of the outer loop. The instructor's own worked example had the outer pointer (right) move 10 times while the inner pointer (left) moved only 7 times total across the whole run, with a window constrained to at most 3 elements wide. Because the total movement of both pointers combined is bounded by roughly 2 times n, the constant gets dropped and the whole thing simplifies to O(n).

This is precisely the trap flagged in the nested loop discussion above: seeing two loops and assuming O(n squared) without checking whether the inner pointer actually resets is one of the most common analysis mistakes at this stage.

### Pattern Three: Nested Loops With Two Independent Input Sizes

**Signal that tells you this pattern applies:** you have two separate collections of possibly different sizes, commonly seen in grid or matrix problems, or in problems comparing two different lists against each other, such as certain graph problems where n might represent the number of cities and m the number of edges.

**General template:**

```
function compareTwoLists(listA of size n, listB of size m):
    for i in range(0, n):
        for j in range(0, m):
            # do work comparing listA[i] and listB[j]

    return result
```

**How this differs from plain O(n squared):** the total work here is n times m. This only equals n squared if n and m happen to be the same value, which will typically be stated explicitly in a problem's constraints. If they differ, for example n equals 4 and m equals 3, you get n times m equals 12 total operations, not 16 (which is what 4 squared would give you). Getting this distinction right matters a great deal when you are checking whether your solution fits inside a problem's time limit.

---

## Worked Code Example: Contains Duplicate

This was the main live coding problem of the session, and it is a genuinely great one for practicing complexity analysis because the instructor deliberately walked through four different approaches, from worst to best, showing the trade offs at each step.

**Problem statement:** given an integer array `nums`, return true if any value appears at least twice in the array, and return false if every element is distinct.

**Reading the constraints first, before writing any code, is itself part of the skill being taught here.** The array size is not fixed, so this cannot be constant time. The instructor also mentioned a rough rule of thumb for C++: roughly 10 to the 8th power operations is generally considered the safe ceiling for a solution to run within typical time limits.

### Approach One: Brute Force With Two Nested Loops

The first instinct, and the correct first step according to the instructor, is to start simple, brute force it, and only optimize from there.

```cpp
// Brute force approach: compare every pair of elements.
// This is O(n squared) time and O(1) extra space.
bool containsDuplicate(vector<int>& nums) {
    int n = nums.size();
    for (int i = 0; i < n - 1; i++) {
        for (int j = i + 1; j < n; j++) {
            if (nums[i] == nums[j]) {
                return true;
            }
        }
    }
    return false;
}
```

With the constraint given in the session (n up to roughly 10 to the 5th power), this approach requires roughly 10 to the 5th times 10 to the 5th, or 10 to the 10th power operations in the worst case, which is far above the roughly 10 to the 8th power ceiling for C++. In the live session this approach actually failed with a time limit exceeded result on submission, which is a great real world demonstration of why complexity analysis matters before you even finish writing the code.

### Approach Two: Sort First, Then Scan Once

If duplicate values are sorted, they end up sitting right next to each other, which means a single linear pass afterward is enough to catch them.

```cpp
// Sort the array first, then scan once for adjacent duplicates.
// Sorting costs O(n log n), the scan afterward costs O(n),
// and since n log n dominates n for large n, the total is O(n log n).
// Space complexity here is O(1) extra space (ignoring the space the sort itself may use internally).
bool containsDuplicate(vector<int>& nums) {
    int n = nums.size();
    sort(nums.begin(), nums.end());

    for (int i = 1; i < n; i++) {
        if (nums[i] == nums[i - 1]) {
            return true;
        }
    }
    return false;
}
```

Why the total is O(n log n) and not O(n log n plus n): this is rule two from earlier in action. The sort's n log n term always outgrows the plain n term from the scan as n gets large, so the lower order term gets dropped, leaving just O(n log n).

### Approach Three: A Hash Set

A set naturally stores only distinct elements, which makes it a very natural fit for a "have I seen this before" question.

```cpp
// Use a hash set to track which values have already been seen.
// Time complexity is O(n), since each element is processed once.
// Space complexity is O(n) in the worst case, since we may need to
// store every element before finding a duplicate (or confirming none exist).
bool containsDuplicate(vector<int>& nums) {
    unordered_set<int> seen;
    for (int num : nums) {
        if (seen.count(num) > 0) {
            return true;
        }
        seen.insert(num);
    }
    return false;
}
```

This trades space for time. It runs in O(n), faster in terms of raw time complexity than the sorting approach, but it now uses O(n) extra auxiliary space, whereas the sorting approach used only constant extra space. Neither approach is objectively "better" in every situation; which one you would reach for depends on whether time or memory is the tighter constraint for your particular scenario.

### Approach Four: A Hash Map (Frequency Counting)

A hash map works almost identically to the set approach here, but it stores a count for each value rather than just marking presence. This makes it a more general tool, since it also naturally answers questions about *how many times* something appears, not just *whether* it appears.

```cpp
// Use a hash map to count occurrences of each value.
// Time complexity: O(n). Space complexity: O(n) in the worst case.
bool containsDuplicate(vector<int>& nums) {
    unordered_map<int, int> seen;
    for (int num : nums) {
        if (seen[num] >= 1) {
            return true;
        }
        seen[num]++;
    }
    return false;
}
```

**Trade off summary for this problem:**

| Approach | Time | Space | Notes |
|---|---|---|---|
| Brute force, two nested loops | O(n squared) | O(1) | Fails for large n, exceeded the time limit in the live session |
| Sort, then single scan | O(n log n) | O(1) extra | Great when memory is tight |
| Hash set | O(n) | O(n) | Fast, natural fit for "seen before" questions |
| Hash map | O(n) | O(n) | Same complexity as the set, but more general since it tracks counts, not just presence |

---

## Extra Code Snippets (Reinforcement Examples Beyond the Transcript)

These are not from the live session. They are extra examples meant to strengthen the same ideas in a slightly different context.

**Reinforcing O(1):** checking whether a number is even or odd never depends on how large the number is.

```python
# Always exactly one operation, regardless of how large n is.
def is_even(n):
    return n % 2 == 0
```

**Reinforcing O(log n) beyond arrays:** this same halving idea applies to finding how many times you can divide a number by two before reaching one, which is the core idea behind the height calculation in a balanced binary tree.

```python
# Counts how many times n can be halved before reaching 1.
# Each step cuts the remaining value in half, so this is O(log n).
def count_halvings(n):
    steps = 0
    while n > 1:
        n = n // 2
        steps += 1
    return steps
```

**Reinforcing the two pointer pattern with a different problem:** finding whether a sorted array contains two numbers that add up to a target value.

```python
# Two pointer approach for finding a pair that sums to a target
# in a SORTED array. Left only ever moves right, right only ever moves left,
# so combined they move at most n times total. This is O(n) time, O(1) space.
def has_pair_with_sum(sorted_arr, target):
    left = 0
    right = len(sorted_arr) - 1

    while left < right:
        current_sum = sorted_arr[left] + sorted_arr[right]
        if current_sum == target:
            return True
        elif current_sum < target:
            left += 1
        else:
            right -= 1

    return False
```

**Reinforcing O(n times m) with two different input sizes:** counting how many pairs can be formed by taking one item from each of two differently sized lists.

```python
# listA has n items, listB has m items.
# This is O(n times m), NOT O(n squared) unless n happens to equal m.
def count_all_pairs(list_a, list_b):
    pairs = []
    for a in list_a:
        for b in list_b:
            pairs.append((a, b))
    return pairs
```

---

## Practice and Brush Up Snippets

Try writing these yourself before checking the exercises section below. These are meant to be quick, standalone typing practice, not full problems.

1. Write a function that returns the last element of an array. State its time complexity.
2. Write a function that sums every element in an array once. State its time complexity.
3. Write a function that, given a sorted array, checks whether a target value exists, without using a built in search function. State its time complexity.
4. Write a function that generates every pair (i, j) where i and j are both indices into the same array of size n, including pairs where i equals j. State its time complexity.
5. Write a function that generates every possible subset of a small set of three elements using the take or leave it recursive template. State its time complexity.

---

## Exercises

### Easy

**Exercise 1:** What is the time complexity of the following function?

```python
def print_first_and_last(arr):
    print(arr[0])
    print(arr[-1])
```

<details>
<summary>Check your answer</summary>

O(1). No matter how large the array is, this always does exactly two operations. The size of the array never affects the number of steps taken.

</details>

**Exercise 2:** What is the time complexity of the following function?

```python
def print_all(arr):
    for item in arr:
        print(item)
```

<details>
<summary>Check your answer</summary>

O(n). The number of print operations grows in direct proportion to the number of elements in the array.

</details>

### Medium

**Exercise 3:** What is the time complexity of the following function, and why is it not what it might first appear to be?

```python
def mystery(arr):
    for i in range(len(arr)):
        for j in range(5):
            print(arr[i], j)
```

<details>
<summary>Check your answer</summary>

O(n), not O(n squared). The outer loop runs n times, but the inner loop always runs exactly 5 times regardless of n. Five is a constant, so the total work is 5 times n, which simplifies to O(n) once we drop the constant, following rule one from this session.

</details>

**Exercise 4:** Two lists, `cities` with n items and `roads` with m items, are compared like this. What is the time complexity, expressed properly?

```python
def compare(cities, roads):
    for c in cities:
        for r in roads:
            check(c, r)
```

<details>
<summary>Check your answer</summary>

O(n times m). This is only equal to O(n squared) if n and m happen to be the same size. Since cities and roads are two genuinely different collections that could have different sizes, the correct and safe way to express this is O(n times m), not O(n squared).

</details>

### Hard (Stretch)

**Exercise 5:** A function uses two pointers, left and right, over a single array of size n. Right moves forward through the entire array exactly once. Left only ever moves forward and, across the entire run of the function, moves a total of at most n times as well, but it never resets to zero between iterations of right. What is the overall time complexity, and why does this not become O(n squared) even though there are two nested loop structures involved?

<details>
<summary>Check your answer</summary>

O(n). Even though there are two pointers moving, each one individually moves at most n times across the *entire* run of the algorithm, not n times for every single step of the other pointer. Total combined movement is bounded by roughly 2 times n, and since we ignore constants, this simplifies to O(n). This is exactly the two pointer pattern covered in this session's pattern breakdown, and recognizing this shape, rather than assuming nested loops automatically mean O(n squared), is one of the most valuable analysis skills to build early.

</details>

**Exercise 6 (stretch):** You are asked to generate every possible ordering of a set of 6 distinct items. Without writing any code, estimate roughly how many total orderings this will produce, and name the time complexity class this falls into.

<details>
<summary>Check your answer</summary>

6 factorial, which is 6 times 5 times 4 times 3 times 2 times 1, equal to 720 total orderings. This falls into O(n factorial), the factorial time complexity class, since we are generating every possible permutation. This is a strong reminder from the session that factorial time only stays practical for very small values of n.

</details>

---

## Interview Framing Notes

The bootcamp is explicitly interview focused, so here is everything from the session that speaks directly to interview performance, gathered in one place.

* **State your constraints out loud before coding.** The instructor's own process on Contains Duplicate was to first note that the array size was not fixed, meaning constant time was off the table, then check the given constraint on n to figure out roughly what complexity budget was available. Interviewers want to see this reasoning happen, not just a correct final answer.

* **Start with brute force, then optimize out loud.** Jumping straight to the most optimized answer without narrating the reasoning that got you there is a missed opportunity in an interview setting. Walking through why the brute force is too slow, and what insight lets you improve it, is itself part of what is being evaluated.

* **Know your language's rough operations per second ceiling.** The instructor used roughly 10 to the 8th power operations as a mental ceiling for C++ within a typical time limit. Having a rough sense of this for whatever language you use lets you sanity check a proposed approach against the given constraints before you even finish writing it.

* **Different approaches carry different trade offs, and interviewers may care about a specific one.** For Contains Duplicate, the sorting approach kept memory constant but cost more time complexity relative to the hash based approaches, while the hash set and hash map approaches were faster in time but used extra space. Be ready to discuss why you would pick one over another depending on what the interviewer says they care about, whether that is raw speed, memory constraints, or code simplicity.

* **Best, average, and worst case matter, and interviewers usually want worst case by default.** When someone in an interview says "what is the Big O here" without further qualification, they are almost always asking for the worst case, so default to reasoning about worst case scenarios unless told otherwise.

* **Pattern recognition, not memorized answers, is the actual underlying skill being tested.** The instructor was explicit that grinding random problems without a learning sequence builds a false sense of progress rather than real understanding. Reading constraints, noticing data types, and recognizing which broad family a problem belongs to (recursive, graph based, single array, and so on) before you even start coding is what experienced problem solvers actually do.

* **Build strong confidence in one language before spreading across several.** The instructor described starting with Python before later moving to C++ for competitive programming, specifically for speed and execution reasons. The advice given was to build fluency in your strongest language first, understanding its memory behavior and constraints deeply, rather than trying to be equally fluent everywhere at once.

---

## Terminology Glossary

**Big O notation:** a mathematical way of describing the upper bound, or worst case growth rate, of an algorithm's time or space usage as its input size grows.

**Time complexity:** a measure of how the number of operations an algorithm performs grows as its input size grows. Not a measurement of literal seconds.

**Space complexity:** a measure of how much memory an algorithm needs as its input size grows, calculated as input space plus extra auxiliary space.

**Auxiliary space:** any extra memory an algorithm uses beyond the input itself, such as additional arrays, hash maps, or the hidden memory used by a stack of recursive calls.

**Constant time, O(1):** an algorithm whose number of operations never changes regardless of input size.

**Linear time, O(n):** an algorithm whose number of operations grows directly in proportion to input size.

**Logarithmic time, O(log n):** an algorithm whose number of operations grows very slowly because the remaining input shrinks by a constant factor, typically half, at every step. The logarithm here is base 2 unless stated otherwise.

**Linearithmic time, O(n log n):** a combination of a linear pass and a logarithmic, divide and halve process, common in efficient sorting algorithms.

**Quadratic time, O(n squared):** an algorithm whose number of operations grows in proportion to the square of the input size, typically caused by two nested loops that both scale with n.

**Exponential time, O(2 to the n):** an algorithm whose number of operations doubles with every additional element in the input, commonly seen when generating every possible subset of a set.

**Factorial time, O(n factorial):** an algorithm whose number of operations equals the factorial of the input size, commonly seen when generating every possible ordering, or permutation, of a set.

**Polynomial time:** any complexity class where the input size n is raised to some fixed power, such as n, n squared, or n cubed.

**Nonpolynomial time:** complexity classes that grow faster than any fixed power of n, such as exponential and factorial time.

**Binary search:** an algorithm that repeatedly halves a sorted search space by comparing the target to a middle element, achieving O(log n) time.

**Binary search tree (BST):** a tree structure where every node's left subtree holds smaller values and right subtree holds larger values, allowing the same halving logic as binary search.

**Balanced tree:** a tree where the nodes are distributed so that each level is filled before the next begins, keeping the height, and therefore search time, close to log n.

**Skewed tree:** a degenerate binary search tree, often caused by inserting already sorted data without rebalancing, where every node becomes a child on the same side, causing the tree to behave like a straight line, meaning O(n) search time instead of O(log n).

**Two pointer pattern:** a technique using two indices moving through a structure, often one from each end or both moving forward, where the total combined movement across the whole run stays bounded by roughly the input size, giving O(n) time despite looking like nested loops.

**Set:** a data structure that stores only distinct elements, useful for checking whether something has already been encountered.

**Hash map:** a data structure that stores key and value pairs, useful for counting occurrences or looking up associated data quickly, typically in close to O(1) time per operation.

**Divide and conquer:** an algorithmic strategy that repeatedly breaks a problem into smaller pieces, solves those pieces, and combines the results. Binary search and many efficient sorting algorithms follow this strategy.

**Brute force:** the most straightforward, often least optimized, approach to a problem, typically the correct starting point before working toward a more efficient solution.

---

## Recap and What Comes Next

Today covered the foundation everything else in this bootcamp will build on: what Big O actually measures, the two formal simplification rules (ignore constants, drop lower order terms), the difference between time and space complexity, and a full walk through of the major complexity classes from constant time all the way up to factorial time. We then extracted two reusable patterns, binary search and the two pointer technique, and applied everything to a complete worked interview problem, Contains Duplicate, moving from a brute force O(n squared) approach all the way to optimized O(n) approaches using a set and a hash map.

The instructor was explicit that **space complexity will be the focus of the next session**, continuing directly from the space complexity groundwork already laid today. The bootcamp's broader roadmap for week one also centers on Big O analysis, binary search, and hash maps as the core basics before moving into more advanced structures and patterns in later weeks, so treat this document as the base layer everything else gets stacked on top of.
