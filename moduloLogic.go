package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	if number%2 == 0 { // even
		if number > 100 {
			fmt.Println("Large Even Number")
		} else {
			fmt.Println("Small Even Number")
		}
	} else {
		fmt.Println("Odd Number")
	}
}
