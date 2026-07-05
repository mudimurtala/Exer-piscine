# MODULE 4 — Algorithms

To make this module comprehensive, practical, and easy to digest, we'll divide it into **4 parts**.

---

# Part 1 — Algorithm Foundations

- What is an Algorithm?
- Why Algorithms Matter
- Characteristics of a Good Algorithm
- Thinking Algorithmically
- Writing Algorithms (Plain English, Flowcharts, Pseudocode)
- Inputs, Outputs, and Steps
- Algorithm Design Basics
- Common Algorithmic Thinking Patterns
- Hands on Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 2 — Searching & Sorting Algorithms

- Linear Search
- Binary Search
- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort (Conceptual)
- Quick Sort (Conceptual)
- When to Use Each Algorithm
- Time Complexity Comparison
- Hands on Coding Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 3 — Problem Solving Patterns & Big O

- Step-by-Step Problem Solving
- Brute Force vs Optimized Solutions
- Two Pointers Technique
- Sliding Window
- Hashing for Faster Lookups
- Recursion Fundamentals
- Divide and Conquer
- Greedy Thinking
- Dynamic Programming (Beginner Introduction)
- Big O Deep Dive
- Time & Space Complexity Analysis
- Hands on Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 4 — Thinking Like an AI Builder with Algorithms

- How AI Systems Use Algorithms
- Choosing the Right Algorithm
- Reading Algorithm Interview Questions
- Common Coding Assessment Patterns
- End-to-End Problem Solving Framework
- Common Beginner Mistakes
- Comprehensive Module Assessment
- Algorithm Cheat Sheet
- Module Summary

---

# AI BUILDER ASSESSMENT BOOTCAMP — MODULE 4: Algorithms

## Part 1 — Algorithm Foundations

**Estimated Study Time:** 7–8 Hours
**Difficulty:** ⭐⭐⭐☆☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Explain what an algorithm is.
✅ Explain why algorithms are important.
✅ Distinguish between a program and an algorithm.
✅ Identify the characteristics of a good algorithm.
✅ Think algorithmically before writing code.
✅ Write algorithms using plain English, flowcharts, and pseudocode.
✅ Break large problems into smaller steps.
✅ Solve simple algorithmic reasoning questions.

---

### Chapter 1 — What is an Algorithm?

Many beginners believe that programming is about writing code. It isn't. Programming is about solving problems, and code is only one way of expressing the solution. That solution is called an **algorithm**.

**Definition:** An **algorithm** is a finite, step-by-step procedure for solving a problem or accomplishing a task. Notice the important words: step-by-step, procedure, solve a problem. An algorithm is simply a sequence of instructions.

**Real Life Analogy:** Imagine your mother asks you to make tea. She might say: boil water, put tea into a cup, pour hot water, add sugar, stir, then serve. That is an algorithm, notice there is no programming language involved, just clear instructions. Another example is directions from your house to school: leave the house, walk to the junction, turn left, continue straight, then enter the school gate. Also an algorithm.

**Builder Insight:** Algorithms existed long before computers. Humans have always used algorithms, cooking recipes, instruction manuals, traffic rules, and assembly guides are all examples. Computers simply execute algorithms much faster.

---

### Chapter 2 — Algorithm vs Program

Many beginners confuse these, but they are not the same. Think of building a house: the blueprint comes first, the actual building comes later. The blueprint is like an **algorithm** (the idea, the plan, the logic), while the finished house is like a **program** (the implementation of the algorithm in a programming language).

**Example:** The algorithm "take two numbers, add them, display the answer" can become this Python program:

```python
a = int(input())

b = int(input())

print(a + b)
```

Different languages can implement the same algorithm.

---

### Chapter 3 — Why Algorithms Matter

Imagine two students solve the same problem. Student A needs 30 seconds, Student B needs 30 minutes. Who has the better solution? Student A, and the difference is usually the algorithm.

Good algorithms make programs faster, simpler, easier to maintain, easier to debug, and more efficient. Poor algorithms make programs slow, confusing, and difficult to improve.

**Real Life Example:** Suppose you're looking for your name in a class register. Method One is reading every name (Ali, Amina, John... Mudi) one at a time. Method Two works because the names are arranged alphabetically, so you immediately jump close to "M." The second method is a better algorithm.

---

### Chapter 4 — Characteristics of a Good Algorithm

A good algorithm has several important properties.

**1. It Must Be Correct.** If the answer is wrong, the algorithm is useless. Correctness comes first.

**2. It Must Finish.** An algorithm should eventually stop. A bad algorithm simply repeats forever, while a good algorithm repeats until all numbers are processed.

**3. It Should Be Efficient.** If two algorithms solve the same problem, prefer the one that uses less time or memory.

**4. It Should Be Clear.** Someone else should understand your logic. If you cannot explain your algorithm, you probably don't understand it well enough.

**5. It Should Accept Input.** For example, a Student Score.

**6. It Should Produce Output.** For example, a Grade of A. Every useful algorithm transforms input into output.

---

### Chapter 5 — Thinking Algorithmically

Professional programmers rarely jump straight into coding, instead they think. Suppose I ask you: "Find the tallest student in a classroom." Would you immediately open VS Code? No, you would first think.

A possible thinking process: look at the first student, assume they are the tallest, compare with the next student, and if the next student is taller update the tallest, continue until everyone has been checked, then return the tallest student. Notice that no code has been written, yet the solution already exists. That is algorithmic thinking.

**Algorithm Design Process:** Whenever you solve a problem, follow this framework: understand the problem, identify the input, identify the expected output, break the problem into smaller steps, arrange the steps logically, test the steps manually, write code, then test the code. Many beginners skip directly to the last two steps, professionals don't.

