# AI BUILDER ASSESSMENT BOOTCAMP — MODULE 1

## Computational Thinking

| **Study Time** | **Difficulty** |
|:---|:---|
| 4–6 Hours | ⭐⭐☆☆☆ (Beginner → Intermediate) |

### Learning Objectives

By the end of this module, you'll be able to:

- Think like a programmer
- Break large problems into smaller ones
- Recognize patterns and similarities
- Remove unnecessary information
- Design simple algorithms
- Solve computational thinking questions confidently
- Read programming problems with clarity

---

# Chapter 1: What is Computational Thinking?

**Scenario:** Someone asks you to cook Jollof Rice.

Would you throw everything into the pot? Of course not! You'd naturally think through steps:
1. Gather ingredients → 2. Wash rice → 3. Prepare tomatoes → 4. Fry ingredients → 5. Add water → 6. Add rice → 7. Cook → 8. Serve

**Key insight:** You solved the problem *before* starting. That's **Computational Thinking**.

---

## Definition

**Computational Thinking** is the process of breaking a problem into manageable pieces, identifying patterns, removing unnecessary details, and creating step-by-step solutions.

It's the foundation of:
- Programming
- Artificial Intelligence
- Data Science
- Cybersecurity
- Robotics
- Software Engineering

**Remember:** If programming is *writing* instructions, Computational Thinking is *designing* those instructions.

---

## Why Does It Matter?

**Scenario:** Two developers face the same project.

- **Developer A** immediately starts coding
- **Developer B** spends ten minutes understanding the problem

**Who finishes first?** Developer B.

> **Why?** Coding is easy. Thinking is hard.
> 
> Professional programmers spend far more time thinking than typing.

---

## Real-Life Examples

**Going to School:** You naturally ask: *"What time should I leave? Which road is faster? Do I have money? Is it raining?"* Without realizing it, you're solving a problem using Computational Thinking.

**Google Maps:** When finding the fastest route, it doesn't randomly choose roads—it follows logical steps. That sequence of steps is an algorithm made possible by Computational Thinking.

**Washing Clothes:** You don't wash items randomly. You group similar items together because grouping saves time. **That's computational thinking.**

---

## The Four Pillars of Computational Thinking

Every programming problem can usually be solved using these four skills:

| # | Pillar | Definition |
|:---|:---|:---|
| **1** | **Decomposition** | Break problems into smaller pieces |
| **2** | **Pattern Recognition** | Find similarities between problems |
| **3** | **Abstraction** | Focus on important details, ignore the rest |
| **4** | **Algorithm Design** | Create step-by-step solutions |

Think of them as **four legs of a table**—remove one, and the table becomes unstable. We'll study each in detail.

---

---

# Pillar One: Decomposition

## What is Decomposition?

**Decomposition** means breaking a large problem into smaller, easier problems.

**Real-world example:** Moving into a new house. Instead of saying *"I'll move everything,"* you divide the work:
> Bedroom → Kitchen → Bathroom → Living room → Garage

**Suddenly, the task becomes easier.** That's decomposition.

---

## Programming Example

Instead of saying *"Build Facebook,"* break it into: Login System, Registration, User Profiles, Messaging, Notifications, Friends, News Feed.

**Professional software is never built all at once.** It's built one small piece at a time.

---

## Another Example

**Lecturer says:** *"Build an AI Chat Application."*

- **Beginner sees:** One giant task
- **Experienced developer sees:** Many smaller tasks

```
Frontend → Backend → Database → Authentication → AI API → Testing → Deployment
```

**That's decomposition.**

---

## ✅ Assessment Question

| | |
|:---|:---|
| **Question** | You've been asked to build an online bookstore. Which approach demonstrates decomposition? |
| **A** | Start coding immediately. |
| **B** | Break the project into Search, Cart, Payment, Login, and Orders. |
| **C** | Read motivational quotes. |
| **D** | Install ten programming languages. |
| **Answer** | ✅ **B** |

---

## Practice Questions

