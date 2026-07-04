# MODULE 3 — Data Structures

To keep this module comprehensive while remaining easy to study, we'll divide it into **4 parts**.

---

# Part 1 — Linear Data Structures (Foundations)

- What is a Data Structure?
- Why Data Structures Matter
- Choosing the Right Data Structure
- Arrays
- Lists (Dynamic Arrays)
- Linked Lists
- Strings as Data Structures
- Time Complexity of Basic Operations
- Real world Analogies
- Hands on Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 2 — Specialized Data Structures

- Stacks
- Queues
- Deques
- Priority Queues
- Hash Tables (Maps/Dictionaries)
- Sets
- Real world Applications
- Choosing Between Them
- Time Complexity Comparison
- Hands on Coding Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 3 — Non Linear Data Structures

- Trees
- Binary Trees
- Binary Search Trees (BST)
- Heaps
- Tries (Prefix Trees)
- Graphs
- Graph Traversal (BFS & DFS)
- Real world Applications
- Time Complexity
- Hands on Exercises
- Logical Reasoning Questions
- AI Builder Assessment Questions
- Builder Checkpoint
- Fast Review

---

# Part 4 — Thinking Like an AI Builder with Data Structures

- How Interview Questions Test Data Structures
- How to Choose the Right Data Structure
- Reading Data Structure Problems
- Common Problem Solving Patterns
- Beginner Friendly Introduction to Big O
- Time Complexity Cheat Sheet
- Memory (Space Complexity) Basics
- Common Mistakes
- Comprehensive Module Assessment
- Data Structures Cheat Sheet
- Module Summary

---

# AI BUILDER ASSESSMENT BOOTCAMP — MODULE 3: Data Structures

## Part 1 — Linear Data Structures (Foundations)

**Estimated Study Time:** 6–7 Hours
**Difficulty:** ⭐⭐⭐☆☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Explain what a data structure is.
✅ Explain why data structures are important.
✅ Understand the difference between static and dynamic collections.
✅ Understand Arrays.
✅ Understand Lists (Dynamic Arrays).
✅ Understand Linked Lists.
✅ Understand Strings as a data structure.
✅ Understand the strengths and weaknesses of each.
✅ Choose the right data structure for simple problems.

---

### Chapter 1 — What is a Data Structure?

Before we define it, let's think about something. Imagine you walk into a library. Thousands of books are inside. Now imagine the books are simply thrown on the floor, with no shelves, no labels, no categories. Could you find one particular book? Probably not.

Now imagine another library, where books are arranged by subject, author, shelf number, and category. Finding a book now takes only a few seconds. The books didn't change, only the **way they were organized** changed. Programming is exactly the same.

**Definition:** A **data structure** is a way of organizing and storing data so that it can be used efficiently.

Notice something: a data structure is **not** data, it is the **organization of the data**. Think of it like this: data is the clothes, and the data structure is the wardrobe. The same clothes can either be folded neatly or scattered all over the room. The clothes haven't changed, only the organization has.

### Why Data Structures Matter

Suppose Facebook has over two billion users. Imagine if every time someone logged in, Facebook searched every user one after another (User 1, then User 2, then User 3... all the way to User 2,000,000,000). Logging in would take forever. Instead, Facebook stores data using efficient data structures, which is why searching feels almost instant.

### Real Life Analogy

Imagine your phone contacts. You don't save them as "Friend 1, Friend 2, Friend 3." Instead, you organize contacts so you can search instantly. Your contacts are stored using data structures.

### Another Analogy

Think of a supermarket. Different products belong in different places: fruits go to the fruit section, milk goes to the dairy section, bread goes to the bakery, and drinks go to the beverage section. Everything has an appropriate place, and programs should organize data the same way.

### AI Builder Perspective

As an AI Builder, you'll often work with user information, chat history, images, documents, API responses, sensor readings, and machine learning datasets. One of your first questions should always be: **"How should I store this information?"** That question is about data structures.

---

### Chapter 2 — Why Choosing the Right Data Structure Matters

Imagine someone asks you to carry water. Would you use a basket or a bucket? Both can carry things, but only one is appropriate. Programming is the same: different problems require different data structures. Choosing the wrong one can make a program slower, harder to understand, and wasteful of memory. Choosing the right one makes the program faster, cleaner, and easier to maintain.

---

### Chapter 3 — Linear vs Non Linear Data Structures

There are many data structures, but they are often grouped into two broad categories.

**Linear Data Structures** — data is arranged one after another: `A → B → C → D`. Examples: Arrays, Lists, Linked Lists, Stacks, Queues.

**Non Linear Data Structures** — data branches into multiple paths:
```
       A
      / \
     B   C
    / \   \
   D   E   F
```
Examples: Trees, Graphs, Heaps, Tries.

We are focusing only on **linear** structures in this part.

---

### Chapter 4 — Arrays

Arrays are usually the first data structure programmers learn, because they are simple.

**Imagine this:** Suppose you have five lockers (Locker 0 through Locker 4). Each locker stores exactly one item. To access an item, you simply provide the locker number. Arrays work exactly like this.

**Definition:** An **array** is a collection of elements stored in consecutive memory locations. The important phrase is **consecutive memory locations**. Imagine houses on a street, each house beside the next one; arrays are stored like that inside memory.

**Visual Representation:**
```
Index
0     1     2     3     4
+-----+-----+-----+-----+-----+
| 12  | 25  | 18  | 40  | 90  |
+-----+-----+-----+-----+-----+
```
Notice that each value has an **index**, which is its address.

**Accessing an Element**

```python
numbers = [12,25,18,40,90]
print(numbers[2])
```
Output: `18` (because counting starts at 0, 1, 2).

