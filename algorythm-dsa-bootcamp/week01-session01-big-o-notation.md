# Week 1, Session 1: Big O Notation

**Program:** Algorythm Data Structures and Algorithms Bootcamp
**Instructor:** Wise (Airbnb Engineer, DSA Coach)
**Session type:** Live lecture (Session 1 of 3 for Week 1)
**Topic:** Big O Notation, Time Complexity, Space Complexity, an introduction to Binary Search and Binary Search Trees, and a first practice problem (Contains Duplicate)

---

## How to use this file

Read this top to bottom the first time. Do not skip the analogies, they are doing more work than the formal definitions. Every code block is something you should actually type out and run yourself, not just read. The exercises at the bottom are yours to attempt before checking the answers.

---

## 1. What Big O Notation Actually Is

Big O notation is a mathematical way to describe how the runtime or memory use of an algorithm grows as the size of its input grows. It does not tell you how many seconds something takes. It tells you the shape of the growth curve.

Think of it like this. If you double the size of the input, does your algorithm's work stay the same, double, quadruple, or explode? Big O gives you a label for that shape, independent of what computer you are running it on.

The instructor's framing is worth holding onto for the whole bootcamp: **coding is the easy part, this is the thinking part.** Every time you choose a data structure, design a database query, or figure out why a feature slows to a crawl once real users hit it, you are reasoning in Big O terms whether you realize it or not.

### Why it matters in an interview

When an interviewer asks "what is the time complexity of this" and then follows up with "can you optimize it," they are not testing whether you memorized a definition. They are checking whether you actually understand what your code is doing under the hood, and whether you can reason about trade offs out loud. This is one of the most common interview questions there is, so get comfortable narrating your own complexity as you write code, not just calculating it after the fact.

---

## 2. The Three Core Rules

These three rules are the foundation for everything else in this document. Learn them cold.

### Rule 1: Ignore constants

If your algorithm does one operation per item in a list of size n, that is O(n). If it does two operations per item, that is technically 2n, but we still call it O(n), because the 2 is a constant multiplier and Big O only cares about the growth shape, not the exact count.

```python
# One operation per item -> O(n)
for item in my_list:
    print(item)

# Two operations per item -> still simplifies to O(n), not O(2n)
for item in my_list:
    print(item)
    print(item * 2)
```

Why do we drop the constant? Because as n gets huge (think millions), the difference between n and 2n stops mattering compared to the difference between n and n squared. Big O is about long term scaling behavior, not a precise operation count.

### Rule 2: Always take the worst case (the upper bound)

Big O describes the **upper bound**, meaning the worst case scenario. If part of your algorithm is O(n squared) and another part is O(n), your overall algorithm is described as O(n squared), because that is the part that will dominate and become the bottleneck as input grows.

This is also called discarding the lower order term. You keep the term that grows fastest and throw away the rest.

```python
def example(arr):
    # This nested loop is O(n^2)
    for i in arr:
        for j in arr:
            print(i, j)

    # This single loop is O(n), but it gets discarded
    # because O(n^2) already dominates
    for i in arr:
        print(i)

# Overall time complexity: O(n^2)
```

A student in the session asked what "lower bound" and "upper bound" mean here. The instructor's answer: if you have an O(n squared) chunk and an O(n) chunk in the same algorithm, the upper bound (worst case, the one that takes longer) is O(n squared), and the lower bound is O(n). Big O always reports the upper bound.

### Rule 3: n log n is multiplication, not addition

This tripped up several students in the session, so pay close attention. If an algorithm sorts a list (O(n log n)) and then does something inside a loop that also costs log n, whether you get **n log n** or **n plus log n** depends entirely on whether the two costs are nested inside each other or happening one after another.

- If one thing happens **after** another (not nested): you add them, for example O(n) + O(log n). If they are different growth rates, you would keep the larger one per Rule 2.
- If one thing happens **inside** a loop that repeats for every element (nested): you multiply them, for example O(n) times O(log n) equals O(n log n).

Concretely: sorting a list with an efficient sort algorithm costs O(n log n) all on its own, that is simply the fundamental cost of sorting, not something built from separate n and log n steps that got added together. If you then run a binary search (O(log n)) on that freshly sorted list, the two steps happen one after another, so the total becomes O(n log n) + O(log n), which simplifies down to just O(n log n) because n log n is the larger, dominating term (Rule 2 again).

