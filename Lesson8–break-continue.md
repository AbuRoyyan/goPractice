## ⚠ Task 3 – User Countdown

Your code:

```go
for i := 5; i >= 1; i-- {
	fmt.Println(i)
}
```

✔ This works, **but it doesn't use user input yet**.

The task was to let the **user choose the number**.

Correct version:

```go
var number int

fmt.Println("Enter a number:")
fmt.Scanln(&number)

for i := number; i >= 1; i-- {
	fmt.Println(i)
}
```

Example run:

```
Enter a number:
5

5
4
3
2
1
```

---

# What You Learned in Lesson 7

You now know:

✅ Basic `for` loops
✅ Loop counters (`i++`, `i--`)
✅ Loop conditions
✅ Skipping numbers (`i += 2`)
✅ Countdown loops

This is a **major milestone in Go**. 🚀

---

# Lesson 8 – `break` and `continue`

These control how loops behave.

## `break`

Stops a loop completely.

Example:

```go
for i := 1; i <= 10; i++ {

	if i == 5 {
		break
	}

	fmt.Println(i)
}
```

Output:

```
1
2
3
4
```

The loop **stops when i becomes 5**.

---

## `continue`

Skips the current loop step.

Example:

```go
for i := 1; i <= 5; i++ {

	if i == 3 {
		continue
	}

	fmt.Println(i)
}
```

Output:

```
1
2
4
5
```

The number **3 is skipped**.

---

# Practice Tasks

## Task 1 – Stop at 7

Print numbers **1 to 10**, but stop when the number becomes **7**.

Expected output:

```
1
2
3
4
5
6
```

---

## Task 2 – Skip Number 5

Print numbers **1 to 10**, but **skip 5**.

Expected output:

```
1
2
3
4
6
7
8
9
10
```

---

## Task 3 – Secret Number Game 🎯

Create a program:

1. The secret number is **7**
2. Ask the user to guess the number
3. Keep asking **until the user guesses correctly**
4. When correct, print:

```
Correct! You guessed the number.
```

(Hint: use a **`for` loop + break**)

---

Send me your **code and output** again, and after that we’ll move to **Lesson 9 (Functions in Go)** — where Go starts becoming really powerful. 🚀
