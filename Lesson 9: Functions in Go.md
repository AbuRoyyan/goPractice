Great job, Nurudeen! 👏 Your program worked exactly as expected.

Your output:

```
Guess the secret number:
6
Wrong guess. Try again.
Guess the secret number:
7
Correct! You guessed the number.
```

✅ This shows you correctly used:

* a **loop (`for`)**
* **user input (`fmt.Scanln`)**
* a **condition (`if guess == secretNumber`)**
* **`break` to stop the loop**

These are **core programming skills**, not just in Go but in most languages.

---

# Quick Improvement (Optional)

You can improve the game by giving hints.

Example:

```go
package main

import "fmt"

func main() {

	secretNumber := 7
	var guess int

	for {

		fmt.Println("Guess the secret number:")
		fmt.Scanln(&guess)

		if guess == secretNumber {
			fmt.Println("Correct! You guessed the number.")
			break
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}

	}
}
```

Example run:

```
Guess the secret number:
10
Too high!

Guess the secret number:
4
Too low!

Guess the secret number:
7
Correct! You guessed the number.
```

---

# Lesson 9: Functions in Go 🧩

Functions help you **organize code and reuse logic**.

## Basic Function

Example:

```go
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, welcome to Go!")
}

func main() {
	sayHello()
}
```

Output:

```
Hello, welcome to Go!
```

---

## Function with Parameters

Example:

```go
func greet(name string) {
	fmt.Println("Hello", name)
}
```

Use it:

```go
greet("Nurudeen")
```

Output:

```
Hello Nurudeen
```

---

## Function that Returns a Value

Example:

```go
func add(a int, b int) int {
	return a + b
}
```

Use it:

```go
result := add(5, 3)
fmt.Println(result)
```

Output:

```
8
```

---

# Practice Tasks

## Task 1 – Greeting Function

Create a function:

```
greetUser(name string)
```

Output example:

```
Hello Nurudeen
```

---

## Task 2 – Add Two Numbers

Create a function:

```
addNumbers(a int, b int)
```

Print the result.

Example:

```
5 + 3 = 8
```

---

## Task 3 – Age Calculator

Create a function:

```
calculateAge(currentAge int)
```

It should print:

```
Next year you will be 30
```

(if input is `29`)

---

Send me your **code**, and I’ll review it like before. After that we’ll move to **Lesson 10 (Arrays & Slices)** — where Go becomes much more powerful. 🚀