Contrast that with something like: for every element in a list of size n, sort a smaller list of size n. That is O(n) multiplied by O(n log n), which equals O(n squared log n), because the sort is nested inside the loop and repeats for every single element.

---

## 3. Time Complexity vs Space Complexity

**Time complexity** measures how the number of steps your algorithm performs grows as input size grows. It is not a stopwatch measurement in seconds, it is a description of how the *work* scales. Double the input, and an O(n) algorithm roughly doubles its work, while an O(n squared) algorithm roughly quadruples its work. Scale the input 10x, and that same O(n squared) algorithm's cost balloons 100x.

**Space complexity** measures how much memory your algorithm needs to run. This breaks down into two parts:

```
space complexity = input space + auxiliary space
```

- **Input space** is the memory taken up by the data you were given (for example, the array passed into your function).
- **Auxiliary space** (also called extra space) is any additional memory your algorithm creates on its own: new arrays, hash sets, hash maps, recursive call stacks, and so on.

A key clarification from the session: if you did not create it, it does not count as auxiliary space. An array that was passed into your function as input is input space, not auxiliary space, even if you read from it. You only start counting auxiliary space once your algorithm allocates its own new memory.

If your algorithm solves a problem using only the input it was given, with zero extra data structures created, its auxiliary space is O(1), meaning constant, because you are using the same fixed amount of extra memory no matter how large the input grows.

### A quick real world contrast: sorting in place vs creating a new array

This came up directly in the session as an example: quick sort typically sorts a list **in place**, meaning it rearranges the existing array without creating a large new one, so its auxiliary space is much smaller. Merge sort, in its classic implementation, creates temporary arrays to merge sorted halves together, so it uses more auxiliary space. Both can have the same time complexity while differing meaningfully in space complexity. This is the kind of "trade off" thinking interviewers want to hear you talk through out loud.

---

## 4. The Time Complexity Ladder

Here is the ladder from fastest to slowest that the session walked through, in order.

### O(1): Constant Time

No matter how big the input gets, the algorithm always does the same fixed amount of work.

**Analogies from the session** (these are genuinely good, keep them):
- A parent hands over the entire candy bowl to a trick or treating kid in one motion. It does not matter if the bowl has 5 pieces or 5,000 pieces of candy, handing over the whole bowl is one operation.
- Knowing your own birthday. Whether you are 6 or 80 years old, you recall your birthday instantly, from memory, in one step. You are not counting up to it.
- Accessing a book on a shelf when you already know exactly where it is. You reach out, grab it, done. Compare that to walking down a street knocking on every door until you find the right house, which is not O(1), that is the O(n) case below.

```python
def get_first_element(arr):
    return arr[0]  # O(1), always exactly one operation

def get_first_three(arr):
    # Still O(1)! Fixed number of operations regardless of array size
    return arr[0], arr[1], arr[2]
```

**The critical nuance:** accessing an element **by index** is always O(1), because a computer can jump directly to any index in an array. But if you do not know the index and have to search for a value, that is no longer O(1), it becomes O(n) because you may have to check every element.

```python
# O(1): you know exactly where to look
value = arr[7]

# O(n): you don't know where it is, you have to search for it
target = 42
for value in arr:
    if value == target:
        found = True
        break
```

Another nuance raised in the session: iterating over a **fixed, unchanging collection**, like the 26 letters of the English alphabet, is also O(1), because that count never grows no matter what input the function receives. The number 26 acts as a constant, not as n. This holds true even if the alphabet string is defined outside the function and passed in as a parameter, because the size of that particular collection never changes regardless of the input.

One more important warning from the session: **O(1) is not automatically faster than O(n) in real terms.** An O(1) algorithm that always touches 1 million fixed elements is technically "constant," but an O(n) algorithm operating on a small list of 100 to 200 elements will run faster in practice. Big O describes how something *scales*, not which one wins on a specific, fixed size input. Do not confuse growth rate with raw speed on one particular case.

### O(n): Linear Time

The number of operations grows directly, one to one, with the size of the input.

```python
def print_all(arr):
    for item in arr:
        print(item)
    # Runs exactly n times for an array of size n -> O(n)
```

Nuance from the session: what if you only loop through **half** the list?

```python
def print_half(arr):
    for i in range(len(arr) // 2):
        print(arr[i])
```

This is still O(n)! Because n divided by 2 still simplifies down to O(n) once you apply Rule 1 (ignore constants). The growth rate is still linear, just with a smaller constant multiplier.

### O(log n): Logarithmic Time