| **Q1** | Which best describes decomposition? |
|:---|:---|
| A | Ignoring the problem. |
| B | **Breaking a large problem into smaller parts.** |
| C | Programming faster. |
| D | Writing shorter code. |
| **Answer** | ✅ **B** |

| **Q2** | Which task uses decomposition? |
|:---|:---|
| A | Writing an entire OS in one file. |
| B | **Splitting an application into modules.** |
| C | Deleting the project. |
| D | Guessing the solution. |
| **Answer** | ✅ **B** |

---

> 💡 **Builder Tip:** When overwhelmed by a project, don't ask *"How do I build this?"* Instead ask *"What is the smallest piece I can build first?"* **This habit separates beginners from experienced developers.**

---

---

# Pillar Two: Pattern Recognition

## What is a Pattern?

Imagine you're a doctor. Ten patients come with Fever, Headache, and Weakness. **The symptoms repeat.** That repeating similarity is called a **pattern.**

**Programming works exactly the same way.**

---

## Example

You build a Login Page. Next month, you build another Login Page. You notice both require: Username, Password, Validation, Authentication.

**Instead of reinventing everything, you reuse the pattern.**

---

## Why Pattern Recognition Matters

Experienced programmers solve problems faster because they **recognize familiar patterns.**

Instead of thinking: *"I've never seen this,"*  
They think: *"This looks similar to something I've solved before."*

AI systems also work this way. **ChatGPT works largely because it recognizes patterns in language.** It predicts likely next words based on patterns learned during training.

---

## ✅ Assessment Question

| | |
|:---|:---|
| **Question** | Which demonstrates pattern recognition? |
| **A** | Creating a new solution every time. |
| **B** | **Recognizing that multiple applications require login functionality.** |
| **C** | Deleting repeated code without understanding it. |
| **D** | Ignoring similarities. |
| **Answer** | ✅ **B** |

---

## Pattern Practice

**Sequence 1:** 2, 4, 6, 8, 10, 12, ?
- **Answer:** 14
- **Pattern:** Add 2 each time

**Sequence 2:** 3, 6, 12, 24, 48, ?
- **Answer:** 96
- **Pattern:** Multiply by 2 each time

---

## Logical Reasoning Exercise

**Scenario:** A company notices that every time users increase, server response time also increases.

**Question:** What pattern is identified?

**Answer:** There's a **relationship between increasing users and slower server performance.** Recognizing this recurring relationship is pattern recognition and indicates the need to improve scalability.

---

> ❌ **Common Mistake:** Thinking every problem is completely new.
> 
> Most programming problems have already been solved in some form. Learn to ask: *"What does this remind me of?"*

---

### ✅ Builder Checkpoint

Before moving ahead, confirm you can:
- [ ] Explain Computational Thinking in your own words
- [ ] Understand why programmers think before coding
- [ ] Explain decomposition
- [ ] Identify patterns in simple examples

**If all checked,** you're building the mindset of an AI Builder!

---

### Fast Review — Pillars 1 & 2

- Computational Thinking is a way of solving problems logically
- Programming begins with thinking, not typing
- **Decomposition** means breaking big problems into smaller ones
- **Pattern Recognition** means finding similarities between problems
- Experienced developers rely heavily on patterns instead of reinventing solutions
- AI systems also recognize patterns, though in much more complex ways

---

---

# Pillar Three: Abstraction

## What is Abstraction?

**Scenario:** Your friend calls and says *"Come to my house."*

Do they explain every road, tree, electric pole, shop, and person you'll pass? No. They say:
> "Turn left at the mosque, then right at the supermarket."

**Why?** They tell you **only the information that matters.** Everything else is unnecessary.

**That's Abstraction.**

---

## Simple Definition

**Abstraction** = Focusing on important details while ignoring unnecessary ones.

Think of it as **filtering**—keep what helps solve the problem, ignore what doesn't.

---

## Why is Abstraction Important?

**Your Laptop:** When you press the power button, the computer performs millions of operations:
- Checks hardware
- Loads OS
- Initializes memory
- Starts drivers
- Launches services