---

### Chapter 6 — Inputs, Outputs, and Processing

Every algorithm has three main parts: **Input → Processing → Output**.

**Example:** Input of 5 and 7, processing that adds them, output of 12.

**Example:** Input of Student Scores, processing that calculates the average, output of the Average Score.

---

### Chapter 7 — Ways to Represent an Algorithm

Algorithms can be written in several ways.

**Method 1 — Plain English.** Problem: find the largest number. Algorithm: read the first number, assume it is the largest, compare with every remaining number, replace if a larger number is found, then display the largest number.

**Method 2 — Flowchart.** Start, input numbers, compare numbers, ask if the largest has been found, if yes display the largest, then end. Flowcharts help visualize logic.

**Method 3 — Pseudocode.** Pseudocode is fake code, it looks like programming but ignores language syntax. For example:

```
START

Read A

Read B

Sum = A + B

Display Sum

END
```

This could later become Python, Go, Java, JavaScript, or any other language.

**Why Pseudocode Is Powerful:** Imagine you're interviewing and the interviewer asks, "Don't code yet, explain your solution." Professionals often begin with pseudocode, since it shows that you understand the logic before worrying about syntax.

---

### Chapter 8 — Common Algorithmic Thinking Patterns

Many problems follow familiar patterns, and recognizing these patterns makes problem-solving easier.

- **Pattern 1 — Repeat.** Need to process every item? Think Loop, e.g. read every student's score.
- **Pattern 2 — Compare.** Need the biggest, smallest, or highest score? Think Comparison.
- **Pattern 3 — Count.** Need to know how many? Think Counter, e.g. `count = count + 1`.
- **Pattern 4 — Accumulate.** Need a total? Think Running Sum, e.g. `total = total + number`.
- **Pattern 5 — Decide.** Need different actions? Think if/else.

**Worked Example:** Problem: calculate the average of five numbers. Input is five numbers, processing adds them together and divides by five, output is the average. Pseudocode:

```
START

Read five numbers

Add them together

Divide total by 5

Display average

END
```

Only after this would you begin writing code.

---

### Hands on Exercises

1. Write an algorithm (in plain English) for brushing your teeth. Do **not** write code, focus only on the sequence of steps.
2. Write pseudocode for calculating the area of a rectangle. Hint: Input is Length and Width, Output is Area.
3. A teacher wants to calculate the total score of 30 students. Identify the Input, Processing, and Output.
4. Write a simple algorithm for finding the smallest number in a list of ten numbers. Do not write code, think only about the logical steps.

---

### Logical Reasoning Challenges

**Challenge 1:** A recipe says: bake for 20 minutes, and if the cake is not ready, continue baking for 5 more minutes, repeating until ready. Is this an algorithm?
✅ **Yes.** It is a sequence of logical instructions that eventually produces a result.

**Challenge 2:** A robot receives these instructions: walk forward, turn right, walk forward, turn right, walk forward, turn right, walk forward, turn right. What shape will the robot trace?
✅ **A square.**

**Challenge 3:** You are given two algorithms that both produce the correct answer. Algorithm A needs 10 steps, Algorithm B needs 1,000 steps. Which is generally preferable?
✅ **Algorithm A**, because it is more efficient.

---

### AI Builder Assessment Questions

1. What is an algorithm?
   A. A programming language  B. A computer  C. A step-by-step procedure for solving a problem  D. A database
   ✅ **C**

2. Which comes first?
   A. Program  B. Algorithm
   ✅ **B**

3. Which is **not** a characteristic of a good algorithm?
   A. Correctness  B. Efficiency  C. Infinite repetition  D. Produces output
   ✅ **C**

4. Which representation ignores programming language syntax?
   A. Python  B. Go  C. Pseudocode  D. JavaScript
   ✅ **C**

5. An algorithm always transforms:
   A. Input into Output  B. Output into Input  C. Variables into Functions  D. Functions into Loops
   ✅ **A**

6. What should you do before writing code?
   A. Choose random variable names  B. Understand the problem  C. Open your IDE immediately  D. Search Stack Overflow
   ✅ **B**

---

### Common Beginner Mistakes

❌ Starting to code before understanding the problem.
❌ Confusing an algorithm with a programming language.
❌ Ignoring the input and expected output.
❌ Thinking pseudocode must follow strict syntax.
❌ Trying to memorize algorithms instead of understanding the logic behind them.

### Builder Tips

As an AI Builder, you'll often use AI tools like ChatGPT or GitHub Copilot to generate code. Before accepting the code, ask yourself: What problem is this algorithm solving? What is the input? What is the output? Can I explain the steps without looking at the code? Is there a simpler approach? If you can answer these questions, you understand the algorithm, not just the code.

### Builder Checkpoint

Before moving to Part 2, make sure you can confidently say:

✅ I understand what an algorithm is.
✅ I know the difference between an algorithm and a program.
✅ I understand why algorithms matter.
✅ I know the characteristics of a good algorithm.
✅ I can break a problem into logical steps.
✅ I can identify inputs, processing, and outputs.
✅ I can write simple pseudocode.
✅ I think about the solution before writing code.

### Fast Review

- An **algorithm** is a step-by-step method for solving a problem.
- A **program** is an implementation of an algorithm in a programming language.
- Good algorithms are **correct**, **finite**, **efficient**, **clear**, and produce meaningful output.
- Every algorithm has **Input → Processing → Output**.
- Algorithms can be represented using **plain English**, **flowcharts**, or **pseudocode**.
- Good programmers don't start by coding, they start by understanding the problem and designing the solution.
- Strong algorithmic thinking is one of the most important skills for technical interviews, AI Builder assessments, and real world software development.

