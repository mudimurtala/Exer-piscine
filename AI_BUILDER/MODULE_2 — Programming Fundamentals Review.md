# MODULE 2 — Programming Fundamentals Review

This module is organized into four parts so the review stays structured and easy to scan.

## Module Map

| Part   | Focus                          | Outcome                                                                                      |
| ------ | ------------------------------ | -------------------------------------------------------------------------------------------- |
| Part 1 | Core Programming Foundations   | Understand variables, data types, operators, expressions, input/output, and type conversion. |
| Part 2 | Control Flow                   | Use decisions, loops, loop control, and basic debugging logic.                               |
| Part 3 | Functions and Data Collections | Work with functions, scope, lists, strings, dictionaries, and sets.                          |
| Part 4 | Programming Like an AI Builder | Read code, find bugs, refactor, handle errors, and think critically about AI-generated code. |

---

# AI Builder Assessment Bootcamp

## MODULE 2

# Programming Fundamentals Review

---

# Part 1 — Core Programming Foundations

**Estimated Study Time:** 3–4 Hours  
**Difficulty:** ⭐⭐☆☆☆ (Beginner → Intermediate)

## Learning Objectives

By the end of this part, you should be able to:

- Explain what variables are and why they exist.
- Distinguish between common data types.
- Use operators correctly.
- Understand expressions.
- Explain how programs receive input and produce output.
- Understand type conversion.
- Read simple code with confidence.

## Chapter 1 — How Computers Think

A computer is incredibly fast, but it is **not intelligent**. It does exactly what you tell it to do, nothing more and nothing less.

If you say, “Go to the shop,” a human may ask follow-up questions. A computer never does that. It waits for precise instructions.

> Programming is simply giving complete and precise instructions.

### Real-Life Analogy

If you teach a robot to make tea, you cannot just say “Make tea.” You must break the task into exact steps.

```
Boil water
↓
Put tea bag into cup
↓
Pour hot water
↓
Wait 3 minutes
↓
Serve
```

The robot succeeds because the instructions are specific.

## Chapter 2 — Variables

### What is a Variable?

Think of containers labeled rice, sugar, and salt. Each container stores something different, and each has a name.

A variable is a **named location that stores data**.

### Simple Definition

A variable is a named place in memory used to store information.

### Visual Illustration

```
age
↓
+------+
| 20   |
+------+

name
↓
+-----------+
| "Mudi"    |
+-----------+
```

### Why Variables Matter

Without variables, changing a value means editing every occurrence manually. With variables, you update one value in one place.

### Examples

Python

```python
age = 20
print(age)
```

Output

```text
20
```

JavaScript

```javascript
let age = 20;
console.log(age);
```

Go

```go
package main

import "fmt"

func main() {
    age := 20
    fmt.Println(age)
}
```

### Naming Variables

Good variable names are descriptive:

```text
student_name
total_score
price
isLoggedIn
```

Avoid vague names:

```text
a
abc
data
x1
```

> Write code for humans first, computers second.

## Chapter 3 — Data Types

Variables store different kinds of information. Those kinds are called **data types**.

### Common Data Types

| Type    | Meaning                     | Examples                            |
| ------- | --------------------------- | ----------------------------------- |
| Integer | Whole numbers               | `5`, `20`, `100`, `-7`              |
| Float   | Numbers with decimal places | `3.14`, `2.5`, `99.99`              |
| String  | Text                        | `"Hello"`, `"Mudi"`, `"AI Builder"` |
| Boolean | Two possible values         | `true`, `false`                     |

### Visual Illustration

```
Integer
↓
42

Float
↓
3.14

String
↓
"Hello"

Boolean
↓
True
```

### Why Data Types Matter

Computers need to know whether they are working with text, whole numbers, decimals, or true/false values. For example, `"Mudi" + 25` does not make sense as a calculation.

### Practice

Identify the data type:

| Value      | Answer  |
| ---------- | ------- |
| `25`       | Integer |
| `3.14159`  | Float   |
| `"Python"` | String  |
| `False`    | Boolean |

## Chapter 4 — Operators

Operators perform actions on data.

### Arithmetic Operators

| Operator | Meaning            |
| -------- | ------------------ |
| `+`      | Addition           |
| `-`      | Subtraction        |
| `*`      | Multiplication     |
| `/`      | Division           |
| `%`      | Remainder (Modulo) |

Examples:

```python
10 + 5
```