**Why Arrays Are Fast**

Imagine I ask, "What is inside locker number 3?" You don't open lockers 0, 1, 2 first, you go directly to 3. That is why arrays are very fast for accessing data.

**Common Operations on Arrays**

- **Read** — very fast, e.g. `numbers[5]`
- **Update** — very fast, e.g. `numbers[5] = 80`
- **Search** — may require checking many elements
- **Insert** — can be slow, because other elements may need to move
- **Delete** — also can be slow, for the same reason

**Example:** Original list `10, 20, 30, 40`. Insert `15`. Result: `10, 15, 20, 30, 40`. Notice everything after 10 had to shift.

---

### Chapter 5 — Lists (Dynamic Arrays)

Python programmers often use Lists. Technically, Python Lists are **dynamic arrays**.

**What Does Dynamic Mean?** Imagine buying a notebook with only 100 pages. When it becomes full, you buy another notebook, similar to a normal array. Now imagine a magical notebook where new pages appear automatically whenever it becomes full. That is a dynamic array.

```python
numbers = [5,10]
numbers.append(15)
numbers.append(20)
```
The list grows automatically.

**Why Dynamic Arrays Are Useful:** You don't need to know beforehand how many items you'll store. Social media comments, chat messages, emails, and products can all grow over time.

**Arrays vs Dynamic Arrays**

| Feature | Array | Dynamic Array (List) |
|---------|-------|----------------------|
| Fixed Size | ✅ Yes | ❌ No |
| Can Grow | ❌ No | ✅ Yes |
| Fast Access | ✅ Yes | ✅ Yes |
| Easy to Use | Moderate | Very Easy |

---

### Chapter 6 — Linked Lists

This is where many beginners become confused. Let's simplify it.

**Imagine a Treasure Hunt:** Each clue tells you where the next clue is (Clue 1 → go to Tree → Clue 2 → go to House → Clue 3 → go to River). Each clue points to another clue. A linked list works exactly like this.

**Definition:** A linked list is a collection of nodes where each node stores data and a reference (pointer) to the next node.

**Visual Representation:**
```
+-------+-------+        +-------+-------+        +-------+-------+
| 10 |  •----|-------->  | 25 |  •----|-------->  | 40 | NULL |
+-------+-------+        +-------+-------+        +-------+-------+
```
The last node points to `NULL`, meaning "there is no next node."

**Why Linked Lists Exist:** Imagine a train, where each carriage connects to the next (Engine → Coach → Coach → Coach). You can add another coach without rebuilding the train. Linked lists work similarly.

**Arrays vs Linked Lists**

Imagine inserting a new student at the beginning. With an array (`Ali, Musa, Fatima`), inserting `Amina` at the front means everything shifts. With a linked list, you simply connect one new node, with no shifting required.

**Which Is Better?** Neither, each has strengths. Arrays offer fast access but slow insertion. Linked lists offer slower access but faster insertion. Professional developers choose based on the problem.

---

### Chapter 7 — Strings as Data Structures

Many beginners think strings are just text, but internally they are collections of characters. For example, `HELLO` can be viewed as index 0 → H, 1 → E, 2 → L, 3 → L, 4 → O.

```python
word = "HELLO"
print(word[1])
```
Output: `E`

**String Operations**

- Length: `len(word)`
- Looping: `for letter in word: print(letter)`
- Searching: `"H" in word`
- Slicing: `word[1:4]` → Output: `ELL`

---

### Choosing Between These Data Structures

- Use an **Array** when you know the approximate size and need very fast access by index.
- Use a **Dynamic Array (List)** when the number of items can grow or shrink, but you still want easy index based access.
- Use a **Linked List** when you frequently insert or remove items, especially at the beginning or middle.
- Use a **String** when you are storing and manipulating text.

---

### Time Complexity (Beginner Introduction)

One of the biggest questions in programming is: "How much work does this operation require?" For now, don't worry about Big O notation, we'll study it deeply in Module 4. Instead, think in simple terms:

| Operation | Array/List | Linked List |
|-----------|------------|-------------|
| Access by Index | ⭐ Very Fast | ❌ Slow |
| Search | 😐 Medium | 😐 Medium |
| Insert at End | ⭐ Fast (usually) | ⭐ Fast |
| Insert at Beginning | ❌ Slow | ⭐ Fast |
| Delete at Beginning | ❌ Slow | ⭐ Fast |

The exact notation (like O(1) or O(n)) will make much more sense after you've mastered these concepts.

---

### Hands on Exercises

1. Create a list of your five favorite programming languages. Print the first and last language.
2. Create a list of ten numbers. Calculate their sum using a loop.
3. Given `names = ["Mudi", "Amina", "John", "Fatima"]`, print every name, then print only the first character of each name.
4. Without using Python's built in `len()` function, count the number of characters in a string using a loop.
   Hint:
   ```python
   count = 0
   for character in word:
       count += 1
   ```

---

### Logical Reasoning Challenges

**Challenge 1:** Suppose you have one million student records. You need to frequently look up students by their position (index), and rarely insert new students. Would you choose an Array or a Linked List?

✅ **Answer: Array.** Accessing elements by index is extremely fast in an array.

**Challenge 2:** Suppose you are building a music playlist where users constantly add, remove, and reorder songs. Which data structure would be more suitable?

**Possible Answer:** A linked list can be a good choice because inserting and removing items is efficient. (In practice, many applications use more advanced or hybrid structures, but for assessment purposes, understanding *why* a linked list is suitable is the key.)

---

### AI Builder Assessment Questions

1. What is a data structure?
   A. A programming language  B. A way of organizing and storing data  C. A database  D. A compiler
   ✅ **B**

