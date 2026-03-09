# 🧮 Exercise: BMI Calculator 2.0

### Formula

```
BMI = weight / (height * height)
```

### Go Code

```go
package main

import "fmt"

func main() {
	var weight float64
	var height float64
	var bmi float64

	fmt.Print("Enter weight (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter height (m): ")
	fmt.Scanln(&height)

	bmi = weight / (height * height)

	fmt.Printf("Your BMI is %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Result: underweight")
	} else if bmi < 25 {
		fmt.Println("Result: normal weight")
	} else {
		fmt.Println("Result: overweight")
	}
}
```

### Example Output

```
Enter weight (kg): 80
Enter height (m): 1.85
Your BMI is 23.37
Result: normal weight
```

### Concepts Practiced

* `float64`
* math calculation
* `if / else if`
* `fmt.Printf` formatting

`%.2f` → prints **2 decimal places**.

---