```text
15
```

```python
20 / 4
```

```text
5
```

```python
17 % 5
```

```text
2
```

Because `17 = 5 × 3 + 2`, the remainder is `2`.

### Comparison Operators

| Operator | Meaning               |
| -------- | --------------------- |
| `>`      | Greater than          |
| `<`      | Less than             |
| `>=`     | Greater than or equal |
| `<=`     | Less than or equal    |
| `==`     | Equal                 |
| `!=`     | Not equal             |

Examples:

```python
10 > 5
```

```text
True
```

```python
10 == 8
```

```text
False
```

### Logical Operators

| Operator | Meaning                             |
| -------- | ----------------------------------- |
| `AND`    | Both conditions must be true        |
| `OR`     | At least one condition must be true |
| `NOT`    | Reverses a condition                |

Example: a student passes only if score ≥ 50 and attendance ≥ 75%.

## Chapter 5 — Expressions

An expression is anything that produces a value.

Examples:

```python
5 + 3
```

```text
8
```

```python
age > 18
```

```text
True
```

## Chapter 6 — Input and Output

Programs receive information and produce information. That is input and output.

### Everyday Example

ATM

```
Input
↓
Enter PIN
↓
Computer processes
↓
Output
↓
Cash
```

### Python Example

```python
name = input("Enter your name: ")
print(name)
```

If the user types `Mudi`, the output is:

```text
Mudi
```

## Chapter 7 — Type Conversion

Sometimes a computer receives numbers as text, such as `"25"`. To calculate with that value, you must convert it.

### Python Example

```python
age = int("25")
print(age + 5)
```

```text
30
```

Without conversion:

```python
"25" + "5"
```

```text
255
```

With conversion:

```python
25 + 5
```

```text
30
```

## Common Beginner Mistakes

- Confusing `=` with `==`.
- Mixing strings with numbers.
- Giving variables meaningless names.
- Forgetting that programming is case-sensitive.

`Age`, `age`, and `AGE` can be different names in many languages.

## Hands-on Exercises

1. Create variables for your name, age, country, and whether you enjoy programming. Print all four.
2. Create two variables: `a = 15` and `b = 10`. Calculate the sum, difference, product, and quotient.
3. Ask the user for their age and display: `Next year you will be ...`

## Logical Reasoning Exercise

What is wrong with this code?

```python
age = "20"
print(age + 5)
```

Answer: `age` is a string while `5` is an integer. Convert the value first.

```python
age = int("20")
print(age + 5)
```

## Assessment Questions

1. What is the purpose of a variable?
   - A. To display pictures
   - B. To store information
   - C. To connect to the Internet
   - D. To create folders
   - Answer: ✅ B

2. Which of the following is a Boolean value?
   - A. `42`
   - B. `"Hello"`
   - C. `True`
   - D. `3.14`
   - Answer: ✅ C

3. Which operator checks whether two values are equal?
   - A. `=`
   - B. `==`
   - C. `+`
   - D. `%`
   - Answer: ✅ B

4. What is the result of `17 % 5`?
   - A. `3`
   - B. `2`
   - C. `5`
   - D. `17`
   - Answer: ✅ B

5. Which statement about variables is true?
   - A. Variable names should be meaningful
   - B. Every variable must contain text
   - C. Variables never change
   - D. Variables can only store numbers
   - Answer: ✅ A

## Builder Checkpoint

Before moving to Part 2, make sure you can confidently say:

- I understand how computers follow instructions.
- I know what variables are and why they are useful.
- I can identify common data types.
- I understand arithmetic, comparison, and logical operators.
- I know what an expression is.
- I understand input and output.
- I can explain why type conversion is sometimes necessary.

## Fast Review

- A computer follows instructions exactly as written.
- Variables are named containers for storing data.
- Common data types include Integer, Float, String, and Boolean.
- Operators perform arithmetic, comparison, and logical evaluation.
- Expressions produce values.
- Programs receive input, process it, and produce output.
- Type conversion changes data from one type to another.
- Good variable names make code easier to read and maintain.

---

# Part 2 — Control Flow

**Estimated Study Time:** 4–5 Hours  
**Difficulty:** ⭐⭐☆☆☆ (Beginner → Intermediate)

## Learning Objectives

By the end of this part, you should be able to:

- Make decisions using conditional statements.
- Repeat tasks efficiently using loops.
- Understand nested loops.
- Know when to use `break` and `continue`.
- Recognize common programming patterns.
- Debug simple logical errors.

