Excellent answers. You understood the core idea. 👏

Let’s confirm each one.

---

# ✅ Your Answers Review

### **1️⃣ When you pass `firstName` instead of `&firstName`, what are you passing?**

You answered:

> **B — A copy of the value**

✔ Correct.

The function receives the value, not the variable itself.

So it **cannot modify the original variable**.

---

### **2️⃣ Why does the variable remain empty?**

You said:

> Because it receives a string not a pointer to the string.

✔ Correct again.

`Scanln` needs the **address of the variable** so it knows **where to store the input**.

Without the address, it has nowhere to write the value.

---

### **3️⃣ What does `&` give the function?**

You said:

> The address to the function.

✔ Exactly right.

More precisely:

```
&variable → memory address of the variable
```

This allows the function to **modify the real variable in memory**.

This concept becomes extremely important later when we study:

* pointers
* memory
* performance
* concurrency

For now, you just need to understand the **idea of passing an address**.

---

# 🟢 Lesson 6 — Conditionals (`if`)

Now your programs will **make decisions**.

Imagine a program asking:

```
Are you old enough to vote?
```

The program must check the age and **choose what to do**.

That is exactly what `if` does.

---

# Basic Structure of an `if` Statement

```go
if condition {
    // code runs if condition is true
}
```

Example:

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are an adult")
    }
}
```

Output:

```
You are an adult
```

Because `20 >= 18` is **true**.

---

# Adding an `else`

Sometimes we want **two possibilities**.

```go
if condition {
    // if true
} else {
    // if false
}
```

Example:

```go
age := 16

if age >= 18 {
    fmt.Println("You can vote")
} else {
    fmt.Println("You are too young to vote")
}
```

Output:

```
You are too young to vote
```

---

# Comparison Operators

These are used inside conditions.

| Operator | Meaning          |
| -------- | ---------------- |
| `==`     | equal            |
| `!=`     | not equal        |
| `>`      | greater than     |
| `<`      | less than        |
| `>=`     | greater or equal |
| `<=`     | less or equal    |

Example:

```go
if score >= 50 {
    fmt.Println("You passed")
}
```

---

# 🧠 Think Before Coding

If a program asks:

```
Enter your age
```

What should happen if the age is **18 or more**?

What if it is **below 18**?

This is exactly what your tasks will test.

---

# 🧪 Classwork (Do NOT ask for answers)

Write small Go programs for each.

---

# 🔹 Task 1 — Voting Eligibility

Program behavior:

1. Ask the user for their age
2. If age **>= 18**

Print:

```
You are eligible to vote
```

Otherwise print:

```
You are not eligible to vote
```

---

# 🔹 Task 2 — Pass or Fail

Program behavior:

1. Ask the user for their exam score
2. If score **>= 50**

Print:

```
You passed the exam
```

Else print:

```
You failed the exam
```

---

# 🔹 Task 3 — Number Checker

Ask the user for a number.

If the number is **greater than 0**

Print:

```
The number is positive
```

Else print:

```
The number is zero or negative
```

---

# ⚠ Important Rules

Do not:

* Skip user input
* Hardcode values

Use:

```
fmt.Scanln()
```

---

# 📩 When you reply

Send:

1️⃣ Your three programs
2️⃣ The outputs you got

Then we will continue to:

**Lesson 7 — Loops (making programs repeat work)**

This is where programming becomes powerful.