This happens when, with each step, you eliminate a large chunk of the remaining input, typically by cutting it in half. As the input grows, the number of steps needed grows *much* more slowly by comparison.

**Analogy from the session:** looking up a word in a physical dictionary. You do not read every entry from A to Z (that would be O(n)). Instead you open to the middle, check whether your word comes before or after that page, and then discard the half you now know cannot contain it. Repeat. Each step throws away roughly half of what remains.

#### Binary Search, walked through in detail

Binary search is the textbook example of O(log n), but it **only works on a sorted list**. This was emphasized repeatedly in the session: you do not get to use binary search on unsorted data. If your data is not already sorted, you would have to sort it first (which costs O(n log n)), and only then could you binary search it (O(log n)) — the combined total in that scenario simplifies to O(n log n), since n log n dominates over log n by Rule 2.

```python
def binary_search(arr, target):
    low = 0
    high = len(arr) - 1

    while low <= high:
        mid = (low + high) // 2

        if arr[mid] == target:
            return mid           # Found it!
        elif arr[mid] < target:
            low = mid + 1        # Target is in the right half, discard the left
        else:
            high = mid - 1       # Target is in the left half, discard the right

    return -1  # Not found
```

Walking through why `while low <= high` uses "less than or equal to": the low and right pointers are meant to converge toward each other as you narrow down the search. Once they pass each other, you know you have searched the entire list and the element is not there. Without that stopping condition, the loop would never terminate.

A student asked why we use two pointers here at all. Binary search is one example of a broader pattern called the **two pointer technique**, which shows up across many different algorithms, not just this one.

### O(log n) for a Balanced Binary Search Tree

A binary search tree (BST) is a tree where every node has at most two children: a left child holding a smaller value, and a right child holding a larger value. Because of that ordering, every comparison you make lets you discard an entire subtree, exactly the same principle as binary search on an array, just expressed as a tree shape instead of a flat list.

```python
class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

def search_bst(node, target):
    if node is None:
        return False
    if node.value == target:
        return True
    elif target < node.value:
        return search_bst(node.left, target)
    else:
        return search_bst(node.right, target)
```

**Why a balanced tree gives you log n:** a fully balanced binary tree doubles the number of nodes at every level going down. Level 0 has 1 node, level 1 has 2, level 2 has 4, level 3 has 8, and so on, level k holds 2 to the power of k nodes. If a tree has height h, it holds roughly 2 to the power of h nodes total. Flip that relationship around: if n is the total node count, then h equals log base 2 of n. Every time you double the number of nodes, you only add a single extra level of height, and that "doubling relationship" is exactly what a logarithm captures.

**The three cases you need to know for interviews:**

| Case | Shape | Complexity |
|---|---|---|
| Best case | Perfectly full and balanced tree | O(log n) |
| Average case | Randomly inserted, roughly balanced | O(log n) |
| Worst case | Skewed tree (every node only has one child, forming basically a line) | O(n) |

A **skewed** tree is the worst case because you lose the "discard half" advantage entirely. If every node only ever has a right child, for example, the tree degenerates into what is functionally a linked list, and searching it requires checking every single node, exactly like a plain linear search. If asked in an interview and the problem does not specify whether a tree is balanced, state both possibilities: "in the best case with a balanced tree this is O(log n), but in the worst case with a skewed tree this degrades to O(n)."

### O(n log n): Linearithmic Time

This is the standard cost of efficient, comparison based sorting. Merge sort, quick sort (average case), and heap sort all run in O(n log n). This is described as roughly the fastest general purpose sorting speed achievable, and is what most programming languages use internally for their built in sort functions.

Do not think of n log n as "an n part plus a log n part added together." It is genuinely its own single complexity class, arising because you are doing log n work, n separate times (or vice versa), nested together, not sequentially.

### O(n squared): Quadratic Time, Nested Loops

This happens when, for every single item in your input, you do another full pass over the input (or something proportional to it).

```python
def print_grid(arr):
    for i in arr:          # outer loop: runs n times
        for j in arr:      # inner loop: runs n times, for EACH outer iteration
            print(i, j)
# Total operations: n * n = n^2
```

If the array has 6 elements, this prints 6 times 6, which is 36 operations total. If the outer loop represents rows and the inner loop represents columns, note carefully: **the outer loop is the rows, the inner loop is the columns** (a detail a couple of students in the session mixed up).

**Important nuance the instructor called out as commonly misunderstood: not every nested loop is automatically O(n squared).** What actually matters is how many times the *innermost* line of code executes in total across the whole run, not simply how many loops are nested inside each other.