Do you need to know all this to turn it on? **No.** You simply press the Power Button.

**The complicated work is hidden. That's abstraction.**

---

## Real-Life Examples

| **Activity** | **You Use** | **You Don't Need to Know** |
|:---|:---|:---|
| **Driving a Car** | Steering wheel, Brake, Accelerator, Gear | Fuel injection, Spark plugs, Engine timing, Transmission |
| **TV Remote** | Volume buttons, Channel buttons | Electrical signals, Infrared, Internal circuits |
| **ATM Machine** | Insert card → Enter PIN → Withdraw | Communicates with bank, Security systems, Databases |

---

## Programming Examples

**Using print():**
```python
print("Hello World")
```

Behind the scenes, Python: finds the function, converts text to bytes, sends to terminal, displays correctly.

**You don't manage these steps.** Python hides them. **That's abstraction.**

**Using sort():**
```python
numbers = [4, 8, 1, 9]
numbers.sort()
```

You don't tell Python "Compare, swap, repeat." You simply call `sort()`. **Python handles the details.**

---

## Hospital Management System

**System breaks into:**
```
Patients → Doctors → Appointments → Billing → Medical Records → Pharmacy
```

Each department focuses **only on what it needs.**
- Pharmacy doesn't know how Billing calculates taxes
- Billing doesn't know how Appointments schedules patients
- Each part hides its internal complexity from others

**That's abstraction.**

---

## Abstraction in AI

**With ChatGPT:** You type `Explain recursion.` and get an answer.

Do you know **how billions of parameters interact?** Which mathematical operations occur? Which GPU performed calculations? **No, and you don't need to.**

You only care about the result. **The complexity is hidden.**

---

## Abstraction vs Decomposition

These are often confused. Let's compare:

| **Aspect** | **Decomposition** | **Abstraction** |
|:---|:---|:---|
| **Question** | How can I break this into pieces? | Which details matter? |
| **Example** | Build E-commerce → Login, Cart, Payment, Orders | Building Login: Keep (Username, Password, Validation), Ignore (DB optimization, Server cooling) |
| **Action** | Divide the problem | Keep relevant, ignore unnecessary |

---

### Memory Aid

- **Decomposition:** "Separate cooking into prep, frying, boiling, serving"
- **Abstraction:** "Right now, only focus on preparing ingredients"

**Different ideas.**

---

> ❌ **Beginner:** "How does every single thing work?"
> 
> ✅ **Experienced:** "What do I need to know right now to solve this?"
> 
> **That's abstraction.**

---

## Assessment Questions

| **Q1** | Abstraction means: |
|:---|:---|
| A | Breaking a problem into smaller pieces. |
| B | Ignoring every detail. |
| C | **Focusing on relevant information while hiding unnecessary complexity.** |
| D | Writing shorter programs. |
| **Answer** | ✅ **C** |

| **Q2** | Best example of abstraction? |
|:---|:---|
| A | Learning every transistor before writing Python. |
| B | **Using `sort()` without knowing its internal implementation.** |
| C | Writing the same code repeatedly. |
| D | Deleting unnecessary files. |
| **Answer** | ✅ **B** |

| **Q3** | NOT necessary for an online shopping cart? |
|:---|:---|
| A | Product Price |
| B | Product Quantity |
| C | Customer's Shipping Address |
| D | **Customer's Favorite Musician** |
| **Answer** | ✅ **D** |

| **Q4** | TRUE about abstraction? |
|:---|:---|
| A | **It hides unnecessary complexity.** |
| B | It means avoiding programming. |
| C | It means ignoring every detail. |
| D | It replaces decomposition. |
| **Answer** | ✅ **A** |

---

## Interactive Exercises

### Food Delivery App

**Available information:** Customer's address, Restaurant name, Driver location, Favorite movie, Food order, Payment status, Driver's vehicle, Primary school

**Question:** Which should be ignored?

**Answer:** Favorite movie, Primary school | **Reason:** Don't help deliver the food.

---

### Result Checking System (University)

