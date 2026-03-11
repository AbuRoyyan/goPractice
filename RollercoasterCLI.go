package main

import "fmt"

func main() {
	var height int
	var age int
	var bill int
	var photo string

	fmt.Println("Welcome to the Rollercoaster!")
	fmt.Println("What is your height in cm? ")
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