```python
# Case 1: inner loop always runs a FIXED number of times (say, 5)
# regardless of n -> this is O(5n), which simplifies to O(n), NOT O(n^2)
def fixed_inner(arr):
    for item in arr:              # runs n times
        for i in range(5):        # always exactly 5, never grows with n
            print(item, i)
```

```python
# Case 2: two DIFFERENT sized inputs -> this is O(n * m), not O(n^2)
# n^2 specifically means the two dimensions are the SAME size
def different_sizes(list_a, list_b):
    for a in list_a:       # runs n times
        for b in list_b:   # runs m times, where m != n
            print(a, b)
# Total: n * m. Only becomes n^2 if n and m happen to be equal.
```

```python
# Case 3: the "two pointer" pattern that looks nested but is actually O(n)
# The inner pointer never resets back to zero for each outer iteration,
# it only ever moves forward across the ENTIRE run, at most n times total.
def sliding_window_example(arr):
    left = 0
    for right in range(len(arr)):
        while (right - left) > 2:   # keeps window at most 3 wide
            left += 1
        print(arr[left:right + 1])
# Even though this LOOKS like a nested loop, left only ever advances
# forward across the whole run, never resetting, so total work is O(n)
```

The distinguishing question to always ask yourself when you see a nested loop: **does the inner loop reset back to the start for every single outer iteration?** If yes, and both loops are the same size, you get O(n squared). If the inner pointer just keeps crawling forward across the entire run without resetting, the total work stays O(n) even though the code visually looks nested.

### O(2 to the n): Exponential Time, Subsets

This shows up in problems asking you to generate every possible subset of a set of elements, a classic backtracking pattern. At every single element, you make a binary decision: include it in the current subset, or do not include it. Since there are n elements and 2 choices at each one, the total number of possible subsets is 2 multiplied by itself n times, written as 2 to the power of n.

```python
def all_subsets(nums):
    result = []

    def backtrack(start, current):
        result.append(current[:])   # record the subset built so far
        for i in range(start, len(nums)):
            current.append(nums[i])       # choose to include nums[i]
            backtrack(i + 1, current)     # explore further with it included
            current.pop()                 # backtrack: undo choosing it
            # the "don't include it" branch is implicitly the next loop iteration

    backtrack(0, [])
    return result
```

For a 3 element list `[1, 2, 3]`, walking the include or exclude decision tree for every element produces 8 total subsets, which is 2 to the power of 3. Space complexity for storing every subset is also O(2 to the n), since you end up holding that many subsets in memory.

This is genuinely one of the more advanced complexity classes covered this early, so if the tree diagram does not click immediately, that is expected. It will get a fully dedicated lesson later in the bootcamp.

### O(n!): Factorial Time, Permutations

This shows up when order matters, meaning you are counting all possible *orderings* of n elements, not just which elements are included. For 3 elements, you have 3 choices for the first slot, then 2 remaining choices for the second slot, then only 1 choice left for the final slot: 3 times 2 times 1 equals 6, which is exactly 3 factorial. The instructor was direct about this one: factorial time problems are rare to encounter in real interviews, so do not over invest in this one compared to the others.

---

## 5. Terminology Glossary

Refer back here whenever a term feels unfamiliar as you go through future sessions.

- **Big O notation**: a way to describe the upper bound, meaning worst case growth rate, of an algorithm's time or space usage as input size grows.
- **Time complexity**: how the number of operations an algorithm performs scales as input size grows.
- **Space complexity**: how much memory an algorithm needs, made up of input space plus auxiliary space.
- **Auxiliary space**: extra memory an algorithm creates on its own, separate from the input it was given.
- **Upper bound**: the worst case scenario, the slowest growing part of your algorithm. This is what Big O reports.
- **Lower bound (in this context)**: a less significant, faster growing part of an algorithm that gets discarded once combined with a slower, dominating part.
- **Constant**: a fixed multiplier or fixed count that does not change as input size grows. Big O ignores constants (Rule 1).
- **In place**: an algorithm that rearranges or solves using its existing input storage rather than allocating a large new data structure, keeping auxiliary space low.
- **Two pointer technique**: an algorithmic pattern using two index variables (often called left and right, or low and high) that move through data, frequently used to avoid nested loops.
- **Balanced binary search tree**: a BST where the left and right subtrees at every node are roughly equal in size or height, keeping search operations at O(log n).
- **Skewed tree**: a binary search tree where nodes mostly only have one child, causing it to behave like a linked list and degrading search to O(n).
- **Backtracking**: an algorithmic technique that explores one possible path fully, then reverses (backtracks) to try alternate choices, commonly used for subsets and permutations.
- **Hash set**: a data structure that stores unique values with no guaranteed order, offering roughly O(1) membership checks (does this value exist in the set).
- **Hash map**: like a hash set, but stores key value pairs rather than just standalone values.

