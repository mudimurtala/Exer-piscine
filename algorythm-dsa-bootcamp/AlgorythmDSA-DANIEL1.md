<!-- week02_session02_space_complexity.md -->

# Week 2, Session 2: Space Complexity

**Topic:** Space Complexity (how much memory an algorithm needs)
**Instructor:** Daniela
**Bootcamp:** Algorithms Bootcamp, Week 2 (building directly on Week 1's time complexity and Big O notation)

---

## Before You Start Reading

This session picks up right where time complexity left off. If time complexity answers "how fast," space complexity answers "how much memory." They are measured with the same Big O language, but they are not the same thing, and mixing them up is a common trap early on. Read this whole document once, then come back to it later (Daniela specifically recommends this) when you learn breadth first search and other data structures, because several ideas here (like the reference sheet and the tree diagram for subsets) are meant to be revisited, not memorized in one sitting.

---

## Concept Breakdown

### What space complexity actually measures

Time complexity tells you how the number of operations grows as your input grows. Space complexity tells you how the amount of memory your program uses grows as your input grows. You need both, because algorithms trade one for the other constantly. Some solutions burn more time to save memory. Some solutions burn more memory to save time. Neither is automatically "better," it depends on the constraints of the problem you are solving.

Why would you ever care about memory if your laptop has plenty of it? Two reasons come up in real engineering and in interview style problems:

1. You might be running on constrained hardware (embedded systems, older machines, systems with strict memory budgets). Use too much memory and your program simply crashes with something called a memory limit exceeded error (often abbreviated MLE on competitive programming and interview practice sites).
2. On sites like LeetCode, and in real interviews, you are explicitly graded on space efficiency, not just correctness and speed.

### The two components of space complexity

Every program's space complexity is made of two parts:

**Input space**, the memory needed just to hold the data you were given. If you are handed an array of n numbers, storing that array already costs you memory, whether or not you do anything else with it.

**Auxiliary space** (also called extra space), the memory your program uses beyond what it was handed. Any additional variables, arrays, lists, or recursive call stacks you create count here.

```
space complexity = input space + auxiliary space
```

Here is the part that trips people up early: when someone asks "what is the space complexity of your algorithm," they almost always mean the auxiliary space, not the total. This is because the input space is usually obvious and unavoidable, it is set by the problem itself, not by your solution. So if you loop through an array of n elements using one temporary variable, the input technically costs O(n) memory, but your algorithm's space complexity is described as O(1), because that one variable is the only extra memory your code introduces.

### How memory allocation actually works behind the scenes

This is the part of the session that goes beyond what most people assume, and it is worth internalizing carefully because it explains why dynamic arrays behave the way they do.

Naive intuition says memory grows one unit at a time as you use it. That is not how it works. When you create a fixed size array, the compiler reserves the entire block of memory immediately, at the moment you declare it, regardless of whether you have filled it yet. If you declare an array of 100 elements plus one extra variable, your program has already claimed 101 units of memory the instant that line runs, even if the array is still empty. This upfront reservation is called allocation.

This matters for how you reason about space complexity: you count memory based on what your code claims it needs, not based on how much of it happens to be filled with meaningful values at any given moment.

### Dynamic arrays and why doubling matters

Languages like C++ often make you declare a fixed size array upfront:

```cpp
// C++: you must state the size in advance
int myArray[100];
```

Python style variant, using the equivalent idea in a language with static sizing (not actually Python code, just the conceptual parallel):

```python
# Python does not require this at all, shown here only for contrast
# In C++ style languages you commit to a size before filling it
my_array = [None] * 100
```

Python lists (and similarly, dynamic arrays or vectors in other languages) do not require you to declare a size. You start empty and append over time:

```python
my_list = []
my_list.append(1)
my_list.append(2)
```

So how does the underlying memory get managed if you never told it how big to be? The naive guess is that it adds memory one unit at a time as you append, or adds a fixed chunk like plus 100 each time it runs out of room. Both of those turn out to be bad designs. If you add a fixed chunk of, say, 100 units every time you run out of space, and your list eventually grows to a million elements, you end up reallocating roughly ten thousand times. Each reallocation is expensive because it involves copying everything that already existed into a new, bigger block of memory, which itself takes linear time relative to how much is being copied. Doing that ten thousand times makes your program slow.

The actual solution used by dynamic arrays (including Python's list under the hood) is to double the capacity every time the current block fills up. Start with a capacity of one. Once that is used, double to two. Once that is used, double to four, then eight, then sixteen, and so on.

This might look wasteful at first glance, but it is the opposite. Because doubling grows so fast, you only need roughly thirty doubling steps to reach about one billion elements. Compare that to the fixed chunk approach, which would need about ten million steps of plus 100 to reach the same size. Far fewer reallocations means far less total copying, even though each individual reallocation copies more.

In the worst case, if you append n times, a doubling strategy uses at most roughly 2 times n units of memory, because right before a doubling event you were using close to half of your current capacity, and the new capacity becomes double the old one. This "worst case looks like double the final size, but the total copying work stays proportional to n across the whole sequence of appends" is the mechanism behind an idea you will meet later called amortized time, where a rare expensive operation gets "spread out" and averaged across many cheap ones so that the overall cost per operation still comes out linear.

Reserving memory in advance, if your language allows it (C++'s `reserve` function is one example), skips the doubling process entirely because you are telling the compiler the final size upfront. It can make your code a little faster in practice, but it does not change the underlying space complexity classification, since you would end up at the same linear space either way. No interview or LeetCode problem actually requires you to do this, it is a minor performance tweak, not a correctness requirement.

---

## Terminology Callouts

**Space complexity**: A measurement of how much memory an algorithm needs, expressed as a function of input size, using the same Big O notation you use for time complexity.

**Memory limit exceeded (MLE)**: An error that occurs when a program tries to use more memory than the system allows. Think of it like trying to pour more water into a cup than the cup can hold, the excess has nowhere to go and the system fails.

**Input space**: The memory required just to store the data you were given as input to the problem.

**Auxiliary space (extra space)**: Any memory your algorithm uses beyond the input itself, such as new variables, new arrays, or recursive call stacks. When people casually say "space complexity" of an algorithm, they almost always mean this part specifically.

**Allocation**: The moment a program reserves a block of memory. Daniela's analogy: imagine reserving 100 seats at a restaurant the moment you book, even if only one guest has arrived so far. The seats are already claimed, whether or not anyone is sitting in them yet.

**Dynamic array**: A list like structure (Python's `list`, C++'s `vector`) that can grow as you add elements, without you having to declare a fixed size upfront, unlike a static array.

**Static array**: An array whose size must be fixed at the moment it is created, common in languages like C++ where you write something like `int arr[100];` and cannot change that size later without creating a new array.

**Doubling strategy (dynamic array growth)**: The technique dynamic arrays use internally, where the underlying memory block doubles in size every time it fills up, rather than growing by a fixed amount each time. This keeps the total number of expensive reallocation events small even as the list grows very large.

**Constant space, O(1)**: Memory usage that does not depend on the size of the input at all. Whether your input has one element or a billion, you use the same fixed amount of extra memory.

**Linear space, O(n)**: Memory usage that grows directly in proportion to the input size n. Double the input, roughly double the memory.

**Quadratic space, O(n squared)**: Memory usage that grows in proportion to n multiplied by itself, common when you build a two dimensional structure like a grid or matrix sized n by n.

**Exponential space, O(2 to the power of n)**: Memory usage that roughly doubles with every additional input element. This shows up in problems involving all subsets of a set, since each new element doubles the number of possible subsets.

**Subset**: Any selection of elements from a set, including the empty selection and the full set itself. For a set of three items, there are 2 to the power of 3, which is eight, possible subsets, because each item independently is either included or excluded.

**In place**: A requirement that a solution modify the given input directly, without allocating a new array, list, or other structure to hold the result. This usually forces you toward constant extra space.

**Two pointers**: A technique where you track two index positions (pointers) moving through a structure, often used to rearrange or scan data using only a fixed number of extra variables, which is why it pairs naturally with in place, constant space requirements.

**Amortized (mentioned as a preview)**: A way of describing cost that averages an occasional expensive operation across many cheap ones, so the effective cost per operation still looks linear or constant over the long run. This is why doubling a dynamic array's capacity still results in linear total copying work, even though individual doubling events are relatively costly.

**Hash table**: A data structure that lets you insert, look up, and delete elements in average constant time, O(1), by using a hashing function to map keys to memory locations. It is extremely versatile but cannot preserve ordering information like "which element was added first," which matters for problems like breadth first search.

**Queue**: A structure that lets you add elements and always retrieve the earliest added element first (first in, first out). Needed for breadth first search precisely because a hash table cannot tell you which element arrived first.

**Big O of V plus E**: A space or time expression used specifically for graphs, where V is the number of vertices (nodes) and E is the number of edges (connections). It uses two separate variables instead of one because a graph's shape depends independently on how many nodes it has and how many connections between them exist, which is why graphs cannot be placed cleanly on a single variable Big O comparison chart.

---

## Pattern Breakdown: Classifying Space Complexity by Growth Rate

This is the single most transferable skill from this session. You will meet dozens of different problems, but almost all of their space costs fall into one of a small handful of growth shapes. Once you can recognize which shape a piece of code produces, you do not need to re derive it from scratch every time.

### The general signal to look for

Ask yourself: as the input size n grows, does the number of extra variables or extra storage locations my code creates:

- stay exactly the same no matter what n is → constant
- grow directly proportional to n → linear
- grow proportional to n multiplied by itself → quadratic
- grow proportional to n multiplied by itself multiple times (three dimensional structures, and so on) → cubic, or higher powers
- roughly double for every additional input element → exponential

The general shape you are checking is: "for every unit increase in n, how many additional units of memory does my structure need?" If the answer is "always the same fixed number," you are constant. If the answer is "one more unit of memory per one more input element," you are linear. If the answer is "an entire new row or column's worth of memory per one more input element," you are quadratic. If the answer is "the total memory needed roughly doubles," you are exponential.

### Constant space, O(1)

**Signal:** Your code uses a fixed, small number of variables that never grows no matter how big the input becomes. Looping through an array and reusing a single temporary variable is the classic example, because that one variable gets reassigned each iteration rather than accumulating.

**General template:**

```python
def process(items):
    total = 0          # one fixed variable
    for item in items:
        total += item  # reused, not grown
    return total
```

**How the transcript's example fits this pattern:** Daniela showed a loop that printed each element of a list using a single reused variable named `elements`. Even though the loop runs n times, only one variable ever exists at a time, so the auxiliary space stays O(1) regardless of how large the input list is.

### Linear space, O(n)

**Signal:** Your code builds a new structure (array, list, set) whose final size is directly tied to the input size, usually because you are storing one new item per input item.

**General template:**

```python
def build_result(items):
    result = []
    for item in items:
        result.append(transform(item))
    return result
```

**How the transcript's example fits this pattern:** The reversed array example built a brand new list containing every element of the input array, just in reverse order. Since the new list ends up holding exactly n elements, the space complexity is O(n). Daniela's broader point: O(n) space is extremely common and generally acceptable, because most problems already require storing the input somewhere, so matching that with O(n) auxiliary space is often the practical ceiling rather than a red flag.

### Quadratic space, O(n squared)

**Signal:** Your code builds a two dimensional structure, a grid, table, or matrix, whose dimensions both scale with n. Watch for phrases like "for each element, store something about every other element" or explicit 2D array creation sized n by n.

**General template:**

```python
def build_matrix(n):
    matrix = [[0] * n for _ in range(n)]
    return matrix
```

**How the transcript's example fits this pattern:** The identity matrix example created an n by n two dimensional list. With n equal to four, that is 16 total cells. With n equal to eight, that jumps to 64 cells. Notice the growth rate compared to linear: going from n equal to four to n equal to eight only doubled n, but it quadrupled the total memory used. That accelerating gap is the signature of quadratic growth. Practically, Daniela's rule of thumb was that once n climbs past roughly 5000, quadratic space starts becoming impractical.

### Exponential space, O(2 to the power of n)

**Signal:** The problem explicitly involves generating or storing every possible subset, combination, or binary choice sequence over n items. Each item independently contributes a factor of two (include it or do not), so the total possibilities multiply by two per item.

**General template (generating all subsets):**

```python
def all_subsets(items):
    subsets = [[]]
    for item in items:
        # for every subset we already have, create a new version that also includes item
        subsets += [current + [item] for current in subsets]
    return subsets
```

**How the transcript's example fits this pattern:** Daniela used a set of three items (A, B, C) and showed all eight subsets, matching 2 to the power of 3. She also reframed this as a binary tree with n levels, where each node branches into two children, one representing "did not pick this element" and one representing "picked this element." The number of leaf nodes at the bottom level is 2 to the power of n, which is exactly the number of subsets. A critical nuance she raised: this exponential space cost only applies if you actually store every subset somewhere, such as in a 2D array. If your task is only to print each subset as you generate it, without saving them all, your auxiliary space drops to O(n), because at any given moment you only need enough memory to hold one subset in progress (at most n elements), not all of them simultaneously.

### Quick reference: growth rate summary

| Space class | Big O | Rough size where it starts becoming risky (per Daniela) |
|---|---|---|
| Constant | O(1) | Never, this scales to any input size |
| Linear | O(n) | Around 10 million and beyond |
| Quadratic | O(n squared) | Around 5000 and beyond |
| Cubic | O(n cubed) | Around 500 and beyond |
| Exponential | O(2 to the power of n) | Around 20 to 25 and beyond |

The practical use of this table: before you write a single line of code for a solution, look at the input size limit given in the problem. If n can be as large as 100,000 (a very common LeetCode constraint), you already know a quadratic or worse solution is doomed to fail on memory or time, and you should not waste effort implementing it. This lets you filter out bad approaches before investing coding time in them.

---

## Worked Code Examples

### The two pass "move zeroes" solution (from the live walkthrough)

Daniela solved the problem "Move Zeroes" live: given an integer array, move all zeros to the end while keeping the relative order of the non zero elements, and do it in place, without allocating a new array. The array size constraint given in the problem was up to roughly 10,000 elements.

She first walked through a version that technically solves the problem but violates the in place constraint, by building a second list:

```python
def move_zeroes_first_attempt(nums):
    result = []
    for element in nums:
        if element != 0:
            result.append(element)
    for element in nums:
        if element == 0:
            result.append(element)
    # this produces the correct values, but it is a new list,
    # which breaks the "do not make a copy" requirement of the problem
    return result
```

This runs in O(n) time (two separate passes over n elements, and constants are dropped in Big O, so "two passes" and "four passes" both simplify to O(n)) and O(n) space, since the new `result` list ends up holding n elements. Daniela was explicit that although this is a working solution, it does not satisfy the actual constraint of the problem, since it creates a second array rather than modifying `nums` directly.

She then began building toward a better approach during the session (counting the zeros first, and considering two pointers), and confirmed that a full O(n) time, O(1) space in place solution exists using the two pointer technique, but explicitly deferred the complete walkthrough to the next class due to time. No bug needed correcting in the code she did complete, since she was explicit that the two pass version above was intentionally shown as the "what not to submit" version.

### C++ reserve example (from a student question)

A student asked about the advantage of pre reserving memory for a data structure in C++:

```cpp
// C++: telling the compiler in advance how much space we expect to need
std::vector<int> nums = {1, 2, 3, 4, 5};
nums.reserve(100);
```

Python equivalent concept (Python lists do not expose a direct reserve call, since they manage this doubling internally, but this is the closest conceptual parallel: pre sizing a structure instead of growing it incrementally):

```python
# Python list, conceptually similar intent to reserve,
# though this creates actual elements rather than just reserving capacity
nums = [1, 2, 3, 4, 5]
nums.extend([0] * 95)  # not a true equivalent, shown only to illustrate the idea
```

Daniela's clarification: reserving in advance can make your code a little faster by skipping the repeated doubling steps, but it does not change the final space complexity classification. Whether you reserve upfront or let the array grow through doubling, you still end up at O(n) total space once everything is stored. She was clear that no interview or LeetCode problem will ever require you to use `reserve`, it is purely a minor performance optimization, not something that changes correctness or Big O class.

---

## External Code Snippets (Additional Reinforcement)

These were not part of the transcript, they are extra examples built to reinforce the same four space complexity classes using slightly different scenarios than the ones Daniela used.

**Constant space, a different scenario (finding the maximum value):**

```python
def find_max(numbers):
    current_max = numbers[0]  # just one variable, reused every iteration
    for number in numbers:
        if number > current_max:
            current_max = number
    return current_max
```

No matter how long `numbers` is, this uses exactly one extra variable. This is constant space, and it pairs naturally with the "reused temporary variable" signal from the pattern breakdown.

**Linear space, a different scenario (building a frequency count):**

```python
def count_frequencies(words):
    frequency = {}
    for word in words:
        frequency[word] = frequency.get(word, 0) + 1
    return frequency
```

In the worst case, every word is unique, so `frequency` ends up holding n separate keys. That is O(n) auxiliary space, the same class as the reversed array example, even though this uses a dictionary instead of a list.

**Quadratic space, a different scenario (all pairs distance table):**

```python
def all_pairs_table(points):
    n = len(points)
    distances = [[0.0] * n for _ in range(n)]
    for i in range(n):
        for j in range(n):
            distances[i][j] = euclidean_distance(points[i], points[j])
    return distances
```

Just like the identity matrix example, this allocates an n by n grid, so it is O(n squared) space, even though the underlying task (comparing every point to every other point) is conceptually different from building an identity matrix.

**Exponential space, a different scenario (generating all binary strings of length n):**

```python
def all_binary_strings(n):
    if n == 0:
        return [""]
    smaller = all_binary_strings(n - 1)
    return [s + "0" for s in smaller] + [s + "1" for s in smaller]
```

Each additional character doubles the number of strings, exactly mirroring the "each element either picked or not picked" logic from the subsets example. This confirms that exponential space is not unique to subsets specifically, it shows up any time a problem has n independent binary choices.

---

## Practice and Brush Up Snippets

Try completing these yourself before checking any explanation. They are meant to be typed and tested, not just read.

**1. Fill in the blank to make this constant space:**

```python
def sum_array(numbers):
    total = ____
    for number in numbers:
        total += number
    return total
```

**2. Complete this function so it builds a new list containing only the even numbers (this will end up being linear space, think about why):**

```python
def evens_only(numbers):
    result = []
    for number in numbers:
        if ____:
            result.append(number)
    return result
```

**3. This function is supposed to build an n by n grid of zeros. Fix the bug (hint, check what happens to the inner list):**

```python
def build_grid(n):
    row = [0] * n
    grid = [row for _ in range(n)]  # bug is here
    return grid
```

**4. Write a one line change that would turn this exponential subset generator into an O(n) space version, assuming you only need to print each subset rather than store them all:**

```python
def all_subsets(items):
    subsets = [[]]
    for item in items:
        subsets += [current + [item] for current in subsets]
    return subsets  # currently stores everything, change this to print instead
```

---

## Exercises

### Easy

**Exercise 1:** What is the space complexity (auxiliary space specifically) of the following function?

```python
def double_values(numbers):
    doubled = []
    for number in numbers:
        doubled.append(number * 2)
    return doubled
```

<details>
<summary>Check your answer</summary>

O(n). The function creates a new list `doubled` that grows to hold exactly one entry per input element, so the auxiliary memory scales directly with the input size n.

</details>

**Exercise 2:** True or false, and explain briefly: "If a function's input space is O(n), the function's overall space complexity must be described as O(n)."

<details>
<summary>Check your answer</summary>

False in the way space complexity is usually described in practice. When people describe "the space complexity of an algorithm," they almost always mean the auxiliary space specifically, not the combined total including input space. A function could receive O(n) input but only use O(1) extra memory beyond that input, in which case its space complexity would be described as O(1), even though the total memory footprint including input is technically O(n).

</details>

### Medium

**Exercise 3:** A problem gives you an array with n up to 100,000 and asks you to solve it. A classmate suggests an approach that builds a two dimensional n by n table to solve it. Should you implement this approach? Why or why not?

<details>
<summary>Check your answer</summary>

No, you should not implement it. With n up to 100,000, an O(n squared) space solution would require on the order of 10 billion units of memory, far beyond what any realistic system allows, and it would fail with a memory limit exceeded error. Based on the general size thresholds discussed in the session, quadratic space starts becoming impractical once n passes roughly 5000, and 100,000 is far past that. You should look for an alternative approach, likely one using O(n) or O(1) auxiliary space, before writing any code.

</details>

**Exercise 4:** Explain, in your own words, why a dynamic array's doubling strategy results in roughly 2 times n memory usage in the worst case, rather than something much larger.

<details>
<summary>Check your answer</summary>

Right before any doubling event occurs, the array was using close to half of its new, doubled capacity (since doubling happens exactly when the previous capacity fills up). So immediately after doubling, at most about half of the newly allocated space is actually filled with real data, meaning the total allocated capacity is at most roughly double the number of elements actually stored. This is why the worst case space usage is bounded by roughly 2 times n rather than growing without limit.

</details>

### Hard (stretch)

**Exercise 5:** You are asked to generate all subsets of a set of n elements, but instead of printing each subset immediately, your task requires you to return them all as a single list of lists, ready for further processing later. What is the space complexity, and is there any way to avoid the exponential cost while still returning all subsets as required?

<details>
<summary>Check your answer</summary>

The space complexity is O(2 to the power of n), because there are 2 to the power of n total subsets, and the requirement to return them all as a list of lists forces you to store every single one simultaneously. Unlike the print only version, there is no way to avoid this exponential cost while still satisfying the stated requirement, since the output itself inherently contains 2 to the power of n items. This is an example of a case where the exponential space is not a flaw in your algorithm, it is a direct consequence of what the problem is asking you to produce. The lesson from the session is important here: before assuming you have a bad solution, check whether the problem's required output itself is exponential in size, in which case no algorithm could do better.

</details>

**Exercise 6:** A friend proposes using a hash table to speed up a problem that fundamentally requires building and storing an n by n grid of relationships between every pair of elements. Would using a hash table change the space complexity? Explain why or why not.

<details>
<summary>Check your answer</summary>

No, it would not change the space complexity. If the underlying task genuinely requires storing information about every pair among n elements, that is n squared distinct pieces of information no matter what structure you use to hold them, an array, a hash table, or anything else. Inserting n squared items into a hash table, even though each individual insert is average case O(1), still totals O(n squared) insertions overall, and O(n squared) space to hold them. A hash table changes access speed characteristics, it does not reduce the fundamental amount of distinct information that must be stored.

</details>

---

## Interview Framing Notes

This bootcamp is explicitly interview focused, so treat these points as directly applicable, not optional extras.

**Always state your complexities out loud.** Daniela was direct about this in response to a student question: you do not need to formally define what time complexity or space complexity mean in an interview unless specifically asked to. But you should, at minimum, always state the time complexity of your solution once you land on one, and ideally the space complexity as well. This signals that you can evaluate your own solution with concrete metrics, not just that it "works."

**Use the given input size to reverse engineer the expected complexity class.** If a problem tells you n can be as large as 100,000, you already know an O(n) or O(n log n) solution is expected, and anything O(n squared) or worse will fail. If a problem gives you a suspiciously small n, like 20 or 25, that is often a deliberate signal that the intended solution is exponential, commonly involving subsets or combinations. If the input size is not stated, ask for it, since your algorithm choice should depend on it.

**If asked "can this be improved," know the two honest answers.** Sometimes there genuinely is a better approach, and you are expected to try to find it. Sometimes there is not, and the correct move is to say so plainly rather than force a fake optimization. The deciding question to ask yourself: does the problem's required output or storage genuinely need that much space (for example, a full pairwise grid, or a full subset listing), or is your current approach just storing more than necessary out of habit. If the requirement itself is inherently that large, no algorithm can beat it. If you are storing something you do not actually need to keep (like all generated subsets when you only needed to print them), that is your opportunity to improve.

**Do not assume a hash table is always the answer just because its listed complexities look best across the board.** A student raised exactly this question in the session. Daniela's response: hash tables offer average case O(1) access, insert, and search, which looks unbeatable on a reference table, but they cannot solve every problem. They have no concept of order, so problems that depend on sequence, like finding the shortest path or processing elements in the order they arrived (which requires a queue, as in breadth first search) cannot be solved with a hash table alone. Also worth remembering, the O(1) for hash tables is an average case guarantee, not a guaranteed worst case the way array access is. Choose your data structure based on what operations the problem actually needs, not just which row of the reference table looks fastest.

**Recognize "in place" and "without extra space" as constant space signals.** When a problem explicitly says you must modify the input directly and cannot create a copy, that is a direct instruction that your solution needs O(1) auxiliary space. This often points you toward techniques like two pointers, which use only a small fixed number of extra variables to rearrange data within the existing structure.

---

## Terminology Glossary

- **Space complexity**: How much memory an algorithm needs, as a function of input size.
- **Time complexity**: How many operations an algorithm performs, as a function of input size (covered in the prior session, referenced here for contrast).
- **Memory limit exceeded (MLE)**: The error triggered when a program uses more memory than the system allows.
- **Input space**: Memory needed to store the given input.
- **Auxiliary space (extra space)**: Memory used beyond the input itself. This is what "space complexity of an algorithm" usually refers to in practice.
- **Allocation**: The act of a program reserving a block of memory upfront, before it is necessarily filled.
- **Dynamic array**: A growable list style structure that manages its own resizing (Python's list, C++'s vector).
- **Static array**: A fixed size array whose length must be declared in advance and cannot change.
- **Doubling strategy**: The technique dynamic arrays use to grow their capacity, doubling in size each time the current capacity is exhausted, which keeps total copying work linear overall.
- **Amortized cost**: The average cost per operation when occasional expensive operations (like a resize) are spread out across many cheap ones.
- **Constant space, O(1)**: Memory usage independent of input size.
- **Linear space, O(n)**: Memory usage directly proportional to input size.
- **Quadratic space, O(n squared)**: Memory usage proportional to input size multiplied by itself, common with two dimensional structures.
- **Cubic space, O(n cubed)**: Memory usage proportional to input size raised to the third power, common with three dimensional structures.
- **Exponential space, O(2 to the power of n)**: Memory usage that roughly doubles per additional input element, common with subset or combination generation.
- **Subset**: Any selection (including none, or all) of elements from a set.
- **In place**: A constraint requiring a solution to modify the given input directly rather than allocate a new structure, typically implying constant extra space.
- **Two pointers**: A technique using two tracked index positions to process or rearrange a structure using a fixed number of extra variables.
- **Hash table**: A structure offering average case constant time insert, search, and delete, but with no inherent concept of ordering.
- **Queue**: A structure that retrieves the earliest added element first, essential for algorithms like breadth first search.
- **O(V plus E)**: A graph specific complexity expression, where V is the vertex (node) count and E is the edge (connection) count, kept as two separate variables because their relative sizes vary independently across different graphs.

---

## Recap and What Comes Next

This session covered how to reason about memory the same rigorous way you already reason about speed. The core ideas to hold onto: space complexity equals input space plus auxiliary space, but in practice you almost always report just the auxiliary part. Memory is not filled gradually the way intuition suggests, it is allocated upfront in blocks, which is exactly why dynamic arrays use a doubling strategy internally rather than growing one unit or one fixed chunk at a time. You now have five named growth shapes to pattern match against, constant, linear, quadratic, cubic and beyond, and exponential, along with rough size thresholds for when each one starts becoming risky in practice. You also saw why input size limits given in a problem are not just flavor text, they are a direct signal for which complexity class your solution needs to land in before you ever write a line of code.

The live problem, Move Zeroes, was left unfinished on purpose. Daniela worked through a correct but non compliant two pass solution (O(n) time, O(n) space, but it violates the in place requirement) and began pointing toward the number of zeros and the two pointers technique as the path to a compliant O(n) time, O(1) space solution. The next session, described as a deep dive, will complete this problem properly and go further into how to construct two pointer based, in place solutions from scratch. You should attempt to solve Move Zeroes yourself before that session, using what you now know about constant space constraints, so the walkthrough lands as confirmation and refinement rather than an entirely new idea.

Beyond that, Daniela flagged that the space and time reference sheet covering data structures like arrays, linked lists, stacks, hash tables, and queues, along with the graph specific O(V plus E) notation, is meant to be revisited repeatedly as you learn each new data structure in future weeks, rather than memorized all at once now.