2. Which data structure stores elements in consecutive memory locations?
   A. Graph  B. Linked List  C. Array  D. Tree
   ✅ **C**

3. Which data structure automatically grows when more elements are added?
   A. Fixed Array  B. Dynamic Array (List)  C. Binary Tree  D. Stack
   ✅ **B**

4. In a linked list, each node contains:
   A. Only data  B. Data and a reference to the next node  C. Two arrays  D. A database connection
   ✅ **B**

5. Which statement about strings is true?
   A. Strings are numbers.  B. Strings are collections of characters.  C. Strings cannot be indexed.  D. Strings only exist in Python.
   ✅ **B**

---

### Common Beginner Mistakes

❌ Thinking an array and a Python list are exactly the same thing in every programming language.
❌ Forgetting that indexing starts at **0**.
❌ Using a linked list when fast random access is required.
❌ Assuming every data structure is equally good for every problem.
❌ Memorizing definitions without understanding *why* one structure is chosen over another.

### Builder Tips

As an AI Builder, you'll often receive data from APIs, machine learning models, or databases. Before writing code, pause and ask: Do I need fast lookups? Will this collection grow over time? Will I insert and delete items frequently? Is the order of items important? These questions will naturally guide you toward the right data structure.

Professional developers don't ask "What's my favorite data structure?" They ask, **"Which data structure best fits this problem?"**

### Builder Checkpoint

Before moving to Part 2, make sure you can confidently say:

✅ I understand what a data structure is.
✅ I know why organizing data matters.
✅ I understand the difference between linear and non linear data structures.
✅ I can explain how arrays work.
✅ I understand what dynamic arrays (lists) are.
✅ I understand the basic idea behind linked lists.
✅ I know that strings are collections of characters.
✅ I can choose an appropriate data structure for simple scenarios.

### Fast Review

- A **data structure** organizes data for efficient use.
- Choosing the right data structure can make a program faster and easier to maintain.
- **Linear data structures** arrange data in sequence.
- **Arrays** provide fast access by index but are expensive to insert into the middle or beginning.
- **Dynamic arrays (lists)** automatically grow as needed.
- **Linked lists** connect nodes together, making insertion and deletion easier but indexed access slower.
- **Strings** are ordered collections of characters that support indexing, iteration, searching, and slicing.
- Understanding *when* to use each structure is more valuable than simply memorizing definitions.

---

## Part 2 — Specialized Data Structures

**Estimated Study Time:** 6–7 Hours
**Difficulty:** ⭐⭐⭐☆☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Explain what a Stack is and where it is used.
✅ Explain what a Queue is and where it is used.
✅ Understand Double Ended Queues (Deque).
✅ Understand Priority Queues.
✅ Understand Hash Tables (Maps/Dictionaries).
✅ Understand Sets.
✅ Choose the correct data structure for different real world problems.

---

### Chapter 1 — Why Specialized Data Structures Exist

In Part 1, we learned that different problems require different data structures. Now imagine these situations: a browser remembers every page you visited, a printer prints documents one after another, a hospital attends to emergency patients before less critical ones, a phone stores contacts by name, and a classroom attendance sheet should not contain duplicate names. Would one data structure solve all these problems? No, each problem requires a different way of organizing data. That is why specialized data structures exist.

---

### Chapter 2 — Stack

**Imagine a Stack of Plates:** Suppose you wash five plates and stack them (Plate 5 on top down to Plate 1 at the bottom). When you need one plate, which one do you remove? The top plate, you cannot remove Plate 2 first. This is exactly how a Stack works.

**Definition:** A **Stack** is a data structure that follows the **Last In, First Out (LIFO)** principle. The last item added is the first one removed.

**Visual Representation:** Push 10 → stack is `[10]`. Push 20 → stack is `[20, 10]`. Push 30 → stack is `[30, 20, 10]`. Pop removes 30 first, leaving `[20, 10]`. The last item (30) is removed first.

**Stack Operations**

- **Push** — adds an item to the top.
- **Pop** — removes the top item.
- **Peek (Top)** — looks at the top item without removing it.
- **Is Empty** — checks whether the stack contains anything.

```python
stack = []
stack.append(10)
stack.append(20)
stack.append(30)
print(stack)
stack.pop()
print(stack)
```
Output: `[10,20,30]` then `[10,20]`

**Real World Applications**

- **Browser Back Button:** Visiting Google → YouTube → GitHub → ChatGPT, then pressing Back makes ChatGPT disappear first, exactly like a stack.
- **Undo Feature:** Typing "Hello" then pressing Undo makes the last action disappear.
- **Programming Function Calls:** When functions call other functions, they are managed using a Call Stack.

**Builder Tip:** Whenever you hear "Undo" or "Go Back," think **Stack**.

---

### Chapter 3 — Queue

Imagine people buying tickets: Person A, then Person B, then Person C. Who gets served first? Person A, not Person C.

**Definition:** A Queue follows the **First In, First Out (FIFO)** principle. The first person entering is the first person leaving.

**Visual Representation:** Front → A, B, C → Back. After removing one item, A leaves and the front becomes B, C.

**Queue Operations**

- **Enqueue** — add to the back.
- **Dequeue** — remove from the front.
- **Front** — look at the first item.
- **Is Empty** — check whether the queue is empty.

```python
from collections import deque
queue = deque()
queue.append("A")
queue.append("B")
queue.append("C")
print(queue)
queue.popleft()
print(queue)
```

**Real World Applications:** Hospital Queue, Bank Queue, Airport Check in, Customer Support Tickets, Printing Documents.

**Builder Tip:** Whenever you hear "First Come, First Served," think **Queue**.

**Stack vs Queue**