---

## Part 2 — Searching & Sorting Algorithms

**Estimated Study Time:** 8–10 Hours
**Difficulty:** ⭐⭐⭐⭐☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Explain what searching algorithms are.
✅ Explain what sorting algorithms are.
✅ Understand how Linear Search works.
✅ Understand how Binary Search works.
✅ Understand Bubble Sort.
✅ Understand Selection Sort.
✅ Understand Insertion Sort.
✅ Understand the ideas behind Merge Sort and Quick Sort.
✅ Compare searching and sorting algorithms.
✅ Analyze their time complexity at a beginner-friendly level.

---

### Chapter 1 — Why Searching and Sorting Matter

Imagine you're working at Google, and a user searches for "AI Builder." Google has billions of webpages, so how can it find the correct ones almost instantly? Algorithms. Imagine Amazon, where a customer wants products sorted from cheapest to most expensive, again, algorithms. Searching and sorting are among the most common operations in programming, almost every software system performs them.

---

### Chapter 2 — Searching Algorithms

Searching answers one question: **"Where is the item I'm looking for?"** There are many searching algorithms, but the two most important for beginners are Linear Search and Binary Search.

---

### Chapter 3 — Linear Search

Imagine you're looking for your name in an attendance list (Ali, John, Fatima, Grace, Mudi). You start from the top, check Ali (not found), John (not found), Fatima (not found), Grace (not found), then Mudi (found). This is Linear Search.

**Definition:** Linear Search checks items **one by one** until it finds the target.

**Example:** Find 7 inside `3, 5, 1, 7, 8`. Check 3, not found. Check 5, not found. Check 1, not found. Check 7, found.

```python
numbers = [3, 5, 1, 7, 8]

target = 7

for number in numbers:
    if number == target:
        print("Found!")
        break
```

**Time Complexity:** Worst case `O(n)`, because every element might need to be checked.

**Advantages:** Very simple, works on any list, no sorting required.
**Disadvantages:** Slow for large datasets.

---

### Chapter 4 — Binary Search

Imagine a dictionary. Do you start reading from page one? No, you jump close to the middle, then eliminate half, then eliminate another half. This is Binary Search.

**Requirement:** The data **must already be sorted**. Without sorting, Binary Search cannot work correctly.

**Example:** Find 17 inside `2, 5, 8, 11, 17, 22, 30`. The middle is 11, and since 17 is greater, ignore the left half and search only `17, 22, 30`. The new middle is 22, too large, so search left, finding 17.

**Why Binary Search Is Fast:** Every comparison removes half the remaining data. For 1,000,000 items, Linear Search may check all 1,000,000, while Binary Search needs only about 20 comparisons, a huge difference.

```python
numbers = [2,5,8,11,17,22,30]

left = 0

right = len(numbers)-1

target = 17

while left <= right:

    middle = (left + right) // 2

    if numbers[middle] == target:
        print("Found")
        break

    elif numbers[middle] < target:
        left = middle + 1

    else:
        right = middle - 1
```

**Time Complexity:** `O(log n)`.

**Linear Search vs Binary Search**

| Linear Search | Binary Search |
|---------------|---------------|
| Works on any list | Requires sorted data |
| Checks one by one | Eliminates half each step |
| O(n) | O(log n) |
| Easy to implement | Faster on large sorted data |

---

### Chapter 5 — Why Sorting Matters

Suppose your teacher wants scores arranged from highest to lowest: `78, 25, 99, 61, 42` becomes `25, 42, 61, 78, 99`. Sorting makes searching, ranking, and reporting easier.

---

### Chapter 6 — Bubble Sort

Bubble Sort repeatedly compares neighboring items, and if they are in the wrong order, it swaps them. For example, `5, 2, 4` compares 5 and 2, swaps to get `2, 5, 4`, then compares 5 and 4, swaps to get `2, 4, 5`, which is sorted.

**Visual Example:** Pass 1 on `5 2 4 1` produces `2 5 4 1`, then `2 4 5 1`, then `2 4 1 5`. Pass 2 produces `2 1 4 5`, then `1 2 4 5`. Done.

```python
numbers = [5,2,4,1]

for i in range(len(numbers)):

    for j in range(len(numbers)-1):

        if numbers[j] > numbers[j+1]:

            numbers[j], numbers[j+1] = numbers[j+1], numbers[j]

print(numbers)
```

**Time Complexity:** Worst case `O(n²)`.
**Advantages:** Easy to understand.
**Disadvantages:** Very slow for large datasets.

---

### Chapter 7 — Selection Sort

Selection Sort repeatedly finds the smallest value and places it at the beginning. For example, `9 5 2 7` has a smallest value of 2, which moves forward to give `2 5 9 7`, and repeating this process gives the final result `2 5 7 9`.

**Time Complexity:** `O(n²)`.

---

### Chapter 8 — Insertion Sort

Imagine arranging playing cards, where each new card is inserted into the correct position. For example, starting with `7`, then `7 4` becomes `4 7`, then `4 7 9`, then `4 5 7 9`.

**Time Complexity:** Worst case `O(n²)`, best case `O(n)` when already sorted.

---

### Chapter 9 — Merge Sort (Conceptual)

Instead of sorting everything at once, Merge Sort divides the problem. For example, `8 3 6 2` splits into `8 3` and `6 2`, then splits again into `8, 3, 6, 2`. Each small part is sorted, then merged into `3 8` and `2 6`, and merged again into `2 3 6 8`.

