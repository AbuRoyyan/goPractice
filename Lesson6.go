package main

import "fmt"

func main() {

	// Examples

	// age := 20

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// }

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("You can vote")
	// } else {
	// 	fmt.Println("You are to young to vote")
	// }

	// Task 1 - Voting Eligibility
	// var age int
	// fmt.Println("What is your age?")
	// fmt.Scanln(&age)
	// if age >= 18 {
	// 	fmt.Println("You are eligible to vote")
	// } else {
	// 	fmt.Println("You are not eliglibe to vote")
	// }
	// fmt.Scanln(&age)

	// Task 2 - Pass or Fail
	// var score int
	// fmt.Println("What is your exam score")
	// fmt.Scanln(&score)
	// if score >= 50 {
	// 	fmt.Println("You passed the exam")
	// } else {
	// 	fmt.Println("You failed the exam")
	// }

	// Task 3 - Number Checker

	var number int

	// fmt.Println("Choose a number")
	// fmt.Scanln(&number)
	// if number > 0 {
	// 	fmt.Println("The number is positive")
	// } else {
	// 	fmt.Println("The number is zero or negative")
	// }

	// 	Small Improvement (Optional) 🚀

	// You can also check three cases instead of two:

	// Positive

	// Zero

	// Negative

	// Example:
	fmt.Println("Choose a number")
	fmt.Scanln(&number)
	if number > 0 {
		fmt.Println("The number is positive")
	} else if number == 0 {
		fmt.Println("The number is zero")
	} else {
		fmt.Println("The number is negative")
	}
}