---

## 6. Worked Practice Problem: Contains Duplicate

**Problem statement (LeetCode style, as given in the session):**
Given an integer array `nums`, return `true` if any value appears at least twice in the array. Return `false` if every element is distinct.

The session deliberately worked through this problem starting from the naive, less efficient solution first, then optimizing it, which is exactly the process an interviewer wants to see out loud, so the three approaches below are presented in that same order.

### Approach 1: Brute force, nested loops

```python
def contains_duplicate_brute_force(nums):
    for i in range(len(nums) - 1):
        for j in range(i + 1, len(nums)):
            if nums[i] == nums[j]:
                return True
    return False
```

- **Time complexity:** O(n squared), because of the nested comparison of every pair.
- **Space complexity:** O(1), no extra data structure was created, only the two loop indices.
- A subtle bug that came up live in the session: the inner loop's range needs to correctly stop one before the end (`len(nums) - 1` for the outer range, or careful use of `i + 1` as the inner starting point) to avoid comparing an element against itself or going out of bounds. Watch your off by one boundaries carefully here.

### Approach 2: Hash set built incrementally, one pass

```python
def contains_duplicate_hash_set(nums):
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False
```

- **Time complexity:** O(n), a single pass through the list. Checking membership in a set (`num in seen`) and adding to a set are both O(1) on average.
- **Space complexity:** O(n) in the worst case (a list with no duplicates at all would end up storing every element in the set).
- A bug that happened live in the session is worth learning from directly: an earlier, broken version of this built the set from the *entire* input array first, and then checked membership against that same complete set while looping, which meant the very first element checked would always immediately register as "found," since it was already sitting in the set from the start. The fix is to build the set up incrementally, one element at a time, checking *before* adding on each iteration, exactly as shown above.

### Approach 3: Compare set length to array length

```python
def contains_duplicate_length_compare(nums):
    return len(set(nums)) != len(nums)
```

This is the most concise version. Since a set cannot contain duplicate values by definition, converting the whole array into a set and comparing its length against the original array's length tells you immediately whether any duplicates existed.

- **Time complexity:** O(n), because building a set from an existing array requires iterating over every element once internally.
- **Space complexity:** O(n), for the same reason as Approach 2.

### Comparing the three approaches

| Approach | Time | Space | Notes |
|---|---|---|---|
| Brute force, nested loops | O(n squared) | O(1) | Simple, but slow to scale. Trades time for space. |
| Incremental hash set | O(n) | O(n) | The standard efficient answer. |
| Set length comparison | O(n) | O(n) | Same efficiency, more concise code. |

This table itself is a great habit to build. In the session, the instructor was explicit that discussing this exact kind of time versus space trade off out loud is exactly what interviewers are listening for.

---

## 7. Interview Framing Tips Mentioned in the Session

A handful of genuinely useful interview process comments came up during the Q&A portions, worth capturing separately from the pure algorithm content:

- **Ask clarifying questions before you start coding.** Is the input sorted or not? What data type are we dealing with? Are there constraints on time or space I should know about? If a problem does not explicitly state the input is sorted, assume it is not sorted, even if the given example happens to look sorted.
- **You are allowed to use your language's standard library.** Most interviewers expect you to use built in data structures unless they explicitly restrict you (for example, "solve this in O(1) space" or "you cannot use a hash map").
- **If a language is missing a data structure you need, you can say so out loud and assume it exists.** For example, if using JavaScript and you need a priority queue, you can tell the interviewer "assuming a priority queue data structure exists here" rather than being expected to build one from scratch on the spot, unless they specifically ask you to.
- **Treat the interview as two way.** You are also evaluating whether you want to work with that team. Asking thoughtful clarifying questions demonstrates communication skill and curiosity, not just correctness.
- **Know your chosen language's non built in data structure libraries.** For example, Python does not have a priority queue built directly into the base language, you need the `heapq` module. Know what tools exist and where to reach for them.

---

## 8. Brush Up Coding Exercises

