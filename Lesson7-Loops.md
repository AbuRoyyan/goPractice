# Lesson 7: Loops in Go 🔁

Loops allow your program to **repeat actions automatically** instead of writing the same code many times.

In Go, the main loop is the **`for` loop**.

---

# 1. Basic `for` Loop

Example:

```go
package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

}
```

### Output

```
1
2
3
4
5
```

### How it works

```
for i := 1;   i <= 5;   i++
    start     condition  increase
```

* `i := 1` → start counting from 1
* `i <= 5` → run while i is less than or equal to 5
* `i++` → add 1 each time

---

# 2. Loop with User Input

Example:

```go
package main

import "fmt"

func main() {

	var number int

	fmt.Println("Enter a number:")
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		fmt.Println(i)
	}

}
```

### Example run

```
Enter a number:
5
1
2
3
4
5
```

The loop prints numbers **from 1 to the number the user entered**.

---

# 3. Infinite Loop

Sometimes programs must run **forever**.

Example:

```go
for {
	fmt.Println("This keeps running")
}
```

⚠️ This runs forever unless you stop it with **Ctrl + C**.

---

# Practice Tasks 🧠

Try these 3 tasks.

---

## Task 1 – Count to 10

Write a program that prints:

```
1
2
3
4
5
6
7
8
9
10
```

---

## Task 2 – Even Numbers

Print all **even numbers from 2 to 20**.

Expected output:

```
2
4
6
8
10
12
14
16
18
20
```

---

## Task 3 – User Countdown

Ask the user for a number and count **down to 1**.

Example:

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

💡 **Tip:**
Countdown uses `i--` instead of `i++`.

Example idea:

```go
for i := number; i >= 1; i-- {
```

---

Send me:

* Your **code**
* Your **output**

Then I’ll review it and move you to **Lesson 8 (While-style loops & break/continue)**.