## Chapter 1 — What is Control Flow?

Control flow is the order in which a program executes its instructions.

Think about driving to work:

- If the traffic light is green, continue.
- If it is red, stop.
- If there is traffic, take another route.
- If you reach the office, park.

The next action depends on what happens. Programs work the same way.

> Control flow simply means the order in which a program runs instructions.

## Chapter 2 — Conditional Statements

Programs often need to make decisions.

### Everyday Example

```
Is it raining?
↓
Yes → Carry an umbrella
No  → Leave without one
```

### Python Example

```python
age = 20

if age >= 18:
    print("Adult")
```

### JavaScript Example

```javascript
let age = 20;

if (age >= 18) {
  console.log("Adult");
}
```

### Go Example

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

### if...else

```python
age = 15

if age >= 18:
    print("Adult")
else:
    print("Minor")
```

### Multiple Conditions

Example:

```
Score ≥ 70 → Excellent
Score ≥ 50 → Pass
Otherwise  → Fail
```

```python
score = 65

if score >= 70:
    print("Excellent")
elif score >= 50:
    print("Pass")
else:
    print("Fail")
```

## Chapter 3 — Switch Statements

Some languages use `switch` to compare one value against several possibilities.

```go
switch day {
case "Monday":
    fmt.Println("Work")
case "Saturday":
    fmt.Println("Rest")
default:
    fmt.Println("Unknown")
}
```

## Chapter 4 — Loops

Loops repeat instructions. Instead of writing code many times, you write it once.

### Why Loops Matter

Whether you have 10 students, 100 students, or 10,000 students, loops keep the code manageable.

### The `for` Loop

```python
for i in range(5):
    print(i)
```

```text
0
1
2
3
4
```

Another example:

```python
for i in range(3):
    print("AI Builder")
```

### Understanding Loop Variables

```python
for number in [5, 8, 10]:
    print(number)
```

Iterations:

- Iteration 1: `number = 5`
- Iteration 2: `number = 8`
- Iteration 3: `number = 10`

### The `while` Loop

A `while` loop repeats until a condition becomes false.

```python
count = 1
while count <= 5:
    print(count)
    count += 1
```

### Difference Between `for` and `while`

- Use `for` when you know how many times to repeat.
- Use `while` when the stopping point depends on a condition.

## Chapter 5 — Nested Loops

A nested loop is a loop inside another loop.

```python
for row in range(3):
    for col in range(2):
        print(row, col)
```

Think of a calendar: one loop for months, another for days.

## Chapter 6 — `break`

`break` stops a loop early.

```python
for number in range(10):
    if number == 5:
        break
    print(number)
```

### Real-Life Example

If you search for a friend in a crowd and find them, you stop searching. That is `break`.

## Chapter 7 — `continue`

`continue` skips only the current iteration.

```python
for number in range(6):
    if number == 3:
        continue
    print(number)
```

### Real-Life Example

If one student is absent during attendance, you skip that student and continue with the rest.

## `break` vs `continue`

| Keyword    | Effect                                      |
| ---------- | ------------------------------------------- |
| `break`    | Stops the entire loop                       |
| `continue` | Skips the current iteration and keeps going |

## Chapter 8 — Common Programming Patterns

### Pattern 1 — Counting

```python
count = 0
for number in numbers:
    count += 1
```

### Pattern 2 — Summing

```python
total = 0
for number in numbers:
    total += number
```

### Pattern 3 — Finding the Largest

```python
largest = numbers[0]
for number in numbers:
    if number > largest:
        largest = number
```

### Pattern 4 — Searching

```python
found = False
for number in numbers:
    if number == target:
        found = True
```

## Chapter 9 — Debugging Logic

Many bugs are logic errors, not syntax errors.

Example:

```python
age = 16

if age > 18:
    print("Adult")
```

If age is exactly 18, nothing prints. The intended condition may be `>=`.

Another example:

```python
total = 1
for number in numbers:
    total += number
```

If you want a sum, should `total` start at `0` or `1`?

### Builder Tips

When code gives the wrong result, ask:

1. Which line produced the wrong value?
2. Was the condition correct?
3. Did the loop run the expected number of times?
4. Did the variables change correctly?

## Hands-on Exercises

