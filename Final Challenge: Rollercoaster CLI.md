# 🎢 Final Challenge: Rollercoaster CLI

This exercise combines:

* **nested conditions**
* **variables**
* **user input**

### Go Code

```go
package main

import "fmt"

func main() {
	var height int
	var age int
	var bill int
	var photo string

	fmt.Println("Welcome to the Rollercoaster!")
	fmt.Print("What is your height in cm? ")
	fmt.Scanln(&height)

	if height < 120 {
		fmt.Println("Sorry, you have to grow taller before you can ride.")
		return
	}

	fmt.Print("What is your age? ")
	fmt.Scanln(&age)

	if age < 12 {
		bill = 5
	} else if age <= 18 {
		bill = 7
	} else {
		bill = 12
	}

	fmt.Print("Do you want a photo taken? (y/n): ")
	fmt.Scanln(&photo)

	if photo == "y" || photo == "Y" {
		bill += 3
	}

	fmt.Printf("Your final bill is $%d\n", bill)
}
```

---

### Example Run

```
Welcome to the Rollercoaster!
What is your height in cm? 150
What is your age? 16
Do you want a photo taken? (y/n): y
Your final bill is $10
```

Explanation:

```
Youth ticket = $7
Photo = $3
Total = $10
```

---

# 🧠 Programming Concepts You Just Practiced

✔ nested `if` statements
✔ `else if` chains
✔ variables (`bill`)
✔ user input
✔ arithmetic operations
✔ logical OR `||`

These are **core programming fundamentals** used in almost every language.

---