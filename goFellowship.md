Good exercises. They are perfect for practicing **input + conditions** in Go.

Since you asked to be **guided**, I will **not give you the full solutions**. I will show you the structure you need so you can complete them yourself.

Work through them one by one.

---

# 1️⃣ The Gatekeeper (Security Logic)

### Goal

Practice **multiple conditions** using:

```
if
else if
else
```

and the logical operator:

```
&&
```

(or possibly `||` depending on how you structure it)

---

## Step 1 — Ask for Input

You need two variables:

* `age`
* `vipPass`

Example structure:

```go
var age int
var vipPass string

fmt.Print("Enter your age: ")
fmt.Scanln(&age)

fmt.Print("Do you have a VIP pass (yes/no): ")
fmt.Scanln(&vipPass)
```

---

## Step 2 — Write the Logic

Think carefully about the rules:

1️⃣ If age **>= 18** → Allow entry

2️⃣ If age **< 18 AND vipPass == "yes"** → Allow entry

3️⃣ Otherwise → Deny

The structure will look like this:

```go
if age >= 18 {
    // allow
} else if age < 18 && vipPass == "yes" {
    // allow
} else {
    // deny
}
```

---

## Thinking Question

Why must `"yes"` have **quotes**?

What would happen if you wrote:

```
vipPass == yes
```

---

# 2️⃣ The Grade Master (Range Logic)

### Goal

Practice **numeric ranges with `else if`**.

---

## Step 1 — Ask for Score

```go
var score int

fmt.Print("Enter score: ")
fmt.Scanln(&score)
```

---

## Step 2 — Handle Invalid Scores FIRST

Always check invalid input early.

Ask yourself:

What two conditions make a score invalid?

You should check something like:

```
if score < 0 OR score > 100
```

---

## Step 3 — Write the Grade Logic

Think about ranges carefully.

Example structure:

```go
if score > 100 || score < 0 {
    fmt.Println("Invalid Score!")
} else if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 70 {
    fmt.Println("Grade: B")
} else if score >= 50 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

---

## Thinking Question

Why do we check:

```
score >= 90
```

instead of:

```
score >= 90 && score <= 100
```

What makes that unnecessary?

---

# 3️⃣ The Mini Calculator (String Logic)

### Goal

Practice **string comparison in conditions**.

---

## Step 1 — Ask for Input

You need three variables:

```
num1
num2
operation
```

Example structure:

```go
var num1 int
var num2 int
var operation string

fmt.Print("Enter first number: ")
fmt.Scanln(&num1)

fmt.Print("Enter second number: ")
fmt.Scanln(&num2)

fmt.Print("Enter operation (add/sub/mult): ")
fmt.Scanln(&operation)
```

---

## Step 2 — Conditional Logic

Think like this:

```
if operation == "add"
if operation == "sub"
if operation == "mult"
```

Example structure:

```go
if operation == "add" {
    // print num1 + num2
} else if operation == "sub" {
    // print num1 - num2
} else if operation == "mult" {
    // print num1 * num2
} else {
    fmt.Println("Unknown Operation")
}
```

---

## Thinking Question

Why must `"add"` be written with quotes?

What type of data is `"add"`?

---

# 🎯 Your Task

Write **three separate Go programs**:

1️⃣ `gatekeeper.go`
2️⃣ `grade.go`
3️⃣ `calculator.go`

Run each program and test it.

---

## When you're done

Paste your code for:

* Gatekeeper
* Grade Master
* Mini Calculator

I will review them like a **Go code reviewer** and help you improve them.

This is exactly how real engineers learn.