**Big Idea:** Split, solve, combine, this is called **Divide and Conquer**.

**Time Complexity:** `O(n log n)`.

---

### Chapter 10 — Quick Sort (Conceptual)

Quick Sort chooses one element, called the **Pivot**. Everything smaller goes left, everything larger goes right, and the process repeats. For example, with `7 2 9 4`, choosing 7 as pivot gives a left side of `2 4` and a right side of `9`. Sorting each side gives the final result `2 4 7 9`.

**Time Complexity:** Average `O(n log n)`, worst `O(n²)`.

---

### Chapter 11 — Comparing Sorting Algorithms

| Algorithm | Best Case | Average Case | Worst Case | Stable | Beginner Friendly |
|------------|-----------|--------------|------------|--------|-------------------|
| Bubble Sort | O(n) | O(n²) | O(n²) | Yes | ⭐⭐⭐⭐⭐ |
| Selection Sort | O(n²) | O(n²) | O(n²) | No | ⭐⭐⭐⭐☆ |
| Insertion Sort | O(n) | O(n²) | O(n²) | Yes | ⭐⭐⭐⭐☆ |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | Yes | ⭐⭐⭐☆☆ |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | No | ⭐⭐⭐☆☆ |

**Which Algorithm Should You Use?**

| Situation | Best Choice |
|------------|-------------|
| Small list | Bubble / Insertion |
| Nearly sorted data | Insertion |
| Large datasets | Merge Sort |
| Fast average performance | Quick Sort |
| Unsorted small dataset search | Linear Search |
| Sorted dataset search | Binary Search |

---

### Hands on Exercises

1. Perform a Linear Search for `18` inside `7 10 12 18 20 31`, writing every comparison.
2. Perform a Binary Search for `45` inside `5 10 20 30 45 60 75`, drawing every step.
3. Sort `9 4 2 8 1` using Bubble Sort, showing every swap.
4. Sort `8 6 5 2` using Selection Sort.
5. Arrange `6 4 7 2` using Insertion Sort.

---

### Logical Reasoning Challenges

**Challenge 1:** A phone contact list is already sorted alphabetically. Which search algorithm is usually the better choice?
A. Linear Search  B. Binary Search
✅ **Binary Search**

**Challenge 2:** You receive random unsorted data. Can Binary Search be used immediately?
✅ **No.** The data must first be sorted.

**Challenge 3:** Which sorting algorithm repeatedly swaps neighboring elements?
A. Selection Sort  B. Bubble Sort  C. Merge Sort  D. Quick Sort
✅ **Bubble Sort**

---

### AI Builder Assessment Questions

1. Which search algorithm works on unsorted data?
   A. Binary Search  B. Linear Search  C. Merge Search  D. Quick Search
   ✅ **B**

2. Binary Search requires:
   A. A Stack  B. Sorted Data  C. A Queue  D. A Tree
   ✅ **B**

3. Which algorithm repeatedly swaps neighboring elements?
   A. Bubble Sort  B. Selection Sort  C. Insertion Sort  D. Merge Sort
   ✅ **A**

4. Which sorting algorithm repeatedly selects the smallest remaining element?
   A. Selection Sort  B. Bubble Sort  C. Quick Sort  D. Binary Sort
   ✅ **A**

5. Which search algorithm has an average time complexity of **O(log n)**?
   A. Linear Search  B. Binary Search  C. Bubble Sort  D. Selection Sort
   ✅ **B**

6. Which algorithm follows the Divide and Conquer strategy?
   A. Merge Sort  B. Bubble Sort  C. Selection Sort  D. Linear Search
   ✅ **A**

---

### Common Beginner Mistakes

❌ Trying to use Binary Search on unsorted data.
❌ Memorizing algorithms without understanding how they work.
❌ Believing Bubble Sort is used in production for large datasets.
❌ Confusing searching algorithms with sorting algorithms.
❌ Forgetting that faster algorithms often require additional ideas or constraints (such as sorted input).

### Builder Tips

As an AI Builder, you may rarely implement Merge Sort or Quick Sort from scratch, but you will constantly make decisions such as: Should I search linearly or use indexed data? Should I sort before searching? Is my dataset already sorted? Do I need the fastest algorithm or the simplest one? Understanding **why** an algorithm is chosen is more valuable than memorizing its code.

### Builder Checkpoint

Before moving to Part 3, make sure you can confidently say:

✅ I understand the difference between searching and sorting.
✅ I know how Linear Search works.
✅ I know how Binary Search works and why it requires sorted data.
✅ I understand Bubble, Selection, and Insertion Sort.
✅ I understand the basic ideas behind Merge Sort and Quick Sort.
✅ I can compare common searching and sorting algorithms.

### Fast Review

- **Linear Search** checks items one by one (**O(n)**).
- **Binary Search** repeatedly halves a **sorted** search space (**O(log n)**).
- **Bubble Sort** swaps adjacent elements until the list is sorted.
- **Selection Sort** repeatedly places the smallest remaining element in its correct position.
- **Insertion Sort** inserts each new element into its proper place within the sorted portion.
- **Merge Sort** uses **Divide and Conquer** and runs in **O(n log n)**.
- **Quick Sort** partitions around a pivot and is very fast on average (**O(n log n)**).
- Choosing the right algorithm depends on the size of the data, whether it is already sorted, and the requirements of the problem.

---

## Part 3 — Problem Solving Patterns & Big O

