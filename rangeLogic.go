package main

import "fmt"

func main() {
	var score int

	fmt.Println("Enter score: ")
	fmt.Scanln(&score)

	if score < 0 || score > 100 {
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

}