1. Write a program that determines whether someone can vote. Requirement: age ≥ 18. Output `Eligible` or `Not Eligible`.
2. Print all numbers from 1 to 20.
3. Print only the even numbers between 1 and 20. Hint: use `%`.
4. Calculate the sum of 1 through 100.
5. Find the largest number in `[12, 7, 35, 4, 18]`.

## Logical Reasoning Exercise

Predict the output:

```python
for i in range(4):
    if i == 2:
        continue
    print(i)
```

Answer:

```text
0
1
3
```

## Assessment Questions

1. What is the purpose of an `if` statement?
   - A. Repeat code
   - B. Store data
   - C. Make decisions
   - D. Create variables
   - Answer: ✅ C

2. Which loop is generally used when you know how many times to repeat?
   - A. while
   - B. for
   - C. switch
   - D. break
   - Answer: ✅ B

3. What does `break` do?
   - A. Skips one iteration
   - B. Ends the entire loop
   - C. Restarts the loop
   - D. Creates a variable
   - Answer: ✅ B

4. What does `continue` do?
   - A. Ends the loop
   - B. Skips only the current iteration
   - C. Repeats forever
   - D. Creates a function
   - Answer: ✅ B

5. What is a nested loop?
   - A. Two variables
   - B. A loop inside another loop
   - C. A function inside a function
   - D. Two programs running together
   - Answer: ✅ B

## Builder Checkpoint

Before moving to Part 3, make sure you can confidently say:

- I understand how control flow affects program execution.
- I can write `if`, `else`, and `elif` statements.
- I know when to use `for` and `while` loops.
- I understand the difference between `break` and `continue`.
- I recognize common looping patterns like counting, summing, searching, and finding the largest value.
- I know how to begin debugging simple logic errors.

## Fast Review

- Control flow determines the order in which a program executes.
- Conditional statements allow programs to make decisions.
- `if`, `else`, and `elif` handle different conditions.
- `switch` provides another way to choose between multiple cases in some languages.
- `for` loops repeat a known number of times.
- `while` loops repeat until a condition becomes false.
- Nested loops are loops inside other loops.
- `break` stops a loop immediately.
- `continue` skips the current iteration and continues with the next one.
- Common tasks include counting, summing, searching, and finding the largest value.
- Debugging starts by understanding the logic, not by guessing.

---

# Part 3 — Functions and Data Collections

**Estimated Study Time:** 5–6 Hours  
**Difficulty:** ⭐⭐⭐☆☆ (Beginner → Intermediate)

## Learning Objectives

By the end of this part, you should be able to:

- Explain what a function is and why functions are important.
- Understand parameters and arguments.
- Understand return values.
- Explain variable scope.
- Work confidently with arrays/lists.
- Manipulate strings.
- Understand maps (dictionaries).
- Understand sets.
- Recognize which data collection is appropriate for different situations.

## Chapter 1 — Functions

If you do the same morning routine every day, it would be useful to have one reusable command for it.

> A function is a reusable block of code that performs a specific task.

Instead of repeating code, write it once and call it whenever needed.

### Why Functions Matter

Large programs become messy without functions. Real applications like messaging systems separate responsibilities into different functions such as sending messages, deleting messages, uploading images, logging in, and logging out.

### Python Example

```python
def greet():
    print("Welcome to AI Builder!")

greet()
```

### JavaScript Example

```javascript
function greet() {
  console.log("Welcome!");
}

greet();
```

### Go Example

```go
func greet() {
    fmt.Println("Welcome!")
}

greet()
```

### Real-Life Analogy

A microwave is a good analogy: you press `Start`, it performs its task, then returns control to you.

## Chapter 2 — Parameters and Arguments

If your friend says “Greet Mudi,” then tomorrow “Greet Amina,” you do not need a new function each time. You pass the name into the same function.

### Definition

- Parameter: the variable inside the function.
- Argument: the actual value provided in the function call.

### Python Example

```python
def greet(name):
    print("Hello", name)

greet("Mudi")
greet("Amina")
```

### Easy Trick

| Term      | Meaning      |
| --------- | ------------ |
| Parameter | Placeholder  |
| Argument  | Actual value |

## Chapter 3 — Return Values

Not every function prints something. Many functions calculate a result and return it.

```python
def add(a, b):
    return a + b

answer = add(5, 3)
print(answer)
```

```python
def square(number):
    return number * number
```

Returning a value is like asking an ATM for your balance: the machine sends the answer back.

## Chapter 4 — Variable Scope