**Estimated Study Time:** 10–12 Hours
**Difficulty:** ⭐⭐⭐⭐☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Approach programming problems like an engineer.
✅ Break complex problems into smaller ones.
✅ Understand Brute Force vs Optimized solutions.
✅ Understand the Two Pointers technique.
✅ Understand the Sliding Window technique.
✅ Use Hashing to speed up lookups.
✅ Understand the basics of Recursion.
✅ Understand Divide and Conquer.
✅ Understand Greedy Algorithms.
✅ Understand the idea behind Dynamic Programming.
✅ Analyze Time and Space Complexity with confidence.

---

### Chapter 1 — The AI Builder Problem Solving Framework

Experienced developers don't immediately begin coding, instead they ask themselves a series of questions. Whenever you see a programming problem, train yourself to think like this: understand the problem, identify the input, identify the expected output, identify the constraints, ask whether you can solve it manually, describe the steps, look for ways to improve those steps, then write and test the code. Never skip the thinking stage.

**Example:** For the problem "find the largest number in a list," most beginners immediately write code, but professionals first think: look at the first number, assume it is the largest, compare with the next number, replace if necessary, repeat until the end, then return the largest number. Only after understanding this logic do they begin coding.

---

### Chapter 2 — Brute Force vs Optimized Solutions

There are usually many ways to solve a problem, some slow, some fast.

**Brute Force** simply tries everything. For example, finding a name in an unsorted list (Ali, John, Grace, Fatima, Mudi) means checking every name, which is simple and correct but not always efficient. Advantages: easy to write and understand. Disadvantages: slow on large datasets.

**Optimized Solutions** look for a smarter strategy. If the names are sorted (Ali, Fatima, Grace, John, Mudi), you can use Binary Search instead, which is much faster.

**Builder Tip:** A brute-force solution is not "bad." Many interviews prefer a correct brute-force solution before an optimized one, correctness first, optimization second.

---

### Chapter 3 — Two Pointers Technique

Imagine two fingers pointing at opposite ends of a list:

```
1   2   3   4   5

↑           ↑
Left      Right
```

Instead of using one pointer, you use two. This is useful for sorted arrays, finding pairs, removing duplicates, and comparing elements.

**Example:** To find whether two numbers add up to 10 in `1 2 3 4 6 8 9`, set Left = 1 and Right = 9, giving `1 + 9 = 10`, found immediately without checking every possible pair.

```python
numbers = [1,2,3,4,6,8,9]

left = 0
right = len(numbers) - 1

target = 10

while left < right:

    total = numbers[left] + numbers[right]

    if total == target:
        print("Found")
        break

    elif total < target:
        left += 1

    else:
        right -= 1
```

---

### Chapter 4 — Sliding Window

Imagine looking through a moving window over `1 2 3 4 5 6` with a window size of 3: `[1 2 3]`, then `[2 3 4]`, then `[3 4 5]`, then `[4 5 6]`. Instead of recalculating everything, the window slides. This is used for finding maximum or minimum sums, longest substrings, and consecutive values.

**Example:** Find the largest sum of three consecutive numbers in `2 4 1 6 8 3`. The window sums are `2+4+1=7`, `4+1+6=11`, `1+6+8=15`, `6+8+3=17`, so the largest is 17. Without Sliding Window, you would repeat many calculations.

---

### Chapter 5 — Hashing

Hashing usually means using a Dictionary (Hash Table). Instead of searching repeatedly, you store information for quick lookup.

**Example:** Count word frequencies in "AI is amazing and AI is useful" to get the dictionary `AI → 2, is → 2, amazing → 1, and → 1, useful → 1`.

```python
words = ["AI","is","AI"]

count = {}

for word in words:

    count[word] = count.get(word,0)+1

print(count)
```

Output: `{'AI':2,'is':1}`. Hashing is extremely common in AI, backend systems, and coding interviews.

---

### Chapter 6 — Recursion

Recursion is one of the most misunderstood topics, but we'll build it slowly.

**What Is Recursion?** A function calls itself. This sounds strange, but many real-world problems naturally repeat themselves, like walking upstairs one similar step at a time. A recursive function solves a small version of the same problem.

**Example:** A countdown function:

```python
def countdown(n):

    if n == 0:
        print("Done")
        return

    print(n)

    countdown(n-1)

countdown(5)
```

Output: `5, 4, 3, 2, 1, Done`.

**Base Case:** Every recursive function **must** know when to stop. Without a stopping condition, it will continue forever until the program crashes, e.g. `if n == 0: return`. This is called the **Base Case**.

**Builder Tip:** Many beginners fear recursion because they imagine magic. Think of it as a function repeatedly asking a smaller version of the same question until there is nothing left to solve.

---

### Chapter 7 — Divide and Conquer

Divide a large problem, solve each small part, then combine the answers. For example, 16 numbers divide into 8, then 4, then 2, then 1. This idea powers Merge Sort, Binary Search, and Quick Sort.

---

### Chapter 8 — Greedy Algorithms

A Greedy Algorithm always chooses what seems best **right now**, never looking far into the future. For example, giving ₦70 in change using available notes of ₦50, ₦20, and ₦10, Greedy chooses ₦50 then ₦20. Greedy is simple, sometimes it gives the optimal solution, sometimes it doesn't.

---

### Chapter 9 — Dynamic Programming (Beginner Introduction)

Dynamic Programming (DP) sounds scary, but the main idea is simple. Imagine climbing stairs, to reach step 5 you've already calculated the best ways to reach step 3 and step 4. Why calculate them again? Store the answers and reuse them, that's Dynamic Programming.

**Core Idea:** Don't solve the same problem repeatedly, remember previous answers. For example, instead of recalculating "What is 15 × 15?" every time, you memorize the answer once. DP works similarly.