| Stack | Queue |
|---------|--------|
| Last In, First Out | First In, First Out |
| Push | Enqueue |
| Pop | Dequeue |
| Undo | Waiting Line |

---

### Chapter 4 — Double Ended Queue (Deque)

Sometimes you need to insert and remove items from both ends. Imagine a train where passengers may enter from the front or the back. A Deque allows this.

**Operations:** Add Front, Add Back, Remove Front, Remove Back.

```python
from collections import deque
dq = deque()
dq.append(10)
dq.appendleft(5)
dq.pop()
dq.popleft()
```

**Real World Example:** A music playlist, where you may skip forward, go back, insert a song at the beginning, or insert a song at the end.

---

### Chapter 5 — Priority Queue

Imagine a hospital where Patient A has a minor injury, Patient B is having a heart attack, and Patient C has a broken finger. Should they be treated in arrival order? No, emergency patients come first. That is a Priority Queue.

**Definition:** A Priority Queue removes the highest priority item first, not necessarily the oldest one.

**Visual Example:** Priority 5 → Heart Attack, Priority 3 → Broken Leg, Priority 1 → Cold. The highest priority is served first.

**Real World Applications:** Hospital Emergency Rooms, CPU Task Scheduling, Network Routers, AI Search Algorithms, Operating Systems.

---

### Chapter 6 — Hash Tables (Maps/Dictionaries)

One of the most important data structures. Imagine an English dictionary: you search "Apple" and immediately find its meaning, you don't read every page. Hash tables work similarly.

**Definition:** A Hash Table stores Key → Value pairs.

```python
student = {
    "name": "Mudi",
    "age": 25,
    "country": "Nigeria"
}
print(student["country"])
```
Output: `Nigeria`

**Real World Applications:** Phone Contacts (Name → Phone Number), Student Records (Student ID → Student Information), API Responses, JSON Data, User Profiles, Caching, AI Applications.

**Why Hash Tables Are Fast:** Suppose you know someone's phone number. Would you read every contact? No, the name points directly to the information. Hash tables work similarly.

---

### Chapter 7 — Sets

Imagine attendance recorded as: Mudi, Mudi, Mudi, John, John, Fatima. Attendance should really just be: Mudi, John, Fatima. Duplicates disappear.

**Definition:** A Set stores only unique values.

```python
languages = {
    "Python",
    "Go",
    "Python",
    "JavaScript"
}
print(languages)
```
Possible Output: `{'Python','Go','JavaScript'}`

**Common Set Operations:** Add, Remove, Check Membership, Union, Intersection, Difference.

```python
A = {1,2,3}
B = {3,4,5}
print(A & B)
```
Output: `{3}`

**Real World Applications:** Unique Visitors, Unique Email Addresses, Unique Product IDs, Removing Duplicate Records, Machine Learning Dataset Cleaning.

---

### Choosing the Right Data Structure

- Need an Undo Button → **Stack**
- Need a Waiting Line → **Queue**
- Need Emergency Priority → **Priority Queue**
- Need Fast Lookup → **Hash Table**
- Need Unique Values → **Set**
- Need Insert/Remove at Both Ends → **Deque**

**Comparison Table**

| Data Structure | Main Idea | Real Example |
|---------------|-----------|--------------|
| Stack | LIFO | Browser Back |
| Queue | FIFO | Bank Line |
| Deque | Both Ends | Playlist |
| Priority Queue | Highest Priority First | Hospital |
| Hash Table | Key → Value | Contacts |
| Set | Unique Values | Attendance |

---

### Hands on Exercises

1. Implement a stack using a Python list. Perform Push, Pop, and Peek.
2. Implement a queue using `collections.deque`. Add A, B, C, then remove one item and observe the result.
3. Create a dictionary representing a student, including Name, Course, and Score. Retrieve each value using its key.
4. Create two sets, `A = {1,2,3,4}` and `B = {3,4,5,6}`, and find their Union, Intersection, and Difference.

---

### Logical Reasoning Challenges

**Challenge 1:** A browser stores every page you visit. Which data structure should it use?
A. Queue  B. Stack  C. Set  D. Dictionary
✅ **Answer: Stack.** The most recently visited page is the first one returned when you click Back.

**Challenge 2:** A printer receives jobs in the order Document A, Document B, Document C. Which document should print first?
✅ **Answer: Document A**, because printers normally use a Queue.

**Challenge 3:** A hospital receives patients in this order: Cold, Heart Attack, Broken Finger. Who should be treated first?
✅ **Answer: Heart Attack**, because hospitals use priority based systems.

---

### AI Builder Assessment Questions

1. Which data structure follows Last In, First Out?
   A. Queue  B. Stack  C. Set  D. Dictionary
   ✅ **B**

2. Which data structure follows First In, First Out?
   A. Queue  B. Stack  C. Tree  D. Set
   ✅ **A**

3. Which data structure stores unique values?
   A. Stack  B. Queue  C. Set  D. Array
   ✅ **C**

4. Which data structure stores Key → Value pairs?
   A. Queue  B. Hash Table  C. Stack  D. Tree
   ✅ **B**

5. Which specialized data structure allows insertion and deletion from both ends?
   A. Deque  B. Priority Queue  C. Stack  D. Set
   ✅ **A**

6. Which data structure is most suitable for implementing an Undo feature?
   A. Queue  B. Stack  C. Dictionary  D. Set
   ✅ **B**

---

### Common Beginner Mistakes

❌ Thinking a Stack and Queue work the same way.
❌ Using a List when a Set is needed to remove duplicates.
❌ Using a Queue for Undo operations.
❌ Forgetting that dictionaries are designed for fast lookups using keys.
❌ Assuming Priority Queues always process items in the order they arrive.