**Available:** Student ID, Student Name, Exam Scores, Blood Group, Favorite Color, Registered Courses

**Question:** What's unnecessary for checking results?

**Answer:** Blood Group, Favorite Color | **Reason:** Don't contribute to results.

---

### Library System

**Available:** Book Title, Author, ISBN, Shelf Number, Borrower's Name, Shoe Size

**Question:** What should be ignored?

**Answer:** Borrower's Shoe Size | **Reason:** No relationship to borrowing books.

---

## Real Programming Scenario

**Manager says:** *"Create a login page."*

**Beginner worries about:** AI, Database scaling, Cloud servers, Cybersecurity, Docker, Kubernetes

**Experienced developer asks:** *"What do I need to finish the login page?"*

**Usually only:** Username field, Password field, Validation, Submit button, Authentication

**Everything else can wait. That's abstraction.**

---

> 💡 **Builder Tips:** When solving problems, ask:
> 1. What information is essential?
> 2. What information can wait?
> 3. Am I solving today's problem, or worrying about tomorrow's?
>
> **These three questions save countless hours.**

---

> ❌ **Common Mistakes:**
> - Trying to understand every detail before starting
> - Overcomplicating simple problems
> - Collecting information that isn't needed
> - Solving problems that don't exist yet

---

### ✅ Builder Checkpoint — Pillar 3

Before Algorithm Design, confirm:
- [ ] I understand what abstraction is
- [ ] I know how it differs from decomposition
- [ ] I can identify unnecessary information in a problem
- [ ] I understand why languages provide abstractions like `print()` and `sort()`
- [ ] I can apply abstraction when designing software

**If all checked,** you've mastered the third pillar!

---

### Fast Review — Pillar 3: Abstraction

- Abstraction means focusing on what matters and hiding unnecessary details
- It reduces complexity and helps programmers solve problems more efficiently
- Real-world examples: driving a car, using an ATM, a TV remote, using ChatGPT
- In programming, functions like `print()` and `sort()` are abstractions that hide complex internal operations
- **Decomposition** breaks a problem into smaller parts; **Abstraction** decides which details are important
- Skilled programmers don't try to understand everything at once—they focus on what's necessary for the current problem

---

---

# Pillar Four: Algorithm Design

## What is an Algorithm?

**Question:** How do you make a cup of tea?

**Answer:**
1. Boil water
2. Put tea bag in cup
3. Pour hot water
4. Let it steep
5. Add sugar/milk (optional)
6. Stir
7. Drink

Notice: You didn't just say *what* to do—you said **the exact order** in which to do it.

**That ordered list is called an algorithm.**

---

## Simple Definition

**Algorithm** = A step-by-step procedure for solving a problem or completing a task.

Think of it as a **recipe**—it tells a computer how to solve a problem, just as it tells a chef how to prepare food.

---

## Why Are Algorithms Important?

**Scenario:** Travel from Kano to Abuja.

Would you randomly choose roads, or follow the fastest route? **Obviously, the best route.**

**That route is an algorithm.**

Algorithms help us solve problems:
- **Faster** — Efficient solutions save time
- **More accurately** — Precise steps reduce errors
- **More consistently** — Same input, same output

Without algorithms, **programming would be guessing.**

---

## Everyday Algorithms

**Brushing Teeth:** Pick toothbrush → Apply toothpaste → Brush teeth → Rinse mouth → Wash toothbrush

Change the order? The result may not be correct.

**Logging into Email:** Open Browser → Visit Website → Enter Email → Enter Password → Click Login → Access Account

Each step depends on the previous one.

**Cooking Rice:** Wash rice → Boil water → Add rice → Cook → Serve

(A sequence of steps)

---

## Algorithms in Programming

**Task:** Find the largest number in a list: `[5, 9, 3, 12, 7]`

**What a computer needs:**
```
Start → Assume first is largest → Compare next → If larger, update → Repeat → Output largest → Stop
```

**Code Implementation:**
```python
numbers = [5, 9, 3, 12, 7]
largest = numbers[0]
for number in numbers:
    if number > largest:
        largest = number
print(largest)  # Output: 12
```