**Builder Tip:** For AI Builder assessments, you are unlikely to implement advanced DP from scratch, but understanding the idea of **saving previous work** is important.

---

### Chapter 10 — Big O Deep Dive

You've already seen Big O in Module 3, now let's understand it better.

**What Big O Measures:** Big O measures **how the amount of work grows as the input grows**, it is about **growth**, not seconds. A good algorithm should still perform reasonably well whether you have 10, 100, or 1,000,000 items.

- **O(1) — Constant Time:** Accessing `numbers[4]` takes roughly the same amount of work whether the list has 10 items or 10 million.
- **O(log n) — Logarithmic Time:** Binary Search, where each comparison cuts the remaining work in half (1000, 500, 250, 125...).
- **O(n) — Linear Time:** Linear Search, where you may inspect every element.
- **O(n log n):** Many efficient sorting algorithms belong here, including Merge Sort, Quick Sort (average), and Heap Sort.
- **O(n²):** Comparing every element with every other element, as in Bubble Sort and Selection Sort.

**Comparing Growth (n = 1,000)**

| Big O | Approximate Work |
|--------|------------------|
| O(1) | 1 |
| O(log n) | 10 |
| O(n) | 1,000 |
| O(n log n) | 10,000 |
| O(n²) | 1,000,000 |

Notice how quickly O(n²) grows, this is why algorithm choice matters.

---

### Chapter 11 — Time Complexity vs Space Complexity

Time Complexity asks "How much work does the algorithm perform?" Space Complexity asks "How much extra memory does the algorithm need?" For example, `numbers = [1,2,3]` uses memory for three numbers, while `numbers = [1,2,3,4,5,6]` uses more memory. Sometimes an algorithm is faster because it uses more memory, sometimes it saves memory but becomes slower, engineering is often about balancing these trade-offs.

---

### Hands on Exercises

1. Given `1 3 5 7 9 11`, use the Two Pointers technique to determine whether two numbers add up to **12**, writing each step.
2. Find the maximum sum of three consecutive numbers in `5 2 8 1 9 4` using the Sliding Window approach.
3. Write a dictionary that counts how many times each fruit appears in: Apple, Orange, Apple, Banana, Orange, Apple.
4. Trace the output of `countdown(3)` using the recursive countdown function from Chapter 6.
5. Classify each of these by its Big O category: Binary Search, Linear Search, Bubble Sort, Merge Sort, Quick Sort (average).

---

### Logical Reasoning Challenges

**Challenge 1:** You need to repeatedly look up student records by Student ID. Which approach is generally better?
A. Search the list every time  B. Store the records in a Dictionary
✅ **B.** A Dictionary provides much faster lookups.

**Challenge 2:** A list is already sorted and you need to search it many times. Which algorithm is usually the better choice?
A. Linear Search  B. Binary Search
✅ **B**

**Challenge 3:** Which programming technique avoids solving the same subproblem repeatedly?
A. Greedy  B. Dynamic Programming  C. Bubble Sort  D. Queue
✅ **B**

---

### AI Builder Assessment Questions

1. The Two Pointers technique is most useful when:
   A. Working with sorted arrays  B. Printing text  C. Creating variables  D. Connecting to the internet
   ✅ **A**

2. Sliding Window is commonly used for:
   A. Maximum or minimum values in consecutive elements  B. Drawing windows on the screen  C. Connecting databases  D. Sorting dictionaries
   ✅ **A**

3. What is the most important part of every recursive function?
   A. Loop  B. Base Case  C. Variable  D. Array
   ✅ **B**

4. Which idea is the foundation of Dynamic Programming?
   A. Repeat work many times  B. Save previous results and reuse them  C. Always choose the largest number  D. Never use memory
   ✅ **B**

5. Which Big O notation grows the slowest as the input size increases?
   A. O(n²)  B. O(n)  C. O(log n)  D. O(1)
   ✅ **D**

6. Merge Sort mainly uses which problem-solving strategy?
   A. Greedy  B. Divide and Conquer  C. Sliding Window  D. Two Pointers
   ✅ **B**

---

### Common Beginner Mistakes

❌ Trying to optimize before finding a correct solution.
❌ Using Binary Search on unsorted data.
❌ Forgetting the base case in recursion.
❌ Memorizing Big O values without understanding why they occur.
❌ Believing every problem has only one correct algorithm.
❌ Assuming the fastest algorithm always uses the least memory.

### Builder Tips

When you're building AI-powered applications, you'll often use these patterns indirectly: **Hashing** for caching API responses or counting data, **Sliding Window** for analyzing sequences of tokens or sensor readings, **Two Pointers** for efficient processing of sorted data, **Divide and Conquer** inside many optimized libraries, **Recursion** when traversing folders, trees, or nested JSON, and **Dynamic Programming** in optimization problems and parts of machine learning. The key isn't memorizing every algorithm, it's recognizing the pattern a problem belongs to.

### Builder Checkpoint

Before moving to Part 4, make sure you can confidently say:

✅ I know how to approach a new programming problem.
✅ I understand the difference between brute-force and optimized solutions.
✅ I know when to use the Two Pointers technique.
✅ I understand the Sliding Window pattern.
✅ I know how hashing speeds up lookups.
✅ I understand recursion and the importance of a base case.
✅ I understand Divide and Conquer, Greedy thinking, and the basic idea of Dynamic Programming.
✅ I can explain Big O intuitively.

### Fast Review

