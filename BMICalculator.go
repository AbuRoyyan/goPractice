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