Variables have scope, which means they are only accessible in certain places.

```python
def greet():
    message = "Hello"
```

`message` only exists inside `greet()`.

Global variables can be accessed in many places:

```python
name = "Mudi"

def show():
    print(name)
```

> Use local variables whenever possible. Global variables can make programs harder to understand.

## Chapter 5 — Arrays / Lists

A list stores multiple values in one place.

### Python Example

```python
students = [
    "Mudi",
    "Amina",
    "John"
]
```

Accessing an item:

```python
print(students[0])
```

Adding an item:

```python
students.append("Fatima")
```

Looping through a list:

```python
for student in students:
    print(student)
```

Real-life example: a shopping list.

## Chapter 6 — Strings

Strings are text.

```python
name = "Mudi"
```

Common string operations:

| Operation | Example                | Result       |
| --------- | ---------------------- | ------------ |
| Length    | `len(name)`            | `4`          |
| Uppercase | `name.upper()`         | `MUDI`       |
| Lowercase | `name.lower()`         | `mudi`       |
| Contains  | `"Mu" in name`         | `True`       |
| Join      | `first + " " + second` | `AI Builder` |

## Chapter 7 — Maps (Dictionaries)

Dictionaries store pairs of related values.

```python
student = {
    "name": "Mudi",
    "score": 95
}

print(student["score"])
```

Think of a dictionary as `Key → Value`.

## Chapter 8 — Sets

Sets do not allow duplicates.

```python
numbers = {
    1,
    2,
    2,
    3,
    3,
    3
}
```

The result is unique values only.

## Choosing the Right Collection

| Use        | When                                     |
| ---------- | ---------------------------------------- |
| List       | Order matters and duplicates are allowed |
| Dictionary | You need `Key → Value` relationships     |
| Set        | Duplicates are not allowed               |

### Comparison Table

| Collection | Ordered           | Duplicates             | Access by Key |
| ---------- | ----------------- | ---------------------- | ------------- |
| List       | ✅ Yes            | ✅ Yes                 | ❌ No         |
| Dictionary | ✅ Keys           | ❌ Keys must be unique | ✅ Yes        |
| Set        | Usually unordered | ❌ No                  | ❌ No         |

## Hands-on Exercises

1. Write a function called `multiply()` that receives two numbers and returns their product.
2. Create a list containing five programming languages and print each one.
3. Create a dictionary representing yourself with name, age, country, and favorite language. Print each value.
4. Create a set containing `Python`, `Go`, `Python`, `JavaScript`, `Go` and observe what happens.

## Logical Reasoning Exercise

Predict the output:

```python
numbers = [4, 8, 12]
print(numbers[1])
```

Answer:

```text
8
```

Reason: lists begin at index `0`.

## Assessment Questions

1. Why do programmers use functions?
   - A. To make programs slower
   - B. To reuse code
   - C. To delete variables
   - D. To replace loops
   - Answer: ✅ B

2. What is a parameter?
   - A. The value passed into a function call
   - B. The placeholder variable defined in a function
   - C. A programming language
   - D. A loop
   - Answer: ✅ B

3. Which keyword sends a value back from a function?
   - A. stop
   - B. break
   - C. return
   - D. exit
   - Answer: ✅ C

4. Which collection stores data as Key → Value pairs?
   - A. List
   - B. Set
   - C. Dictionary
   - D. Tuple
   - Answer: ✅ C

5. Which collection automatically removes duplicates?
   - A. List
   - B. Dictionary
   - C. Set
   - D. String
   - Answer: ✅ C

6. What is printed?

   ```python
   items = ["Pen", "Book", "Bag"]
   print(items[2])
   ```

   - A. Pen
   - B. Book
   - C. Bag
   - D. Error
   - Answer: ✅ C

## Common Beginner Mistakes

- Forgetting to call a function after defining it.
- Confusing parameters with arguments.
- Printing instead of returning values.
- Trying to access a local variable outside its function.
- Forgetting that list indexing starts at `0`.
- Using a list when a dictionary would make the solution simpler.
- Expecting a set to keep duplicate values.

## Builder Tips

Ask yourself:

- Can I turn repeated code into a function?
- Am I using the right data structure for this problem?
- Would a dictionary make this lookup easier?
- Do I really need duplicates, or would a set be better?

## Builder Checkpoint

Before moving to Part 4, make sure you can confidently say:

- I understand what functions are and why they are useful.
- I know the difference between parameters and arguments.
- I understand how return values work.
- I know the difference between local and global variables.
- I can work with lists.
- I can manipulate strings.
- I understand dictionaries (maps).
- I know when to use sets.

## Fast Review

- A function is a reusable block of code.
- Parameters are placeholders; arguments are the actual values passed to a function.
- `return` sends a value back to the caller.
- Scope determines where a variable can be accessed.
- Lists store ordered collections of items.
- Strings represent text and provide useful operations.
- Dictionaries store data as key → value pairs.
- Sets store unique values and automatically remove duplicates.
- Choosing the correct data collection makes programs simpler, faster, and easier to maintain.

---

# Part 4 — Programming Like an AI Builder

**Estimated Study Time:** 5–6 Hours  
**Difficulty:** ⭐⭐⭐☆☆ (Intermediate)

## Learning Objectives

By the end of this chapter, you should be able to:

- Read code you did not write.
- Find and fix common bugs.
- Write cleaner and more maintainable code.
- Refactor simple programs.
- Understand common programming errors.
- Debug programs methodically.
- Solve programming assessment questions with confidence.

## Chapter 1 — Reading Unfamiliar Code

A beginner asks, “How do I write code?” An experienced developer asks, “How do I understand existing code?” In practice, reading code is a major part of programming.

### How to Read Code

Start from the top and follow the structure:

```
Read the function name
↓
Read the parameters
↓
Read variable declarations
↓
Read loops
↓
Read conditions
↓
Read the return statement
↓
Ask: What problem is this code solving?
```

### Example

```python
def multiply(a, b):
    result = a * b
    return result
```

Ask:

- What is the function called? `multiply`
- What information does it receive? Two numbers
- What does it return? Their product

> Read unfamiliar code like a detective: every variable, function, and loop is a clue.

## Chapter 2 — Finding Bugs

A bug is an error in a program. Not all bugs are syntax errors; many are logic errors.

### Syntax Error

```python
if age > 18
    print(age)
```

Problem: missing colon.

### Logic Error

```python
age = 18

if age > 18:
    print("Adult")
```

The program runs, but 18 is not treated as an adult. The condition should probably be `>=`.

### Runtime Error

```python
number = 10
print(number / 0)
```

The program starts and then crashes because dividing by zero is impossible.

### Three Types of Errors

```
Syntax Error  → Program cannot start
Runtime Error → Program crashes while running
Logic Error   → Program runs but produces the wrong answer
```

## Chapter 3 — Debugging

Debugging means finding and fixing errors.

### The Debugging Process

```
Observe the problem
↓
Reproduce it
↓
Locate the error
↓
Understand why it happened
↓
Fix it
↓
Test again
```

### Debugging Example

```python
numbers = [5, 10, 15]
print(numbers[3])
```

There are only three items, so valid indexes are `0`, `1`, and `2`. Index `3` does not exist.

## Chapter 4 — Writing Clean Code

Good code works, and it is also easy for another human to understand later.

### Bad

```python
x = 100
y = 0.15
z = x * y
```

### Good

```python
price = 100
tax_rate = 0.15
tax = price * tax_rate
```

### Clean Code Principles

- Use meaningful names.
- Keep functions short.
- Avoid repeating code.
- Write one function for one responsibility.
- Add comments only when necessary.
- Write code for humans, not just computers.

## Chapter 5 — Refactoring

Refactoring means improving code without changing what it does.

### Bad

```python
print("Hello Mudi")
print("Hello Amina")
print("Hello Musa")
```

### Better

```python
def greet(name):
    print("Hello", name)

greet("Mudi")
greet("Amina")
greet("Musa")
```

### Another Example

```python
subtotal = price + tax + shipping
total = subtotal
final = subtotal
```

Avoid repeating calculations.

## Chapter 6 — Error Handling

Programs should respond gracefully instead of crashing unnecessarily.

```python
try:
    number = int(input())
except ValueError:
    print("Please enter a valid number.")
```

If the user enters the wrong thing, the program explains the problem instead of crashing.

## Chapter 7 — Thinking Like an AI Builder

AI can generate code, but you are still responsible for understanding and testing it.

```python
def average(a, b):
    return a + b / 2
```

The correct version is:

```python
return (a + b) / 2
```

### AI Builder Rules

- Never copy code blindly.
- Understand it.
- Run it.
- Test edge cases.
- Improve it.
- Document it.