- Always understand the problem before writing code.
- Start with a correct solution, then optimize it.
- **Two Pointers** use two indices to solve problems efficiently, especially on sorted data.
- **Sliding Window** avoids repeating work when analyzing consecutive elements.
- **Hashing** enables fast lookups and counting.
- **Recursion** solves a problem by solving smaller versions of itself.
- **Divide and Conquer** splits a problem, solves each part, and combines the results.
- **Greedy Algorithms** make the best immediate choice, but aren't always globally optimal.
- **Dynamic Programming** stores previous results to avoid repeated work.
- **Big O** describes how an algorithm's time or memory usage grows as input size increases.

---

## Part 4 — Thinking Like an AI Builder with Algorithms

**Estimated Study Time:** 8–10 Hours
**Difficulty:** ⭐⭐⭐⭐☆ (Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Think like a software engineer instead of just a programmer.
✅ Recognize common algorithmic interview questions.
✅ Choose appropriate algorithms for different problems.
✅ Explain your reasoning before writing code.
✅ Understand how AI systems rely on algorithms.
✅ Develop confidence for AI Builder assessments and coding interviews.

---

### Chapter 1 — What Does It Mean to Think Like an AI Builder?

Many beginners think programming is about writing code. Professional developers know that programming is mostly about solving problems, writing code is only the final step. The real work happens before the keyboard is touched. An AI Builder asks questions like: What exactly is the problem? What information do I already have? What information do I need? Is there a simpler way? Can this be solved more efficiently? Which data structure fits this problem? Which algorithm fits this data structure?

**The AI Builder Workflow:** Understand the problem, identify inputs, identify outputs, understand constraints, choose a data structure, choose an algorithm, write pseudocode, write code, test, optimize, then deploy. This workflow is followed by experienced developers whether they are building a chatbot, a banking app, or an AI system.

---

### Chapter 2 — AI Is Built on Algorithms

Many people think AI is magic, it isn't. AI is software, and software runs algorithms, machine learning algorithms, search algorithms, optimization algorithms, graph algorithms, sorting algorithms, and recommendation algorithms. Every AI product uses algorithms somewhere.

**Example 1 — ChatGPT.** When you ask it to explain recursion, the system must receive your input, break it into tokens, search patterns learned during training, predict the next word, and repeat until the answer is complete. Each of these involves algorithms.

**Example 2 — Google Maps.** When you ask it to take you to the airport, the system builds a graph of roads, finds possible routes, estimates travel time, and chooses the best route. Again, algorithms.

**Example 3 — Netflix.** When you watch science documentaries, Netflix records your activity, compares your interests with other users, and recommends similar content. Recommendation algorithms make this possible.

---

### Chapter 3 — How Professionals Solve Problems

Suppose someone asks you to find duplicate usernames. A beginner might immediately start writing loops, but a professional pauses first and asks: What is the input? How large is the dataset? Do duplicates matter? Do I need speed? Can a Set solve this? Then they code.

**Example:** For the problem "remove duplicate emails," Solution 1 uses nested loops at `O(n²)`, while Solution 2 uses a Set at approximately `O(n)`. Both work, but one scales much better.

---

### Chapter 4 — Reading Hidden Interview Questions

Interviewers rarely ask "Use a Queue," instead they describe a situation. "Design an Undo feature" hints at a Stack. "Process customers in arrival order" hints at a Queue. "Store products by Product ID" hints at a Dictionary. "Suggest search results while typing" hints at a Trie. "Find the shortest route between cities" hints at a Graph with Graph Search.

**Builder Rule:** Don't search for the algorithm, search for the **pattern**.

---

### Chapter 5 — A Practical Problem Solving Framework

Whenever you see a new coding problem, ask these questions in order: What is the problem asking (rewrite it in your own words)? What is the input, e.g. a list of numbers? What is the output, e.g. the largest number? Can I solve it manually, since if you cannot solve it manually you probably cannot code it yet? Write the steps in plain English. Think about efficiency, can anything be improved? Write pseudocode. Write the code. Test with different inputs, including empty input, one item, duplicate values, very large input, and unexpected values when appropriate.

---

### Chapter 6 — Thinking Before Coding

Consider the problem "find the largest number." Many beginners immediately type `max(numbers)`, and that works, but suppose you're in an interview and the interviewer wants to know whether you understand the logic. Instead, explain: assume the first number is the largest, compare with every remaining number, replace when a larger number appears, then return the largest. Now you've demonstrated algorithmic thinking.

---

### Chapter 7 — Explaining Your Thinking

One of the biggest differences between junior and senior developers is communication. Suppose an interviewer asks why you chose a Dictionary. A weak answer is "Because ChatGPT used it." A strong answer is "I chose a Dictionary because I need fast lookups by key, each user has a unique ID, and a Dictionary allows efficient retrieval without scanning the entire collection." Notice the difference, professionals justify decisions.

---

### Chapter 8 — Common Coding Assessment Patterns

Many coding questions are variations of the same ideas. Learn the patterns instead of memorizing hundreds of problems.

| Pattern | Typical Technique |
|----------|-------------------|
| Find something | Search |
| Arrange data | Sort |
| Remove duplicates | Set |
| Count occurrences | Dictionary |
| Process in order | Queue |
| Undo operations | Stack |
| Navigate hierarchies | Tree |
| Navigate networks | Graph |
| Find pairs | Two Pointers |
| Consecutive elements | Sliding Window |

---

### Chapter 9 — Common AI Builder Mistakes

**Mistake 1:** Starting to code immediately, professionals think first.

**Mistake 2:** Ignoring edge cases such as an empty list, a single element, negative numbers, or duplicate values. Always test unusual inputs.

**Mistake 3:** Using the wrong data structure. Choosing a List when a Dictionary is needed can make a solution much slower.

**Mistake 4:** Optimizing too early. Write a correct solution first, improve it later.

**Mistake 5:** Depending entirely on AI. AI tools are assistants, not replacements for understanding. If ChatGPT generates code, you should still be able to explain what it does, why it works, and its limitations.

---

### Chapter 10 — AI Builder Assessment Strategy

When you receive a question, don't panic. Follow this checklist: read carefully, underline important keywords, identify the input, identify the output, recognize the pattern, choose the data structure, choose the algorithm, estimate the time complexity, explain your reasoning, write the code, then test. This simple routine helps reduce mistakes under pressure.

---

### Chapter 11 — Mini Case Studies

**Case Study 1:** Find duplicate usernames. Think: duplicates → Set.

**Case Study 2:** Store customer information using Customer ID. Think: key → value → Dictionary.

**Case Study 3:** Find the fastest route between cities. Think: road network → Graph → shortest-path algorithm.

**Case Study 4:** Process support tickets in arrival order. Think: arrival order → Queue.

**Case Study 5:** Implement an Undo button in a drawing app. Think: most recent action → Stack.

---

### Hands on Exercises

1. For each problem below, identify the most suitable data structure, a suitable algorithm or technique, and why you chose it: Remove duplicate emails, Search sorted student IDs, Auto-complete search, GPS navigation, Count word frequency, Undo operation, Browser history.
2. Given the problem "return the second largest number from a list of numbers," answer: What is the input? What is the output? How would you solve it manually? Which data structure will you use? Can you think of more than one solution?
3. A company wants to build a simple chatbot. List at least five algorithms or data structures that might be useful somewhere in the system, and explain briefly why.

---

### Logical Reasoning Challenges

**Challenge 1:** A customer service system always answers the oldest waiting request first. Which data structure best models this?
A. Stack  B. Queue  C. Graph  D. Set
✅ **Queue**

**Challenge 2:** An application must instantly retrieve a user's profile from their unique ID. Which data structure is most appropriate?
A. List  B. Queue  C. Dictionary  D. Stack
✅ **Dictionary**

**Challenge 3:** A recommendation engine connects users with similar interests. Which data structure best represents those relationships?
A. Graph  B. Array  C. Stack  D. Queue
✅ **Graph**

---

### Final AI Builder Mock Assessment

1. Before writing code, what should you do first?
   A. Open the IDE  B. Memorize syntax  C. Understand the problem  D. Search online
   ✅ **C**

2. Which technique is commonly used to count occurrences efficiently?
   A. Queue  B. Dictionary (Hash Table)  C. Bubble Sort  D. Stack
   ✅ **B**

3. Which statement about AI coding assistants is most accurate?
   A. They replace the need to understand code  B. They always produce the best solution  C. They are tools that assist developers, but developers should still understand the generated code  D. They eliminate the need for algorithms
   ✅ **C**

4. A problem mentions "find the shortest route." Which type of data structure should immediately come to mind?
   A. Queue  B. Graph  C. Stack  D. Set
   ✅ **B**

5. Why is it important to test edge cases?
   A. To make the code longer  B. To ensure the algorithm behaves correctly under unusual inputs  C. To improve internet speed  D. To reduce variable names
   ✅ **B**

6. Which habit best reflects an AI Builder's mindset?
   A. Writing code immediately  B. Copying code without understanding it  C. Understanding the problem, choosing appropriate tools, then implementing and testing  D. Memorizing every algorithm
   ✅ **C**

---

### AI Builder Interview Tips

If you're asked to solve a coding problem: read the problem carefully, restate it in your own words, ask clarifying questions if something is ambiguous, explain your planned approach before coding, mention the data structure and algorithm you're choosing, estimate the time and space complexity, write clean and readable code, test your solution with simple and edge-case inputs, and if time permits, discuss possible improvements. Interviewers often evaluate **how you think**, not just whether your final code is correct.

### Builder Checklist

Before moving to Module 5, make sure you can confidently say:

✅ I know how to approach unfamiliar programming problems.
✅ I can identify common algorithmic patterns.
✅ I choose data structures based on the problem, not habit.
✅ I can explain my reasoning before writing code.
✅ I understand the relationship between algorithms, data structures, and AI systems.
✅ I know how to analyze a basic solution using time and space complexity.
✅ I understand how to work effectively with AI coding assistants without relying on them blindly.

---

### Module 4 Cheat Sheet

| Situation | Recommended Tool |
|-----------|------------------|
| Search unsorted data | Linear Search |
| Search sorted data | Binary Search |
| Sort small/simple data | Insertion Sort / Bubble Sort |
| Sort large datasets | Merge Sort / Quick Sort |
| Remove duplicates | Set |
| Fast lookup | Dictionary |
| Count frequencies | Dictionary |
| Process in arrival order | Queue |
| Undo last action | Stack |
| Find shortest route | Graph + Graph Search |
| Auto-complete | Trie |
| Find pairs in sorted data | Two Pointers |
| Consecutive sequence problems | Sliding Window |

---

### Module 4 Summary

Congratulations! You have completed **Module 4 — Algorithms**. You now understand:

- What algorithms are and why they matter.
- How to think algorithmically before writing code.
- The most common searching and sorting algorithms.
- Core problem-solving patterns used in technical interviews.
- The intuition behind Big O, time complexity, and space complexity.
- How experienced developers approach coding problems.
- How algorithms and data structures work together to power AI applications.

> **Remember this principle:** A great programmer doesn't start by writing code, they start by understanding the problem. Once the problem is clear, the right data structure and algorithm usually become much easier to choose.