**Key insight:** Before writing code, someone first designed the algorithm.

---

## Characteristics of a Good Algorithm

| **Characteristic** | **Description** | **Example** |
|:---|:---|:---|
| **Clear** | Every instruction is easy to understand | ✅ "Add 10 to total" vs ❌ "Do something" |
| **Finite** | Algorithm must eventually stop | ✅ "Repeat until all checked" vs ❌ "Repeat forever" |
| **Correct** | Solves the intended problem | ✅ ATM gives correct amount vs ❌ ATM gives wrong amount |
| **Efficient** | Saves time, memory, computing power | ✅ 1 second vs ❌ 1 minute for same task |

---

## Algorithm Representation

### Flowcharts

Visual representation:
```
Start → Input Age → Age ≥ 18? → Yes → Adult → End
                          ↓
                         No → Minor
```

Flowcharts help programmers understand logic before coding.

---

### Pseudocode

Combines English and logic without syntax concerns:
```
Read age
If age ≥ 18
    Display "Adult"
Else
    Display "Minor"
```

Focuses on **logic, not programming syntax.**

---

## Algorithm vs Program

| **Algorithm** | **Program** |
|:---|:---|
| A **plan** | **Actual code** |
| Make tea → Boil water → Add tea → Serve | `if water_is_hot: make_tea()` |
| **Blueprint** | **Building** |

---

## Professional Algorithm Design Process

```
1. Understand the problem
2. Identify inputs & outputs
3. Break into smaller steps
4. Write the algorithm
5. Test the algorithm
6. Convert to code
7. Debug if necessary
```

**Skipping the algorithm leads to confusing code.**

---

## Interactive Exercises

### Calculator Algorithm

**Task:** Design an algorithm for a calculator before coding.

**Answer:**
```
Start → Input first number → Input second number → Choose operation 
     → Perform calculation → Display answer → Stop
```

---

### Online Payment Algorithm

**Answer:**
```
Start → Enter Card Details → Verify Card → Enter Amount → Confirm Payment
     → Process Transaction → Display Result → Stop
```

---

### Robot Navigation

**Question:** A robot can only move Forward, Left, Right (no jumps). How do you instruct it?

**Answer:** Create a sequence using only allowed movements. **This demonstrates algorithm design.**

---

### Measuring Liquid

**Question:** You have buckets of 10L, 5L, and 2L. Measure exactly 7 litres.

**Hint:** Think step by step. **The lesson is breaking problems into ordered steps,** not just finding the answer.

---

## Assessment Questions

| **Q1** | Calculate student average. What comes FIRST? |
|:---|:---|
| A | Print the average. |
| B | **Read the students' scores.** |
| C | Delete the data. |
| D | Restart the computer. |
| **Answer** | ✅ **B** — Can't calculate average without data |

| **Q2** | Best definition of an algorithm? |
|:---|:---|
| A | A programming language. |
| B | A database. |
| C | **A step-by-step solution to a problem.** |
| D | A computer monitor. |
| **Answer** | ✅ **C** |

---

## Mini Quiz — Algorithm Design

| **Q1** | An algorithm is: |
|:---|:---|
| A | A programming language. |
| B | A computer. |
| C | **A step-by-step method for solving a problem.** |
| D | A database. |
| **Answer** | ✅ **C** |

| **Q2** | NOT a good algorithm characteristic? |
|:---|:---|
| A | Clear |
| B | Finite |
| C | Correct |
| D | **Random** |
| **Answer** | ✅ **D** |

| **Q3** | Why write algorithms before coding? |
|:---|:---|
| A | Make the code look longer. |
| B | **Think through solution before implementing.** |
| C | Avoid using computers. |
| D | Replace programming. |
| **Answer** | ✅ **B** |

| **Q4** | Which is pseudocode? |
|:---|:---|
| A | `print("Hello")` |
| B | **Read numbers → Add → Display result** |
| C | `console.log("Hello");` |
| D | `fmt.Println("Hello")` |
| **Answer** | ✅ **B** |