Attempt these yourself before checking the answers underneath each one. Try to state the time and space complexity of your own solution out loud (or in a comment) before moving to the next problem, exactly like the habit described in the session.

### Exercise 1 (Easy): Classify the complexity

What is the time complexity of the function below?

```python
def mystery_one(arr):
    total = 0
    for num in arr:
        total += num
    return total
```

<details>
<summary>Answer</summary>

O(n) time. A single pass through the array, one operation per element. Space complexity is O(1), since `total` is a single fixed variable regardless of array size.
</details>

### Exercise 2 (Easy): Spot the nested loop trap

Is the function below O(n squared) or O(n)? Explain why using the "does the inner loop reset" question from Section 4.

```python
def mystery_two(arr):
    for i in arr:
        for j in range(3):
            print(i, j)
```

<details>
<summary>Answer</summary>

O(n). The inner loop always runs exactly 3 times no matter how big `arr` is, it is a fixed constant, not something that grows with n. So this is O(3n), which simplifies to O(n) once you drop the constant (Rule 1).
</details>

### Exercise 3 (Medium): Implement binary search from scratch

Without looking back at Section 4, write your own `binary_search(arr, target)` function. Test it against a sorted list of your choosing. State its time and space complexity in a comment above your function.

<details>
<summary>Answer</summary>

```python
def binary_search(arr, target):
    # Time: O(log n), Space: O(1)
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1
```
</details>

### Exercise 4 (Medium): Contains Duplicate, from memory

Without scrolling back to Section 6, write the incremental hash set solution to Contains Duplicate from scratch. Then write the one line set length comparison version.

<details>
<summary>Answer</summary>

```python
# Incremental hash set version
def contains_duplicate(nums):
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False

# One line version
def contains_duplicate_v2(nums):
    return len(set(nums)) != len(nums)
```
</details>

### Exercise 5 (Harder, stretch): Identify the complexity of a combined operation

You are given an unsorted list. You need to find out whether a specific target value exists in it, using a binary search approach. What is the total time complexity of this entire process, including any steps required to make binary search possible? Explain your answer using Rule 2 and Rule 3 from Section 2.

<details>
<summary>Answer</summary>

You cannot binary search an unsorted list directly, so you must first sort it, which costs O(n log n), and only then run binary search on it, which costs O(log n). These two steps happen one after another (not nested), so you would technically add them: O(n log n) + O(log n). Applying Rule 2 (always keep the dominating term, discard the smaller one), this simplifies down to just O(n log n), since n log n grows faster than log n alone and therefore dominates the total.
</details>

### Exercise 6 (Harder, stretch): Subsets by hand

Draw out, on paper or in a text editor, the full include or exclude decision tree for the list `[1, 2]` (two elements only). How many total subsets do you end up with? Confirm it matches 2 to the power of 2.

<details>
<summary>Answer</summary>

Four total subsets: `[]`, `[1]`, `[2]`, `[1, 2]`. That matches 2 to the power of 2, which equals 4, exactly as expected for a 2 element input.
</details>

---

## 9. Recap and How This Connects Forward

Today's session covered:

- What Big O notation is and why engineers reason in it constantly, not just in interviews.
- The three core rules: ignore constants, always report the worst case upper bound, and understand when combined complexities multiply versus when they add.
- The full time complexity ladder from O(1) up through O(n!), with an analogy and worked code example for each.
- Space complexity, broken into input space and auxiliary space, plus a real world example contrasting in place sorting against sorting that requires a new array.
- A first, gentle introduction to binary search and binary search trees, including the important nuance that a skewed tree degrades to O(n).
- A full walk through of the Contains Duplicate problem across three different approaches, moving from brute force to optimal, exactly the kind of iterative problem solving process interviewers want to see.

**Next session (per the instructor)** will go deeper into space complexity specifically, and based on direct feedback from the group, will likely include more worked examples of actually calculating time complexity from scratch on new problems, since several students noted that the algorithms made sense individually but computing the final complexity notation still felt shaky. Keep that in mind as a specific thing to watch for and practice actively in the next transcript.

This entire ladder, O(1), O(log n), O(n), O(n log n), O(n squared), O(2 to the n), O(n!), is the vocabulary the rest of the twelve weeks will be spoken in. Every future data structure and algorithm you learn (linked lists, hashmaps, trees, graphs, sorting algorithms, dynamic programming) will get evaluated against this exact same scale. Getting genuinely comfortable with this session before moving forward is worth the extra time.