## Chapter 8 — Common Assessment Questions

1. What is the first thing you should do after receiving AI-generated code?
   - A. Deploy it immediately
   - B. Delete it
   - C. Read and understand it
   - D. Ignore it
   - Answer: ✅ C

2. Which code is more readable?
   - A. `x = a*b+c`
   - B. `total_price = item_price * quantity + shipping_fee`
   - Answer: ✅ B

3. Which error causes incorrect output even though the program runs?
   - A. Syntax Error
   - B. Runtime Error
   - C. Logic Error
   - D. Compiler Error
   - Answer: ✅ C

## Hands-on Challenges

1. Predict the output:

   ```python
   def square(x):
       return x * x

   print(square(6))
   ```

   Answer: `36`

2. Find the bug:

   ```python
   age = 20
   if age = 20:
       print("Correct")
   ```

   Answer: `==` should be used.

3. Predict the output:

   ```python
   numbers = [2, 4, 6]
   total = 0
   for number in numbers:
       total += number
   print(total)
   ```

   Answer: `12`

4. Refactor this code:

   ```python
   print("Welcome Mudi")
   print("Welcome Musa")
   print("Welcome Amina")
   ```

   Possible answer:

   ```python
   def welcome(name):
       print("Welcome", name)

   welcome("Mudi")
   welcome("Musa")
   welcome("Amina")
   ```

## Logical Reasoning Challenge

A programmer writes:

```python
balance = 500
withdraw = 600
balance -= withdraw
```

The program now says `-100`.

Is this a programming error?

Answer: No. The code followed the instructions correctly. The mistake is in the business logic because it did not check whether enough money existed before allowing the withdrawal.

## AI Builder Scenario

If you inherit a codebase you do not understand, first read and understand the existing code. Do not rewrite everything blindly.

### Final Assessment

1. Why are meaningful variable names important?
   - A. They make programs faster
   - B. They improve readability
   - C. They reduce RAM usage
   - D. They remove bugs automatically
   - Answer: ✅ B

2. Which error prevents a program from starting?
   - A. Logic Error
   - B. Syntax Error
   - C. Runtime Error
   - D. Design Error
   - Answer: ✅ B

3. What is refactoring?
   - A. Changing the behavior of a program
   - B. Improving code without changing its behavior
   - C. Deleting code
   - D. Adding comments
   - Answer: ✅ B

4. What should you do before trusting AI-generated code?
   - A. Deploy immediately
   - B. Understand and test it
   - C. Ignore it
   - D. Rewrite it from scratch every time
   - Answer: ✅ B

5. Which habit best describes a professional programmer?
   - A. Guessing
   - B. Reading error messages carefully
   - C. Ignoring bugs
   - D. Copying code without understanding it
   - Answer: ✅ B

## Module 2 Cheat Sheet

### Variables

Named containers for storing data.

### Data Types

- Integer
- Float
- String
- Boolean

### Operators

- Arithmetic
- Comparison
- Logical

### Control Flow

- if
- else
- elif
- switch
- for
- while
- break
- continue

### Functions

Reusable blocks of code.

### Parameters

Placeholders.

### Arguments

Actual values.

### Return

Sends data back.

### Collections

- List: ordered collection
- Dictionary: key → value
- Set: unique values only

### Errors

| Type    | Meaning          |
| ------- | ---------------- |
| Syntax  | Won't run        |
| Runtime | Crashes          |
| Logic   | Runs incorrectly |

### Debugging

```
Observe
↓
Reproduce
↓
Locate
↓
Understand
↓
Fix
↓
Test
```

### Clean Code

- Meaningful names
- Small functions
- Avoid repetition
- Readable structure
- Simple logic

## Module 2 Summary

Congratulations. You have completed the **Programming Fundamentals Review**.

You reviewed the core building blocks of programming:

- Variables and data types
- Operators and expressions
- Input and output
- Type conversion
- Control flow
- Functions
- Parameters and return values
- Scope
- Lists, strings, dictionaries, and sets
- Reading code
- Debugging
- Refactoring
- Clean code
- AI-assisted programming

These concepts are the foundation of almost every programming language. Whether you are writing Python, JavaScript, Go, Java, or using AI tools like ChatGPT or GitHub Copilot, these fundamentals remain the same.

The strongest AI Builders are not the ones who memorize syntax. They are the ones who understand these programming principles and can apply them to solve real problems.