| **Q5** | First when solving a problem? |
|:---|:---|
| A | Choose a font. |
| B | **Understand the problem.** |
| C | Deploy the application. |
| D | Publish the software. |
| **Answer** | ✅ **B** |

---

> 💡 **Builder Tip:** When asked to build software:
> 
> ❌ Don't ask: "Which programming language should I use?"  
> ✅ Instead ask: "What is the algorithm?"
>
> **Good programmers think solution-first.** The language is just a tool for expressing that solution.

---

### 🏆 Builder Challenge — ATM Algorithm

**Task:** Design an algorithm for ATM cash withdrawal (no code).

**Include:**
- Card insertion
- PIN verification
- Balance checking
- Amount entry
- Cash dispensing
- Receipt (optional)
- Card ejection

Try yourself before looking at solutions.

---

> ❌ **Common Algorithm Mistakes:**
> - Starting to code before understanding the problem
> - Writing code without planning
> - Ignoring edge cases
> - Assuming the computer "knows what you mean"
>
> **Remember:** Computers only do exactly what you tell them.

---

### ✅ Builder Checkpoint — Pillar 4

Before moving on, confirm:
- [ ] I know what an algorithm is
- [ ] I can explain the difference between an algorithm and a program
- [ ] I understand why programmers design algorithms before coding
- [ ] I can write simple pseudocode
- [ ] I can draw a basic flowchart
- [ ] I can design a simple algorithm for an everyday task

**If all checked,** you've mastered algorithm design!

---

### Fast Review — Pillar 4: Algorithm Design

- An **algorithm** is a step-by-step procedure for solving a problem
- Algorithms are like recipes—they describe *how* to complete a task
- Good algorithms are **clear, finite, correct, and efficient**
- **Flowcharts** provide a visual representation of algorithms
- **Pseudocode** helps you focus on logic before worrying about syntax
- Professional developers understand the problem and design an algorithm before coding
- **Programming** is the process of translating an algorithm into a programming language

---

---

# Module 1 Summary

## The Four Pillars of Computational Thinking

| **Pillar** | **Definition** | **Example** |
|:---|:---|:---|
| **1. Decomposition** | Break a large problem into smaller, manageable pieces | Building Facebook: Login, Registration, Profiles, Messaging, Notifications, Friends, News Feed |
| **2. Pattern Recognition** | Find similarities with problems you've solved before | Recognizing login patterns across multiple applications |
| **3. Abstraction** | Focus on important information, ignore unnecessary details | Using `sort()` without knowing how it works internally |
| **4. Algorithm Design** | Create clear, step-by-step solutions before coding | Designing ATM withdrawal steps before coding |

---

## The Computational Thinking Workflow

```
┌─────────────────┐
│   Big Problem   │
└────────┬────────┘
         ↓
┌──────────────────────┐
│   Decomposition      │ (break it down)
│  (break it down)     │
└────────┬─────────────┘
         ↓
┌──────────────────────┐
│ Pattern Recognition  │ (find patterns)
│   (find patterns)    │
└────────┬─────────────┘
         ↓
┌──────────────────────┐
│    Abstraction       │ (what matters?)
│   (what matters?)    │
└────────┬─────────────┘
         ↓
┌──────────────────────┐
│  Algorithm Design    │ (step-by-step)
│ (step-by-step)       │
└────────┬─────────────┘
         ↓
┌──────────────────────┐
│       Code           │ (implementation)
│  (implementation)    │
└──────────────────────┘
```

---

## 🎯 Key Takeaway

**Mastering these four pillars transforms you from someone who merely writes code into someone who can design solutions.**

**That's exactly the mindset expected of an AI Builder.**

---

### Congratulations! 🎉

You've completed Module 1 and learned:
- ✅ What Computational Thinking is
- ✅ Why it matters for programmers
- ✅ The four pillars and how to use them
- ✅ How to approach problems like an AI Builder

You're now ready for the next module. Keep practicing these skills!