### Builder Tips

As an AI Builder, you'll frequently use these structures without realizing it: chat history can be represented as a sequence of messages, API responses are commonly dictionaries (JSON objects), removing duplicate training examples often uses sets, AI task schedulers may rely on priority queues, and function calls inside your programs use a stack automatically. Understanding these structures helps you understand not only your own code, but also the tools and frameworks you'll work with.

### Builder Checkpoint

Before moving to Part 3, make sure you can confidently say:

✅ I understand the difference between LIFO and FIFO.
✅ I know when to use a Stack.
✅ I know when to use a Queue.
✅ I understand what a Deque is.
✅ I understand why Priority Queues exist.
✅ I understand how Hash Tables (Dictionaries) store Key → Value pairs.
✅ I know that Sets automatically remove duplicates.
✅ I can choose the correct specialized data structure for common real world problems.

### Fast Review

- A **Stack** follows **Last In, First Out (LIFO)**.
- A **Queue** follows **First In, First Out (FIFO)**.
- A **Deque** allows insertion and deletion at both the front and the back.
- A **Priority Queue** processes items based on priority instead of arrival time.
- A **Hash Table (Dictionary)** stores **Key → Value** pairs and provides fast lookups.
- A **Set** stores only unique values and is useful for removing duplicates.
- Choosing the right data structure depends on the problem you're solving, not on personal preference.

---

## Part 3 — Non Linear Data Structures

**Estimated Study Time:** 7–8 Hours
**Difficulty:** ⭐⭐⭐⭐☆ (Intermediate)

### Learning Objectives

By the end of this part, you should be able to:

✅ Explain what non linear data structures are.
✅ Understand Trees and why they exist.
✅ Understand Binary Trees.
✅ Understand Binary Search Trees (BST).
✅ Understand Heaps.
✅ Understand Tries (Prefix Trees).
✅ Understand Graphs.
✅ Understand Breadth First Search (BFS).
✅ Understand Depth First Search (DFS).
✅ Recognize real world applications of each.

---

### Chapter 1 — What Are Non Linear Data Structures?

In Part 1, we studied **linear** data structures, where data is arranged one after another: `A → B → C → D`. Every element has only one path forward.

Now imagine something different, like a family tree:
```
        Grandfather
         /       \
     Father      Uncle
      /   \
   You   Sister
```
Notice that one person can connect to many others. This is **non linear**.

**Definition:** A **non linear data structure** stores data in a hierarchical or network like structure rather than in a single sequence.

**Why Do We Need Them?** Imagine trying to represent family relationships, company organization, computer folders, road maps, or Facebook friendships using a normal list. A simple list isn't enough, we need structures that allow **branching**.

---

### Chapter 2 — Trees

Real trees grow downward, but computer trees grow upward:
```
        A
      / | \
     B  C  D
```

**Definition:** A **Tree** is a hierarchical data structure made of nodes connected by edges.

**Important Terminology**

- **Root** — the very first node (A in `A`).
- **Parent** — a node that has children (A is B's parent in `A / B`).
- **Child** — a node connected below another (B is the child in `A | B`).
- **Leaf** — a node with no children. In the tree with A at the root, B, and C splitting into D and E, the leaves are B, D, and E.
- **Edge** — the connection between nodes, e.g. `A ----- B`.
- **Height** — the longest path from the root to a leaf.

**Real Life Examples:** Computer folders (Documents → School, Pictures, Music), company hierarchy (CEO → Managers → Employees), restaurant menus, website navigation, and organization charts.

---

### Chapter 3 — Binary Trees

Suppose every parent is allowed only **two children**:
```
        A
       / \
      B   C
```
That's a Binary Tree.

**Definition:** A Binary Tree is a tree where each node has **at most two children**, called the Left Child and Right Child.

Example:
```
          8
        /   \
       4     10
      / \      \
     2   6      15
```

**Why Binary Trees Are Popular:** Many searching algorithms become much faster, many databases use tree like structures, and many AI systems organize decision making using trees.

---

### Chapter 4 — Binary Search Tree (BST)

Now imagine we introduce one simple rule: every smaller number goes left, every larger number goes right.

Inserting `50, 30, 70, 20, 40, 60, 80` produces:
```
          50
        /    \
      30      70
     /  \    /  \
   20   40 60   80
```
Notice the pattern: everything left is smaller, everything right is larger.

**Searching in a BST:** Suppose we want to find `60`. Start at 50. Is 60 greater? Yes, go right to 70. Is 60 smaller? Yes, go left to 60. Found. Instead of checking every number, the tree eliminates half of the search space each step.

**Real Life Example:** Imagine searching a dictionary. You don't start at page one, you repeatedly divide the search space. BSTs work similarly.

---

### Chapter 5 — Heap

Many beginners confuse Heaps with BSTs, but they are different. Imagine a company where the CEO always has the highest authority and employees below have lower authority, similar to a Max Heap.

**Definition:** A Heap is a specialized tree where the parent always has a higher (Max Heap) or lower (Min Heap) priority than its children.

Max Heap example (parent always larger):
```
        100
       /   \
      80    60
     / \
   40  30
```

Min Heap example (parent always smaller):
```
        5
      /   \
     10    20
    / \
  30  40
```

**Where Are Heaps Used?** Priority Queues, Task Scheduling, Operating Systems, AI Search Algorithms, CPU Scheduling.

---

### Chapter 6 — Trie (Prefix Tree)

Imagine your phone keyboard: you type "Cha" and it immediately suggests "Chat," "ChatGPT," "Chair," and "Challenge." How? Using a Trie.

**Definition:** A Trie stores words by sharing common prefixes.

Example, where CHAT, CHEF, and CHOIR share the prefix CH:
```
          Root
          |
          C
          |
          H
        / | \
      A  E  O
      |
      T
```

**Real Life Uses:** Search Suggestions, Auto complete, Spell Checking, Dictionary Apps, Search Engines.

---

### Chapter 7 — Graphs

Suppose you have cities: Lagos, Abuja, Kano, Kaduna. Roads connect them, but not every city connects directly. This is a Graph.

**Definition:** A Graph consists of Nodes (Vertices) and Connections (Edges).

Example:
```
A ----- B
|       |
|       |
C ----- D
```

Graphs do NOT have a root, graphs may contain cycles, and graphs can connect in many ways.

**Real Life Applications:** Google Maps, Facebook Friend Network, Airline Routes, Internet Connections, Electric Power Grids, Recommendation Systems.

---

### Chapter 8 — Graph Traversal

Suppose we want to visit every city. Two famous methods exist.

**Breadth First Search (BFS):** Think of throwing a stone into water, the ripples spread outward.
```
        A
      /   \
     B     C
    / \   / \
   D  E  F  G
```
BFS visits A, B, C, D, E, F, G, one level at a time.

**Where BFS Is Used:** GPS, Finding Shortest Paths, Social Networks, Web Crawlers, Recommendation Systems.

**Depth First Search (DFS):** Imagine exploring a cave, you continue going deeper until you cannot continue, then you return. DFS explores one branch completely before returning.

**Where DFS Is Used:** Maze Solving, Puzzle Solving, Finding Connected Components, AI Game Search, Backtracking Algorithms.

**BFS vs DFS**

| BFS | DFS |
|------|------|
| Visits level by level | Goes deep first |
| Uses a Queue | Uses a Stack |
| Finds shortest path in unweighted graphs | Uses less memory in some deep searches |
| Good for nearby exploration | Good for exhaustive exploration |

---

### Choosing the Right Structure

| Problem | Data Structure |
|----------|----------------|
| Folder System | Tree |
| Database Index | BST |
| Task Priority | Heap |
| Auto complete | Trie |
| Social Network | Graph |
| GPS Navigation | Graph |
| Search Suggestions | Trie |

---

### Hands on Exercises

1. Draw a family tree with a Grandparent, Parent, and two children. Label the Root, Parent, Child, and Leaf.
2. Insert these values into a Binary Search Tree: `40, 20, 60, 10, 30, 50, 70`. Draw the final tree.
3. Suppose your contacts app must suggest names as you type. Which data structure is most appropriate? Explain why.
4. Draw a graph representing the cities Lagos, Abuja, Kano, and Kaduna, connecting any cities you think have roads between them. Then perform a simple BFS traversal starting from Lagos.

---

### Logical Reasoning Challenges

**Challenge 1:** A company wants to represent its organizational structure. Should it use a Queue, Tree, Stack, or Set?
✅ **Answer: Tree.** Organizations naturally have parent child relationships.

**Challenge 2:** Google Maps must calculate routes between cities. Which data structure best models this?
A. Stack  B. Queue  C. Graph  D. Array
✅ **Answer: C, Graph.**

**Challenge 3:** Your phone instantly suggests words while typing. Which data structure makes this efficient?
A. Queue  B. Trie  C. Heap  D. Stack
✅ **Answer: B, Trie.**

---

### AI Builder Assessment Questions

1. A Binary Tree allows each node to have at most:
   A. One child  B. Two children  C. Three children  D. Unlimited children
   ✅ **B**

2. In a Binary Search Tree, smaller values are stored:
   A. Above the root  B. On the right  C. On the left  D. Randomly
   ✅ **C**

3. Which data structure is best for implementing auto complete?
   A. Graph  B. Queue  C. Trie  D. Stack
   ✅ **C**

4. Which traversal explores one level at a time?
   A. DFS  B. BFS  C. Heap Search  D. Trie Search
   ✅ **B**

5. Which traversal goes as deep as possible before backtracking?
   A. BFS  B. DFS  C. Binary Search  D. Heap Traversal
   ✅ **B**

6. Which data structure naturally models friendships in a social network?
   A. Array  B. Stack  C. Graph  D. Queue
   ✅ **C**

---

### Common Beginner Mistakes

❌ Confusing a Tree with a Graph.
❌ Thinking every Tree is a Binary Tree.
❌ Assuming BSTs and Heaps follow the same rules.
❌ Forgetting that Graphs can contain cycles.
❌ Thinking BFS and DFS visit nodes in the same order.
❌ Trying to memorize traversals without understanding *why* they are useful.

### Builder Tips

You may not implement these data structures from scratch every day, but you'll use systems that rely on them: search engines use **Tries** for suggestions, databases often use **Trees** to index data, navigation apps use **Graphs** to compute routes, AI planning and game algorithms use **Trees** and **Graphs** to explore possible decisions, and task schedulers and priority systems often rely on **Heaps**. Recognizing *which* data structure is being used, and *why*, is a valuable skill for an AI Builder.

### Builder Checkpoint

Before moving to Part 4, make sure you can confidently say:

✅ I understand what non linear data structures are.
✅ I can identify the parts of a tree (root, parent, child, leaf).
✅ I know what makes a Binary Tree different.
✅ I understand the Binary Search Tree property.
✅ I know the purpose of Heaps.
✅ I understand how Tries enable auto complete.
✅ I know what a Graph represents.
✅ I understand the basic ideas behind BFS and DFS.

### Fast Review

- **Trees** organize data hierarchically.
- **Binary Trees** allow at most two children per node.
- **Binary Search Trees (BSTs)** keep smaller values on the left and larger values on the right.
- **Heaps** maintain a priority relationship between parents and children.
- **Tries** store words efficiently by sharing prefixes.
- **Graphs** model networks such as roads, friendships, and the internet.
- **BFS** explores level by level using a **Queue**.
- **DFS** explores deeply before backtracking and is commonly implemented using a **Stack**.
- Non linear data structures are essential for solving many real world AI and software engineering problems efficiently.

---

## Part 4 — Thinking Like an AI Builder with Data Structures

**Estimated Study Time:** 7–9 Hours
**Difficulty:** ⭐⭐⭐⭐☆ (Beginner → Intermediate)

### Learning Objectives

By the end of this chapter, you should be able to:

✅ Think like a software engineer when solving problems.
✅ Choose the appropriate data structure for different scenarios.
✅ Read interview questions without panicking.
✅ Develop intuition for algorithmic thinking.
✅ Understand the basics of Big O Notation.
✅ Understand Time Complexity.
✅ Understand Space Complexity.
✅ Avoid common mistakes beginners make.

---

### Chapter 1 — Stop Memorizing. Start Thinking.

One of the biggest mistakes beginners make is trying to memorize data structures. Professionals don't think this way, instead they ask one simple question: **"What problem am I trying to solve?"** The data structure comes **after** understanding the problem.

**Example:** Suppose you're building a calculator. Would you use a Graph? No. Would you use a Trie? No. A few variables are enough. Now suppose you're building Google Maps. Would variables be enough? No, you need a **Graph**. Different problems require different tools.

**The AI Builder Mindset:** Instead of asking "What data structure should I memorize?" ask "What kind of data am I working with?" For example: Text → String, Unique Items → Set, Name → Phone Number → Dictionary, Undo Button → Stack, Waiting Line → Queue, Road Network → Graph, Folder Structure → Tree.

---

### Chapter 2 — How Interview Questions Are Actually Testing You

Suppose an interviewer says "Design a browser's Back button." They are NOT asking you to build Chrome, they're asking whether you can recognize that this is a Stack.

Another question: "Build a printer queue." They don't care about printers, they're checking whether you understand Queues.

Another: "Suggest words while the user types." They're checking whether you know Tries.

Most interview questions are really asking: **"Can you recognize the hidden data structure?"**

---

### Chapter 3 — Choosing the Right Data Structure

A simple decision guide:

1. Do you need duplicate values? Yes → Use List or Array. No → Use Set.
2. Do you need fast lookup by key? Yes → Dictionary (Hash Table).
3. Do you frequently add and remove from the beginning? Yes → Linked List or Deque.
4. Do you process items in order of arrival? Yes → Queue.
5. Do you always need the most recently added item? Yes → Stack.
6. Are relationships important? Yes → Tree or Graph.

**Decision Tree**

```
Need unique values?
↓
YES → Set
↓
NO
↓
Need key → value?
↓
YES → Dictionary
↓
NO
↓
Need LIFO?
↓
YES → Stack
↓
NO
↓
Need FIFO?
↓
YES → Queue
↓
NO
↓
Need hierarchy?
↓
YES → Tree
↓
NO
↓
Need network?
↓
YES → Graph
```

---

### Chapter 4 — Reading Data Structure Questions

Many assessment questions hide the answer in a keyword:

- "Users join a waiting room." Keyword: waiting → think **Queue**.
- "Remove duplicate emails." Keyword: duplicate → think **Set**.
- "Store user profile by ID." Keyword: ID → think **Dictionary**.
- "Undo previous action." Keyword: undo → think **Stack**.
- "Recommend friends." Keyword: relationship → think **Graph**.

---

### Chapter 5 — Common Problem Patterns

You don't need hundreds of algorithms, many interview questions fall into patterns:

- **Searching** (find a name, number, or student) often uses Dictionary, BST, or Binary Search.
- **Ordering** (sort scores, products, prices) often uses sorting algorithms.
- **Traversal** (visit every folder, city, or employee) often uses DFS or BFS.
- **Removing Duplicates** often uses a Set.
- **Counting** (words, characters, products) often uses a Dictionary.

---

### Chapter 6 — Beginner Friendly Introduction to Big O

This is one of the most feared topics, but it doesn't have to be.

**What is Big O?** Big O describes **how the amount of work grows as the amount of data grows**. It does NOT measure seconds or CPU speed, it measures **growth**.

Imagine three people: Person A checks one book, Person B checks ten books, Person C checks one million books. Who does more work? Obviously Person C. Big O simply gives us a language for describing that growth.

**Imagine Searching a Phone Contact:** Suppose you want to find "Mudi." Method 1 is reading every contact (Ali, Amina, John, Mudi), where worst case you read all contacts and work grows with the number of contacts. Method 2 is using a Dictionary to jump directly, needing only one lookup, which is much faster. This is why data structures matter.

**Common Big O Notations**

- **O(1) — Constant Time:** The amount of work never changes. For example, `numbers[5]` takes roughly the same work whether the array has 10 items or 10 million, like opening locker number 5, you go directly there.
- **O(n) — Linear Time:** You may need to check every item. For example, finding "Mudi" inside an unsorted list (Ali, John, Musa, Fatima, Mudi) means worst case you inspect every name.
- **O(log n) — Logarithmic Time:** Instead of checking one by one, you keep cutting the search space in half, like finding a word in a dictionary by repeatedly dividing the remaining pages. This idea powers **Binary Search**.
- **O(n²) — Quadratic Time:** Every item is compared with every other item, like every student shaking hands with every other student. As the class grows, the number of handshakes grows much faster.

**Visual Intuition:** O(1) always stays tiny. O(log n) grows very slowly. O(n) grows steadily. O(n²) gets large very quickly.

---

### Chapter 7 — Time Complexity of Common Operations

| Data Structure | Access | Search | Insert | Delete |
|----------------|--------|--------|--------|--------|
| Array | Fast | Medium | Slow (middle) | Slow (middle) |
| Dynamic Array | Fast | Medium | Usually Fast (end) | Medium |
| Linked List | Slow | Medium | Fast | Fast |
| Stack | Fast (Top) | Not Typical | Fast | Fast |
| Queue | Fast (Ends) | Not Typical | Fast | Fast |
| Dictionary | Very Fast | Very Fast | Very Fast | Very Fast |
| Set | Very Fast | Very Fast | Very Fast | Very Fast |
| Binary Search Tree* | Fast (average) | Fast (average) | Fast (average) | Fast (average) |

> *Balanced BSTs achieve these average efficiencies. A poorly balanced BST can become much slower.

Don't memorize this table today, understand **why**.

---

### Chapter 8 — Space Complexity

Time asks "How much work?" Space asks "How much memory?" For example, `numbers = [1,2,3]` uses memory for 3 numbers, while `numbers = [1,2,3,4,5,6,7,8]` needs more memory. As data grows, memory usage grows, and that's Space Complexity.

---

### Chapter 9 — Common Beginner Mistakes

❌ Memorizing Big O without understanding it.
❌ Choosing a data structure before understanding the problem.
❌ Thinking "fast" always means "best."
❌ Forgetting that readability matters.
❌ Ignoring memory usage.
❌ Using a List when a Dictionary is clearly better.

---

### Chapter 10 — Real AI Builder Scenarios

1. **A chatbot stores messages.** Possible Answer: List, since messages are kept in order.
2. **A chatbot stores user profiles by user ID.** Possible Answer: Dictionary (ID → Profile).
3. **An AI cleans duplicate email addresses.** Possible Answer: Set.
4. **A recommendation system models friendships.** Possible Answer: Graph.
5. **A search engine suggests words while typing.** Possible Answer: Trie.

---

### Hands on Exercises

1. For each problem below, choose the most appropriate data structure: Browser Back Button, Student Record by ID, Remove Duplicate Emails, GPS Navigation, Auto complete Search, Waiting Line, Undo Feature, Company Hierarchy.
2. Explain **why** you chose each answer. Do not simply write "Stack." Instead write something like: "I chose a Stack because the last page visited should be the first page returned when the user presses Back."
3. Imagine you're building a small AI powered library system. Choose suitable data structures for Books, Members, Borrowing Queue, Search Suggestions, and Categories, and explain your reasoning.

---

### AI Builder Assessment Questions

1. Which data structure is best for removing duplicates?
   A. Queue  B. Set  C. Stack  D. Array
   ✅ **B**

2. Which Big O notation describes work that stays approximately constant regardless of input size?
   A. O(n)  B. O(n²)  C. O(1)  D. O(log n)
   ✅ **C**

3. Binary Search has a time complexity of approximately:
   A. O(n)  B. O(log n)  C. O(n²)  D. O(1)
   ✅ **B**

4. Which question should you ask first when solving a programming problem?
   A. Which language should I use?  B. Which algorithm should I memorize?  C. What problem am I trying to solve?  D. Which IDE is the best?
   ✅ **C**

5. Which data structure best represents friendships on a social network?
   A. Graph  B. Queue  C. Stack  D. Array
   ✅ **A**

6. Space Complexity measures:
   A. Internet speed.  B. CPU temperature.  C. Memory required by an algorithm.  D. Screen resolution.
   ✅ **C**

---

### Mini Mock Assessment

1. A hospital treats the most critical patients first. Which data structure? ✅ **Priority Queue**
2. You need to look up a customer using their Account Number. Which data structure? ✅ **Dictionary (Hash Table)**
3. You want to remove duplicate usernames. Which data structure? ✅ **Set**
4. You need to repeatedly visit folders and subfolders. Which data structure naturally represents this? ✅ **Tree**
5. A navigation app computes routes between cities. Which data structure? ✅ **Graph**

---

### The AI Builder Checklist

Before moving to Module 4, you should be able to answer **YES** to all of these:

✅ I know what a data structure is.
✅ I know the difference between linear and non linear data structures.
✅ I understand Arrays, Lists, Linked Lists, Stacks, Queues, Dictionaries, Sets, Trees, Heaps, Tries, and Graphs.
✅ I can choose the correct data structure for common real world problems.
✅ I understand the intuition behind Big O.
✅ I know the difference between Time Complexity and Space Complexity.
✅ I no longer try to memorize data structures without understanding their purpose.

---

### Module 3 Cheat Sheet

| Problem | Best Data Structure |
|----------|---------------------|
| Ordered Items | List / Array |
| Dynamic Collection | List |
| Frequent Insert/Delete | Linked List |
| Undo | Stack |
| Waiting Line | Queue |
| Both Ends | Deque |
| Highest Priority First | Priority Queue |
| Fast Key → Value Lookup | Dictionary |
| Remove Duplicates | Set |
| Folder Structure | Tree |
| Database Index | Binary Search Tree |
| Auto complete | Trie |
| Friend Network | Graph |

---

### Module 3 Summary

Congratulations! You have completed **Module 3 — Data Structures**. You now understand:

- Why data structures exist.
- How to organize data effectively.
- The strengths and weaknesses of the most common data structures.
- How to recognize hidden data structure questions in interviews.
- The foundations of Big O Notation.
- The difference between Time Complexity and Space Complexity.
- How professional developers think when choosing a data structure.

This module gives you one of the strongest foundations for technical interviews and AI Builder assessments.

> **Remember this principle:** Good programmers know many data structures. Great programmers know when to use each one.